package main

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"log"
	"strings"
	"time"
)

func main() {
	// 创建一个上下文
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// 创建一个超时上下文
	ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// 变量用于存储页面标题
	var pageContent string

	// 运行任务
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://36kr.com/p/2884561632664196`),
		chromedp.WaitVisible(`body`, chromedp.ByQuery), // 确保页面加载完成
		chromedp.OuterHTML(`html`, &pageContent, chromedp.ByQuery),
	)
	if err != nil {
		log.Fatal(err)
	}

	// 打印页面标题
	reader := strings.NewReader(pageContent)

	// 使用 goquery.NewDocumentFromReader 创建一个 goquery.Document
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Fatalf("Failed to create document: %v", err)
	}
	times := doc.Find(".item-time").First()
	fmt.Println(times.Text())
	articleTime := strings.Trim(strings.TrimSpace(times.Text()), "·")
	// 定义时间格式
	layout := "2006-01-02 15:04"
	fmt.Println(articleTime)
	loc, err := time.LoadLocation("Asia/Shanghai")
	t, err := time.ParseInLocation(layout, articleTime, loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(t.Unix())
	/*doc.Find(".kr-flow-article-item").Each(func(i int, s *goquery.Selection) {
		titleElement := s.Find(".article-item-title").First()
		fmt.Println(titleElement.Text())
		thumbElement := s.Find(".scaleBig").First()
		thumbHref, ok := thumbElement.Attr("src")
		if ok {
			fmt.Println(thumbHref)
		}
		// 查找 div.scaleBig 元素

		s.Find("img.scaleBig").Each(func(j int, ss *goquery.Selection) {
			fmt.Println("Found div with class 'scaleBig'")
		})
	})*/
}
