package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

var (
	HttpClient *http.Client
	str        = `{"extra_info":{"app_id":4,"template_id":19,"tag_id":[7],"corp_id":"wwc5c5dbbc5c2ccf25","agent_id":1000033}}`
)

const (
	MaxIdleConnections int = 3
	RequestTimeOut     int = 30
)

func main() {
	HttpClient = &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: MaxIdleConnections,
		},
		Timeout: time.Duration(RequestTimeOut) * time.Second,
	}
	url := "http://10.68.30.129:8080"

	for i := 0; i < 10; i++ {
		go func() {
			resp, err := HttpClient.Post(
				url,
				"application/json",
				bytes.NewReader([]byte(str)),
			)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(resp)
			resp.Body.Close()
		}()
	}

	time.Sleep(time.Hour)
}
