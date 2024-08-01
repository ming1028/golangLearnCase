package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const zhiHuZhuanLanColumn = "https://zhuanlan.zhihu.com/api/recommendations/columns?limit=8&offset=%d&seed=%d"

func main() {
	ZlListHttpGet()
	return
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

func ArticleList(alAddress string) (interface{}, error) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.DisableGPU,
		chromedp.NoSandbox,
		chromedp.Flag("disable-setuid-sandbox", true),
		chromedp.Flag("disable-dev-shm-usage", true),
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	// 创建一个上下文
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	var err error
	expireDate := time.Now().Add(7 * 24 * time.Hour).Format("2006-01-02")
	jsCode := `
        document.cookie =  "__zse_ck=; expires=Mon, 20 Sep 1970 00:00:00 UTC; path=/;", document.cookie =  "__zse_ck=001_WhTwrwyPbcdYTIm3vp/tWEfYMiomyg9=iVraq7VbxGCOOr3caQXBtdR2kLuGjlTNPe9Q5KmbcFjEEFXDN2LGIA3k/Lsw4ugfTelDWC5102QC3JB3lHZ6V6Hm2XLKTlpQ; domain=.zhihu.com; expires=` + expireDate + `T02:32:58.000Z; path=/;";
    `
	err = chromedp.Run(ctx, chromedp.Navigate("https://zhuanlan.zhihu.com"))
	if err != nil {
		logx.Errorw("Failed to navigate to https://zhuanlan.zhihu.com", logx.Field("err", err))
		return nil, err
	}

	var res interface{}
	err = chromedp.Run(ctx, chromedp.Evaluate(jsCode, &res))
	if err != nil {
		logx.Errorw("Failed to execute JavaScript", logx.Field("err:", err))
		return nil, err
	}

	// 创建一个超时上下文
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	// 变量用于存储页面标题
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
		chromedp.Navigate(alAddress),
		chromedp.WaitVisible(`.css-9w3zhd`, chromedp.ByQuery), // 确保页面加载完成
		chromedp.OuterHTML(`html`, &pageContent, chromedp.ByQuery),
		chromedp.Sleep(time.Second*3),
	)
	if err != nil {
		logx.Errorw("chromedpError", logx.Field("address", alAddress))
		return nil, err
	}
	// 打印页面标题
	reader := strings.NewReader(pageContent)

	// 使用 goquery.NewDocumentFromReader 创建一个 goquery.Document
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		logx.Errorf("Failed to create document: %v", err)
		return nil, err
	}
	doc.Find(".css-9w3zhd").Each(func(i int, s *goquery.Selection) {
		// 根据时间过滤，只要有dateModified值才保存
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
	})
	return nil, nil
}

func ZlListHttpGet() ([]string, error) {
	client := resty.New()
	rand.Seed(time.Now().UnixNano())
	randStr := cast.ToString(rand.Intn(1000))
	offset := rand.Intn(10) + 1
	seed := rand.Intn(20) + 1
	urlAddress := fmt.Sprintf(zhiHuZhuanLanColumn, offset, seed)
	reqUrl, _ := url.Parse("http://HP5R1405SO6OHAMD:00E2A978D29A0146@http-dyn.abuyun.com:9020")
	client = client.SetTransport(&http.Transport{Proxy: http.ProxyURL(reqUrl)})
	resp, err := client.R().SetHeaders(map[string]string{
		"Content-Type": "application/json",
		"Cookie":       "__zse_ck=001_V+xVudffC4VOUr4me3VMGTwoLly0dj9bc5EatysPIENAbVVLeR9MfZVtTF99+zya+4D6K+aI1AAoBJB=Gc9g/8wD4A=dWS=kM5/9D9BzRQN0bWhA2C6Eqto/8FfLe2/g",
		"Referer":      "https://www.zhihu.com/",
		"User-Agent":   randStr,
	}).Get(urlAddress)
	if err != nil {
		logx.Infow("专栏列表请求", logx.Field("zlAddress", urlAddress))
		return nil, err
	}
	if resp.Body() == nil {
		logx.Infow("专栏列表请求接口返回body", logx.Field("zlAddress", urlAddress))
		return nil, errors.New("response body is nil")
	}
	zhiHuZhuanLan := new(ZhiHuZhuanLanColumn)
	err = json.Unmarshal(resp.Body(), zhiHuZhuanLan)
	if err != nil {
		logx.Infow("json.Unmarshal", logx.Field("zlAddress", urlAddress), logx.Field("err", err))
		return nil, err
	}
	if resp.StatusCode() != http.StatusOK || len(zhiHuZhuanLan.Data) == 0 {
		logx.Infow("专栏列表请求", logx.Field("zlAddress", urlAddress),
			logx.Field("resp", string(resp.Body())))
		return nil, errors.New(fmt.Sprintf("httpStatusCodeError:%d", resp.StatusCode()))
	}

	ret := make([]string, 0, len(zhiHuZhuanLan.Data))
	for _, d := range zhiHuZhuanLan.Data {
		if d.Url == "" {
			continue
		}
		ret = append(ret, d.Url)
	}
	return ret, nil
}

type ZhiHuZhuanLanColumn struct {
	Data []struct {
		Url string `json:"url"`
	}
}
