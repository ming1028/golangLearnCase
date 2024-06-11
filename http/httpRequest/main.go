package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"strings"
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
	getUrl("https://36kr.com/p/2814959596870537")
	return
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

func getUrl(url string) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return
	}

	// 使用 io.ReadAll 获取响应体内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	doc, err := html.Parse(strings.NewReader(string(body)))
	if err != nil {
		return
	}
	fmt.Println(len(extractText(doc)))
}

func extractText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	var text string
	/*for c := n.FirstChild; c != nil; c = c.NextSibling {
		text += extractText(c)
	}*/
	var c *html.Node
	if n.FirstChild != nil {
		c = n.FirstChild
	} else if n.NextSibling != nil {
		c = n.NextSibling
	}
	if c != nil {
		text = extractText(c)
	}
	return text
}
