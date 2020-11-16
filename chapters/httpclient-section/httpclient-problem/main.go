package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var client *http.Client

func main() {

	tr := &http.Transport{
		TLSHandshakeTimeout: 10 * time.Second,
		//DisableKeepAlives:   false,
		//MaxIdleConns:        500,
		//MaxConnsPerHost: 500,
		//MaxIdleConnsPerHost: 500,

	}
	client = &http.Client{
		Transport: tr,
		Timeout:   time.Second * 30,
	}

	wg := sync.WaitGroup{}
	for i := 0; i < 40000; i++ {
		wg.Add(1)
		go func() {
			GetReply()
			defer wg.Done()
		}()
	}
	wg.Wait()
}

func GetReply() {

	strURL := "https://sampleapis.com/beers/api/ale"
	req, err := http.NewRequest("GET", strURL, nil)
	if err != nil {
		fmt.Println("\nError forming request: " + err.Error())
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("\nError reading response body: " + err.Error())
		if res != nil {
			res.Body.Close()
		}
		return
	}
	if res.StatusCode != 200 {
		fmt.Println("\nError reading response body, status code: " + res.Status)
		if res != nil {
			res.Body.Close()
		}
		return
	}

	res.Body.Close()
	fmt.Println("Result done")
}
