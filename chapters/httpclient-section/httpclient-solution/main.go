package main

import (
	"fmt"
	"github.com/alperhankendi/devnot-workshop/pkg/httpclient"
	"sync"
)

var client *httpclient.HttpCall

func main() {

	strURL := "https://sampleapis.com/"

	client = httpclient.NewHttpClient(strURL, "5s")

	wg := sync.WaitGroup{}
	for i := 0; i < 50000; i++ {
		wg.Add(1)
		go func() {
			GetReply()
			defer wg.Done()
		}()
	}
	wg.Wait()
}

func GetReply() string {

	res := client.Get("beers/api/ale")
	fmt.Printf("Res status: %d", res.StatusCode)
	return ""
}
