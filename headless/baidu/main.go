package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.DisableGPU,
		chromedp.NoSandbox,
		chromedp.Flag("disable-setuid-sandbox", true),
		chromedp.Flag("disable-dev-shm-usage", true),
	)
	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	// 创建一个上下文和取消函数
	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	// 启动浏览器并访问目标页面
	err := chromedp.Run(ctx, chromedp.Navigate("https://author.baidu.com/home/1747829326827698"))
	if err != nil {
		log.Fatalf("Failed to navigate to the page: %v", err)
	}

	// 等待页面加载完成
	err = chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
		// 等待网络空闲
		time.Sleep(2 * time.Second) // 等待初始加载
		for {
			var readyState string
			chromedp.Evaluate(`document.readyState`, &readyState).Do(ctx)
			if readyState == "complete" {
				break
			}
			time.Sleep(500 * time.Millisecond)
		}
		return nil
	}))
	if err != nil {
		log.Fatalf("Failed to wait for page to load: %v", err)
	}

	// 查找并点击文本内容为“文章”的div.s-tab元素
	var res string
	jsCode := `
        var elements = Array.from(document.querySelectorAll('div.s-tab'));
        var target = elements.find(el => el.textContent.trim() === '文章');
        if (target) {
            target.click();
            'Clicked on the tab with text 文章';
        } else {
            'No tab with text 文章 found';
        }
    `
	err = chromedp.Run(ctx, chromedp.Evaluate(jsCode, &res))
	if err != nil {
		log.Fatalf("Failed to click on the tab with text '文章': %v", err)
	}

	if res == "No tab with text 文章 found" {
		log.Fatalf("No tab with text '文章' found")
	}

	// 抓取页面内容
	var pageContent string
	err = chromedp.Run(ctx,
		chromedp.ActionFunc(func(ctx context.Context) error {
			// 等待网络空闲
			time.Sleep(2 * time.Second) // 等待初始加载
			for {
				var readyState string
				chromedp.Evaluate(`document.readyState`, &readyState).Do(ctx)
				if readyState == "complete" {
					break
				}
				time.Sleep(500 * time.Millisecond)
			}
			return nil
		}),
		chromedp.Sleep(time.Second*3),
		chromedp.OuterHTML(`html`, &pageContent, chromedp.ByQuery),
	)
	if err != nil {
		log.Fatalf("Failed to get page content: %v", err)
	}

	// 打印抓取到的页面内容
	fmt.Println(pageContent)
}
