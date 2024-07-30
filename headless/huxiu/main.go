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
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	// 变量用于存储页面标题
	var pageContent string

	// 运行任务
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://www.huxiu.com/article/3300911.html`),
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
	ts := doc.Find(".article__time").First()
	fmt.Println(ts.Text())
	/*doc.Find(".article-item-wrap").Each(func(i int, s *goquery.Selection) {
		titleElement := s.Find(".channel-title").First()
		if titleElement.Nodes == nil {
			titleElement = s.Find(".content-title").First()
		}
		pubTime := s.Find(".bottom-line__time").First()

		href := s.Find(".img-wrap").First()
		if href.Nodes == nil {
			href = s.Find(".content-wrap").First()
		}
		url := ""
		articleUrl, ok := href.Attr("href")
		if ok {
			url = articleUrl
		}
		thumbElement := href.Find(".img").First()
		thumbHref, ok := thumbElement.Attr("src")
		pic := ""
		if ok {
			pic = thumbHref
		}
		fmt.Println(titleElement.Text())
		fmt.Println(pubTime.Text(), url, pic)
	})*/
}
