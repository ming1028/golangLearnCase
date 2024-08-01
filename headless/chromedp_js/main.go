package main

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"log"
	"time"
)

func main() {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	// 创建一个上下文和取消函数
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()
	/*// 读取外部 JavaScript 文件
	jsCode, err := os.ReadFile("./headless/chromedp_js/demo.js")
	if err != nil {
		log.Fatalf("Failed to read external.js: %v", err)
	}*/
	/*jsCode, err := os.ReadFile("./headless/chromedp_js/d2.js")
	if err != nil {
		log.Fatalf("Failed to read external.js: %v", err)
	}*/
	var err error
	jsCode := `
        document.cookie =  "__zse_ck=; expires=Mon, 20 Sep 1970 00:00:00 UTC; path=/;", document.cookie =  "__zse_ck=001_WhTwrwyPbcdYTIm3vp/tWEfYMiomyg9=iVraq7VbxGCOOr3caQXBtdR2kLuGjlTNPe9Q5KmbcFjEEFXDN2LGIA3k/Lsw4ugfTelDWC5102QC3JB3lHZ6V6Hm2XLKTlpQ; domain=.zhihu.com; expires=2024-08-08T02:32:58.000Z; path=/;";
        console.log("Cookie set: " + document.cookie);
    `
	err = chromedp.Run(ctx, chromedp.Navigate("https://zhuanlan.zhihu.com"))
	if err != nil {
		log.Fatalf("Failed to navigate to about:blank: %v", err)
	}

	// 执行 JavaScript 代码
	var res interface{}
	err = chromedp.Run(ctx, chromedp.Evaluate(string(jsCode), &res))
	if err != nil {
		log.Fatalf("Failed to execute JavaScript: %v", err)
	}

	// 打印执行结果
	fmt.Println("JavaScript executed successfully:", res)

	// 等待 cookie 生效
	time.Sleep(100 * time.Second)

	// 启动浏览器并访问一个页面
	var pageContent string
	err = chromedp.Run(ctx,
		chromedp.Navigate(`https://zhuanlan.zhihu.com/p/520275558`),
		network.Enable(),
		chromedp.WaitVisible(`body`, chromedp.ByQuery), // 确保页面加载完成
		chromedp.OuterHTML(`html`, &pageContent, chromedp.ByQuery),
	)
	if err != nil {
		log.Fatalf("Failed to navigate: %v", err)
	}
}
