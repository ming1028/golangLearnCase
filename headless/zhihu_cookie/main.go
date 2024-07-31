package main

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/target"
	"log"
	"strings"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

func main() {
	// 创建一个上下文
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// 创建一个超时上下文
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	// 运行任务
	var cookies []*network.Cookie
	var pageContent string
	err := chromedp.Run(ctx,
		network.Enable(),
		chromedp.Navigate("https://zhuanlan.zhihu.com/"),
		chromedp.WaitVisible(`body`),
		chromedp.Click(`a.ColumnHomeColumnCard[target="_blank"]`, chromedp.NodeVisible),
		chromedp.Sleep(2*time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// 获取所有窗口的目标
			targets, err := chromedp.Targets(ctx)
			if err != nil {
				return err
			}

			// 找到新打开的窗口
			var newTargetID target.ID
			for _, t := range targets {
				if t.OpenerID != "" {
					newTargetID = t.TargetID
					break
				}
			}

			// 切换到新窗口的上下文
			if newTargetID != "" {
				newCtx, cancel := chromedp.NewContext(ctx, chromedp.WithTargetID(newTargetID))
				defer cancel()
				ctx = newCtx
				cookies, _ = network.GetCookies().Do(ctx)
				// 在新窗口中获取 HTML 内容
				return chromedp.Run(ctx,
					chromedp.WaitVisible("body"),
					chromedp.OuterHTML("html", &pageContent),
				)
			}
			return nil
		}),
		/*chromedp.WaitVisible(`body`),
		chromedp.OuterHTML(`html`, &pageContent, chromedp.ByQuery),*/
	)
	if err != nil {
		log.Fatal(err)
	}

	reader := strings.NewReader(pageContent)

	// 使用 goquery.NewDocumentFromReader 创建一个 goquery.Document
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Fatalf("Failed to create document: %v", err)
	}
	s := doc.Find(".css-9w3zhd").First()
	var url string
	s.Find("meta").Each(func(i int, s *goquery.Selection) {
		dp, ok := s.Attr("itemprop")
		if !ok {
			return
		}
		content, ok := s.Attr("content")
		if !ok || content == "" {
			return
		}
		switch dp {

		case "url":
			url = strings.Trim(content, "//")
		default:
			return
		}
	})
	fmt.Println(url)
	// 打印获取到的 Cookie
	fmt.Println("Cookies:")
	for _, cookie := range cookies {
		fmt.Printf("Name: %s, Value: %s, Domain: %s, Path: %s, Expires: %v\n",
			cookie.Name, cookie.Value, cookie.Domain, cookie.Path, cookie.Expires)
	}
}
