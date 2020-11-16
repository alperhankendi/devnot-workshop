package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func main() {

	http.HandleFunc("/", rootHandlerFunc)
	fmt.Printf("Service is started.\n")
	http.ListenAndServe(":9000", nil)
}
func rootHandlerFunc(writer http.ResponseWriter, request *http.Request) {

	durationPeriod, err := strconv.ParseInt(request.URL.Path[1:], 10, 64)

	if err != nil {
		durationPeriod = 50
	}
	time.Sleep(time.Millisecond * time.Duration(durationPeriod))
	writer.WriteHeader(http.StatusOK)
	fmt.Fprintf(writer, "Ok. Sleeping %d...", durationPeriod)
}
