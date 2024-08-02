package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/chromedp/cdproto/network"
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
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	// 创建上下文
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// 定义第一个网址和包含 JavaScript 的 HTML 文件路径
	firstURL := "https://zhuanlan.zhihu.com"
	localFilePath := "D:/www/go/src/golangLearnCase/headless/chromedp_js/d2.js"
	secondURL := "https://zhuanlan.zhihu.com/c_1052934733759131648"

	// 读取本地 HTML 文件内容
	content, err := os.ReadFile(localFilePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// 提取 <script> 标签中的 JavaScript 代码
	// script := extractJavaScript(string(content))

	// 任务列表
	var pageContent, res string
	tasks := chromedp.Tasks{
		// 访问第一个网址
		chromedp.Navigate(firstURL),
		chromedp.Sleep(5 * time.Second), // 等待页面加载

		// 在当前页面执行 JavaScript 代码生成 cookie
		chromedp.Evaluate(string(content), &res),
		chromedp.Sleep(5 * time.Second), // 等待脚本执行

		// 获取并打印 cookies
		chromedp.ActionFunc(func(ctx context.Context) error {
			cookies, err := network.GetCookies().Do(ctx)
			if err != nil {
				return err
			}
			for _, cookie := range cookies {
				log.Printf("Cookie: %s = %s\n", cookie.Name, cookie.Value)
			}
			return nil
		}),

		// 访问第二个网址
		chromedp.Navigate(secondURL),
		chromedp.WaitVisible(`.css-9w3zhd`, chromedp.ByQuery), // 确保页面加载完成
		chromedp.OuterHTML(`html`, &pageContent, chromedp.ByQuery),
		chromedp.Sleep(5 * time.Second), // 等待页面加载
	}

	// 执行任务
	err = chromedp.Run(ctx, tasks)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pageContent)
}

// 提取 <script> 标签中的 JavaScript 代码
func extractJavaScript(htmlContent string) string {
	start := "<script>"
	end := "</script>"
	startIndex := strings.Index(htmlContent, start)
	endIndex := strings.Index(htmlContent, end)
	if startIndex == -1 || endIndex == -1 || startIndex >= endIndex {
		return ""
	}
	return htmlContent[startIndex+len(start) : endIndex]
}
