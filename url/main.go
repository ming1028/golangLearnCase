package main

import (
	"fmt"
	"net/url"
)

func main() {
	// 给定的URL
	link := "https://36kr.com/information/technology/p/2803742467454343"

	// 解析URL
	parsedURL, err := url.Parse(link)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	// 获取域名
	domain := parsedURL.Host
	fmt.Println("Domain:", parsedURL.Scheme, domain)
}
