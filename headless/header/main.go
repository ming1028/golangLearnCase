package main

import (
	"context"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"log"
)

func main() {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("disable-blink-features", "AutomationControlled"),
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36"),
	)
	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	// log the protocol messages to understand how it works.
	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	var pageContent string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://zhuanlan.zhihu.com/powertrain`),
		network.Enable(),
		chromedp.WaitVisible(`.css-9w3zhd`, chromedp.ByQuery), // 确保页面加载完成
		chromedp.OuterHTML(`html`, &pageContent, chromedp.ByQuery),
		network.SetExtraHTTPHeaders(map[string]interface{}{
			"Referer": "https://zhuanlan.zhihu.com/",
			"Origin":  "https://zhuanlan.zhihu.com",
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
}
