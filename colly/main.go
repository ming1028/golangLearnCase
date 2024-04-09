package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

func main() {
	collector := colly.NewCollector()
	// 在请求之前调用
	collector.OnRequest(func(request *colly.Request) {
		fmt.Println(request.URL)
	})
	// 错误
	collector.OnError(func(response *colly.Response, err error) {
		fmt.Println(err, response.StatusCode)
	})
	// 收到响应标头后调用
	collector.OnResponseHeaders(func(response *colly.Response) {
		fmt.Println(response.Headers)
	})
	// 收到回复后调用
	collector.OnResponse(func(response *colly.Response) {
		fmt.Println(response.Request.URL)
	})
	// 如果收到的内容是HTML ,则在之后调用
	collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})
	collector.OnHTML("tr td:nth-of-type(1)", func(e *colly.HTMLElement) {
		fmt.Println("First column of a table row:", e.Text)
	})

	// 接收到的内容是XML ,则在之后调用
	collector.OnXML("//h1", func(e *colly.XMLElement) {
		fmt.Println(e.Text)
	})

	// 回调后调用
	collector.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})
}
