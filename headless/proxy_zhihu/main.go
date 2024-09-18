// Command proxy is a chromedp example demonstrating how to authenticate a proxy
// server which requires authentication.
package main

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"

	"github.com/chromedp/cdproto/fetch"
	"github.com/chromedp/chromedp"
)

func main() {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		// 1) specify the proxy server.
		// Note that the username/password is not provided here.
		// Check the link below for the description of the proxy settings:
		// https://www.chromium.org/developers/design-documents/network-settings
		chromedp.ProxyServer("http-dyn.abuyun.com:9020"),
		// By default, Chrome will bypass localhost.
		// The test server is bound to localhost, so we should add the
		// following flag to use the proxy for localhost URLs.
		chromedp.Flag("proxy-bypass-list", "<-loopback>"),
		chromedp.Flag("enable-automation", false), // 防止监测webdriver
		chromedp.Flag("disable-blink-features", "AutomationControlled"),
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36"),
	)
	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	// log the protocol messages to understand how it works.
	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	// 3) handle the Fetch.AuthRequired event and provide the username/password to the proxy
	// We will disable the fetch domain and cancel the event handler once the proxy is
	// authenticated to reduce the overhead. If your project needs the fetch domain to be enabled,
	// then you should change the code accordingly.
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
						// Chrome will remember the credential for the current instance,
						// so we can disable the fetch domain once credential is provided.
						// Please file an issue if Chrome does not work in this way.
						fetch.Disable(),
					)
					// and cancel the event handler too.
					lcancel()
				}()
			}
		}
	})

	var pageContent string
	if err := chromedp.Run(ctx,
		// 2) enable the fetch domain to handle the Fetch.AuthRequired event
		fetch.Enable().WithHandleAuthRequests(true),
		chromedp.Navigate("https://zhuanlan.zhihu.com/"),
		chromedp.WaitVisible(`body`, chromedp.ByQuery), // 确保页面加载完成
		chromedp.OuterHTML(`html`, &pageContent, chromedp.ByQuery),
	); err != nil {
		panic(err)
	}
	reader := strings.NewReader(pageContent)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Fatalf("Failed to create document: %v", err)
	}
	title := doc.Find(".css-9w3zhd").First()
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
}

// newProxy creates a proxy that requires authentication.
func newProxy() *httputil.ReverseProxy {
	return &httputil.ReverseProxy{
		Director: func(r *http.Request) {
			if dump, err := httputil.DumpRequest(r, true); err == nil {
				fmt.Printf("%s", dump)
			}
			// hardcode username/password "u:p" (base64 encoded: dTpw ) to make it simple
			if auth := r.Header.Get("Proxy-Authorization"); auth != "Basic dTpw" {
				r.Header.Set("X-Failed", "407")
			}
		},
		Transport: &transport{http.DefaultTransport},
		ErrorHandler: func(w http.ResponseWriter, r *http.Request, err error) {
			if err.Error() == "407" {
				fmt.Println("proxy: not authorized")
				w.Header().Add("Proxy-Authenticate", `Basic realm="Proxy Authorization"`)
				w.WriteHeader(407)
			} else {
				w.WriteHeader(http.StatusBadGateway)
			}
		},
	}
}

type transport struct {
	http.RoundTripper
}

func (t *transport) RoundTrip(r *http.Request) (*http.Response, error) {
	if h := r.Header.Get("X-Failed"); h != "" {
		return nil, fmt.Errorf(h)
	}
	return t.RoundTripper.RoundTrip(r)
}
