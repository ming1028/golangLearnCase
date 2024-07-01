package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
	"time"
)

func main() {
	// 创建一个上下文
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// 创建一个超时上下文
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// 变量用于存储页面标题
	var pageContent string

	// 运行任务
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://36kr.com/p/2843272694778497`),
		chromedp.WaitVisible(`body`, chromedp.ByQuery), // 确保页面加载完成
		chromedp.OuterHTML(`html`, &pageContent, chromedp.ByQuery),
	)
	if err != nil {
		log.Fatal(err)
	}

	// 打印页面标题
	fmt.Println("Page Title:", pageContent)
}
