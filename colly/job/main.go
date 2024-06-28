package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

func main() {
	collector := colly.NewCollector()
	// 在请求之前调用
	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("请求地址：", request.URL)
	})

	// 错误
	collector.OnError(func(response *colly.Response, err error) {
		fmt.Println("Error:", err)
		err = response.Request.Retry()
		if err != nil {
			fmt.Println("retry:", err)
		}
	})
	collector.OnResponse(func(response *colly.Response) {
		fmt.Println("response body length:", len(response.Body))

	})
	collector.OnHTML(".job-card-body clearfix", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})
	collector.Visit("https://www.zhipin.com/wapi/zpgeek/search/joblist.json?scene=1&query=%E7%94%B5%E6%B0%94%E8%87%AA%E5%8A%A8%E5%8C%96&city=101020100&experience=&payType=&partTime=&degree=&industry=&scale=&stage=&position=&jobType=&salary=&multiBusinessDistrict=&multiSubway=&page=1&pageSize=30")
}
