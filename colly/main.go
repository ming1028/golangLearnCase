package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

func main() {

	collector := colly.NewCollector(
		colly.AllowedDomains("https://www.collinsdictionary.com/"),
	)
	// 在请求之前调用
	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("123")
		fmt.Println(request.URL)
	})
	// 错误
	collector.OnError(func(response *colly.Response, err error) {
		fmt.Println("123")
		fmt.Println(err, response.StatusCode)
	})
	collector.OnResponse(func(response *colly.Response) {
		fmt.Println("123")
		fmt.Println(string(response.Body))
	})
	collector.OnHTML(".sense", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
		definition := e.ChildText("div.definition")
		fmt.Println("Definition:", definition)
	})

	// 接收到的内容是XML ,则在之后调用
	collector.OnXML("//h1", func(e *colly.XMLElement) {
		fmt.Println(e.Text)
	})

	// 回调后调用
	collector.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})
	err := collector.Visit("https://baidu.com")
	fmt.Println(err)
}
