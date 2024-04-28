package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"time"
)

func main() {
	collector := colly.NewCollector()
	// 在请求之前调用
	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("request", request.URL)
	})
	// 错误
	collector.OnError(func(response *colly.Response, err error) {

		fmt.Println("Error:", err)
		time.Sleep(time.Second)
		err = response.Request.Retry()
		if err != nil {
			fmt.Println("retry:", err)
		}
	})
	collector.OnResponse(func(response *colly.Response) {
		fmt.Println("response body length:", len(response.Body))
	})
	collector.OnHTML(".index_news_list_p_5zOEF", func(e *colly.HTMLElement) {
		attr := e.ChildAttr("a", "href")
		fmt.Println(attr, e.Text)
	})

	// 接收到的内容是XML ,则在之后调用
	collector.OnXML("//h1", func(e *colly.XMLElement) {
		fmt.Println("OnXML", e.Text)
	})

	// 回调后调用
	collector.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	err := collector.Limit(&colly.LimitRule{
		Parallelism: 5,
	})
	err = collector.Visit("https://www.ifeng.com/")
	fmt.Println(err)
}
