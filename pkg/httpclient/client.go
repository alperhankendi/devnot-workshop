package httpclient

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/alperhankendi/devnot-workshop/pkg/log"
	"github.com/valyala/fasthttp"
	"strings"
	"time"
)

type HttpClient struct {
	client  *fasthttp.Client
	req     *fasthttp.Request
	res     *fasthttp.Response
	baseUrl string
	timeout time.Duration
}

type HttpRequest struct {
	httpClient  HttpClient
	verb        string
	urlSegments string
	headers     map[string]string
	cookies     map[string]string
	contentType string
	body        []byte
}

type HttpResponse struct {
	StatusCode int
	Body       []byte
	Error      error
}

type HttpCall struct {
	client      HttpClient
	request     HttpRequest
	response    HttpResponse
	exception   error
	startedTime time.Time
	endedTime   time.Time
	duration    time.Duration
	succeeded   bool
}

func (c *HttpCall) setContentType(contentType string) *HttpCall {
	if len(contentType) != 0 {
		c.request.contentType = contentType
	}
	return c
}

func (c *HttpCall) WithHeader(key, value string) *HttpCall {
	if len(key) != 0 && len(value) != 0 {
		if _, keyExist := c.request.headers[key]; !keyExist {
			c.request.headers[key] = value
		}
	}
	return c
}

func (c *HttpCall) WithBody(body interface{}) *HttpCall {
	if bodyContent, err := json.Marshal(body); err == nil {
		c.request.body = bodyContent
	} else {
		log.Logger.Errorf("failed to serialize request body data %v", body)
	}
	return c
}

func (c *HttpCall) Get(path string) (response *HttpResponse) {
	c.request.verb = "GET"
	c.request.urlSegments = path

	req, err := buildFastHttpRequest(c.client.baseUrl, c.request)

	if err != nil {
		return &HttpResponse{
			Error: err,
		}
	}

	c.client.req = req
	return c.client.Send()
}

func (c *HttpCall) Post(path string) (response *HttpResponse) {
	c.request.verb = "POST"
	c.request.urlSegments = path
	c.setContentType("application/json")

	req, err := buildFastHttpRequest(c.client.baseUrl, c.request)

	if err != nil {
		return &HttpResponse{
			Error: err,
		}
	}

	c.client.req = req
	return c.client.Send()
}

func buildFastHttpRequest(baseUrl string, req HttpRequest) (request *fasthttp.Request, err error) {
	var absoluteUri string
	const slash = "/"
	request = fasthttp.AcquireRequest()

	if len(baseUrl) == 0 {
		return request, errors.New("baseUrl url is should not be null")
	}

	if len(req.urlSegments) == 0 {
		return request, errors.New("request url is should not be null")
	}

	if strings.HasSuffix(baseUrl, slash) {
		absoluteUri = baseUrl + req.urlSegments
	} else {
		absoluteUri = baseUrl + "/" + req.urlSegments
	}

	request.SetRequestURI(absoluteUri)

	if len(req.verb) == 0 {
		return request, errors.New("request verb is should not be null")
	}

	request.Header.SetMethod(req.verb)

	if len(req.contentType) > 0 {
		request.Header.SetContentType(req.contentType)
	}

	if len(req.headers) > 0 {
		for key, value := range req.headers {
			request.Header.Add(key, value)
		}
	}

	if len(req.body) > 0 {
		request.SetBody(req.body)
	}

	return
}

func (c *HttpClient) Send() (response *HttpResponse) {
	c.res = fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(c.req)
	defer fasthttp.ReleaseResponse(c.res)
	response = new(HttpResponse)

	if c.timeout > 0 {
		response.Error = c.client.DoTimeout(c.req, c.res, c.timeout)
	} else {
		response.Error = c.client.Do(c.req, c.res)
	}

	var body []byte
	response.StatusCode = c.res.StatusCode()
	if response.StatusCode == fasthttp.StatusOK {
		// Verify the content type
		//contentType := c.res.Header.Peek("Content-Type")
		//if bytes.Index(contentType, []byte("application/json")) != 0 {
		//log.Logger.Errorf(fmt.Sprintf("Expected content type application/json but got %s\n", contentType))
		//}
		// Do we need to decompress the response?
		contentEncoding := c.res.Header.Peek("Content-Encoding")
		if bytes.EqualFold(contentEncoding, []byte("gzip")) {
			body, _ = c.res.BodyGunzip()
		}
		if bytes.EqualFold(contentEncoding, []byte("brotli")) {
			body, _ = c.res.BodyUnbrotli()
		} else {
			body = c.res.Body()
		}

	}

	response.Body = body
	return
}

func (res *HttpResponse) Bind(v interface{}) *HttpResponse {
	if res.Error != nil {
		return res
	}
	if len(res.Body) > 0 || res.Body != nil {
		res.Error = json.Unmarshal(res.Body, &v)
	} else {
		res.Error = errors.New(fmt.Sprintf("Http Status: %d", res.StatusCode))
	}

	return res
}

func NewHttpClient(baseUrl, timeout string) *HttpCall {
	var timeoutDuration, err = time.ParseDuration(timeout)
	if err != nil {
		log.Logger.Errorf("Failed to parse client timeout's value for %s domain. wrong value is %s", baseUrl, timeout)
		timeoutDuration = 4 * time.Second
	}

	client := &fasthttp.Client{
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	return &HttpCall{
		client: HttpClient{
			client:  client,
			baseUrl: baseUrl,
			timeout: timeoutDuration,
		},
		request: HttpRequest{
			headers: make(map[string]string),
			cookies: make(map[string]string),
		},
	}
}
