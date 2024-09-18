package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/zeromicro/go-zero/core/logx"
)

func main() {
	// 创建一个新的采集器
	c := colly.NewCollector(
		// 设置用户代理
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36"),
	)
	c.OnError(func(_ *colly.Response, err error) {
		logx.Error("collyError", logx.Field("err", err))
	})

	/*c.OnHTML(".kr-flow-article-item", func(e *colly.HTMLElement) {
		fmt.Println(e.DOM.Html())
		thumbElement := e.DOM.Find("article-item-pic.scaleBig").First()
		thumbElement.Length()
		thumbHref, ok := thumbElement.Attr("src")
		if ok {
			fmt.Println(thumbHref)
		}
	})*/
	c.OnHTML("img.scaleBig", func(element *colly.HTMLElement) {
		fmt.Println(element.DOM.Html())
	})
	err := c.Visit("https://36kr.com/information/technology/") // 替换为你要抓取的目标网站
	if err != nil {
		logx.Error("Failed to visit the website", logx.Field("err", err))
	}
}
