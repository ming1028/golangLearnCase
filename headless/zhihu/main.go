package main

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
	"strings"
	"time"
)

func main() {
	// 创建上下文，并启用调试日志
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36"), // 设置自定义 User-Agent
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	// 设置超时时间
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	var pageContent string
	cv := `001_iBWbpQrBbiIHlHWv8IDvHh6I8W/wx+MRCEOC2oUDoPUWBwj+=5Uqdx8aYGazNfqjRAEcau+YmGxkV0dp5h96mExkt7QQ1lYCg4x=8t/f5dM940o4E3=gtThgA45kcy=d`
	err := chromedp.Run(ctx,
		chromedp.ActionFunc(func(ctx context.Context) error {
			expr := cdp.TimeSinceEpoch(time.Now().Add(5 * 24 * time.Hour))
			return network.SetCookie("__zse_ck", cv).
				WithExpires(&expr).
				WithDomain(".zhihu.com").
				Do(ctx)
		}),
		chromedp.Navigate(`https://zhuanlan.zhihu.com/p/520275558`),
		network.Enable(),
		chromedp.WaitVisible(`body`, chromedp.ByQuery), // 确保页面加载完成
		chromedp.OuterHTML(`html`, &pageContent, chromedp.ByQuery),
		network.SetExtraHTTPHeaders(network.Headers(map[string]interface{}{})),
	)
	if err != nil {
		log.Fatal(err)
	}

	// 打印页面标题
	reader := strings.NewReader(pageContent)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Fatalf("Failed to create document: %v", err)
	}

	title := doc.Find(".css-9w3zhd").First()
	fmt.Println(title.Text())
	title.Find("meta").Each(func(i int, s *goquery.Selection) {
		dp, ok := s.Attr("itemprop")
		if !ok {
			return
		}
		content, ok := s.Attr("content")
		if !ok || content == "" {
			return
		}
		switch dp {
		case "headline":
			fmt.Println(content)
		case "url":
			fmt.Println(strings.Trim(content, "//"))
		case "dateModified":
			utcTime, err := time.Parse(time.RFC3339, content)
			if err != nil {
				logx.Errorf("Error parsing time:%v", err)
				return
			}
			fmt.Println(utcTime.Local().Unix())
		default:
			return
		}
	})
	/*doc.Find(".ColumnHomeColumnCard").Each(func(i int, s *goquery.Selection) {
		thumbHref, ok := s.Attr("href")
		if ok {
			fmt.Println(thumbHref)
		}
	})*/

	/*var cookies []*network.Cookie
	if err := chromedp.Run(ctx,
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			cookies, err = network.GetCookies().Do(ctx)
			return err
		}),
	); err != nil {
		log.Fatal(err)
	}

	// 打印所有 Cookie
	for _, cookie := range cookies {
		fmt.Printf("Name: %s, Value: %s\n", cookie.Name, cookie.Value)
	}*/
	//doc.Find(".css-9w3zhd").Each(func(i int, s *goquery.Selection) {
	//	s.Find("meta").Each(func(i int, s *goquery.Selection) {
	//		/*dp, ok := s.Attr("itemprop")
	//		if !ok {
	//			return
	//		}
	//		if dp != "dateModified" {
	//			return
	//		}
	//		t, ok := s.Attr("content")
	//		if !ok {
	//			return
	//		}
	//		utcTime, err := time.Parse(time.RFC3339, t)
	//		if err != nil {
	//			fmt.Println("Error parsing time:", err)
	//			return
	//		}
	//
	//		// 获取本地时区
	//		//localTime := utcTime.Local().Unix()
	//
	//		// 打印本地时间
	//		fmt.Println("Local Time:", utcTime.Unix())*/
	//		dp, ok := s.Attr("itemprop")
	//		if !ok {
	//			return
	//		}
	//		content, ok := s.Attr("content")
	//		if !ok || content == "" {
	//			return
	//		}
	//		switch dp {
	//		case "headline":
	//			fallthrough
	//		case "url":
	//			fmt.Println(strings.Trim(content, "//"))
	//		case "dateModified":
	//			return
	//			utcTime, err := time.Parse(time.RFC3339, content)
	//			if err != nil {
	//				logx.Errorf("Error parsing time:%v", err)
	//				return
	//			}
	//			fmt.Println(utcTime.Local().Unix())
	//		default:
	//			return
	//		}
	//	})
	//})
}
