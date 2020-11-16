package main

import (
	"crypto/tls"
	"fmt"
	"github.com/valyala/fasthttp"
	"sync"
	"time"
)

var client *fasthttp.Client
var lock sync.Mutex
var counter int

func main() {

	client = &fasthttp.Client{
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	wg := sync.WaitGroup{}
	for i := 0; i < 50000; i++ {
		wg.Add(1)
		go func() {
			GetReply()
			lock.Lock()
			counter++
			lock.Unlock()
			defer wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("Total hit : %d\n", counter)
}

func GetReply() string {

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	//req.SetRequestURI("https://sampleapis.com/beers/api/ale")
	req.SetRequestURI("http://localhost:9000/5")
	client.DoTimeout(req, resp, time.Millisecond*500)
	fmt.Printf("Res status: %d\n", resp.StatusCode())
	return ""
}
