package main

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"log"
	"strings"
	"time"
)

func main() {
	// 创建上下文，并启用调试日志
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("disable-setuid-sandbox", true),
		chromedp.Flag("default-charset", "UTF-8"), // 设置默认字符编码
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36"), // 设置自定义 User-Agent
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	// 设置超时时间
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	cookie := []*network.CookieParam{
		{
			Name:     "__zse_ck",
			Value:    "001_Hf=GRNtUeEMju5Fw6/9b0Hf2eRaW0SwqZVm/Gg2QDfByzBiPQ+tDFITe84/uhUci1mX/R30d1rm9gt+O+=jCq1az0o+vf51iYCEGBoCO3dvLMiUI+86MWDaxoJGy4Kj3",
			Domain:   ".zhihu.com",
			Path:     "/",
			HTTPOnly: false,
			Secure:   false,
		},
	}

	var pageContent string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://zhuanlan.zhihu.com/p/520275558`),
		chromedp.WaitVisible(`.Post-Title`, chromedp.ByQuery), // 确保页面加载完成
		chromedp.OuterHTML(`html`, &pageContent, chromedp.ByQuery),
		network.SetCookies(cookie),
	)
	if err != nil {
		log.Fatal(err)
	}

	// 打印页面标题
	reader := strings.NewReader(pageContent)
	//reader := transform.NewReader(strings.NewReader(pageContent), simplifiedchinese.GBK.NewDecoder())
	// 使用 goquery.NewDocumentFromReader 创建一个 goquery.Document
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Fatalf("Failed to create document: %v", err)
	}

	title := doc.Find(".Post-Title").First()
	fmt.Println(title.Text())

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
