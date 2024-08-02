package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
)

func main() {
	/*opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.ProxyServer("http-dyn.abuyun.com:9020"),
		chromedp.Flag("proxy-bypass-list", "<-loopback>"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	ctx, cancel = chromedp.NewExecAllocator(ctx, opts...)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()
	lctx, lcancel := context.WithCancel(ctx)
	chromedp.ListenTarget(lctx, func(ev interface{}) {
		switch ev := ev.(type) {
		case *fetch.EventRequestPaused:
			go func() {
				_ = chromedp.Run(ctx, fetch.ContinueRequest(ev.RequestID))
			}()
		case *fetch.EventAuthRequired:
			if ev.AuthChallenge.Source == fetch.AuthChallengeSourceProxy {
				go func() {
					_ = chromedp.Run(ctx,
						fetch.ContinueWithAuth(ev.RequestID, &fetch.AuthChallengeResponse{
							Response: fetch.AuthChallengeResponseResponseProvideCredentials,
							Username: "HP5R1405SO6OHAMD",
							Password: "00E2A978D29A0146",
						}),
						fetch.Disable(),
					)
					lcancel()
				}()
			}
		}
	})*/

	var err error
	/*err = chromedp.Run(ctx, chromedp.Navigate("https://zhuanlan.zhihu.com"))
	if err != nil {
		panic(err)
	}*/

	content, err := os.ReadFile("D:/www/go/src/golangLearnCase/headless/chromedp_js/d2.js")
	if err != nil {
		panic(err)
	}
	var zse string
	lctx, lcancel := context.WithCancel(context.Background())
	defer lcancel()
	ctx, cancel := chromedp.NewContext(lctx)
	defer cancel()
	err = chromedp.Run(ctx, chromedp.Evaluate(string(content), &zse))
	if err != nil {
		logx.Errorw("Failed to execute JavaScript", logx.Field("err:", err))
		panic(err)
	}
	fmt.Println(zse)
}
