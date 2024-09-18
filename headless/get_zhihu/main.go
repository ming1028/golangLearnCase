package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"net/url"
	"strings"
)

const itemsUrl = "https://www.zhihu.com/api/v4/columns/%s/items?ws_qiangzhisafe=1"

func main() {
	parsedURL, err := url.Parse("https://zhuanlan.zhihu.com/shuhangli")
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	// 提取路径部分
	path := parsedURL.Path

	// 去掉前导的斜杠
	path = strings.TrimPrefix(path, "/")
	client := resty.New()
	client = client.SetHeaders(map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.1.0.0 Safari/537.36",
	})
	urls := fmt.Sprintf(itemsUrl, path)

	cookie := &http.Cookie{
		Name:  "__zse_ck",
		Value: "001_iBWbpQrBbiIHlHWv8IDvHh6I8W/wx+MRCEOC2oUDoPUWBwj+=5Uqdx8aYGazNfqjRAEcau+YmGxkV0dp5h96mExkt7QQ1lYCg4x=8t/f5dM940o4E3=gtThgA45kcy=d",
	}
	client.SetCookie(cookie)
	resp, err := client.R().Get(urls)
	if err != nil {
		return
	}
	fmt.Println(string(resp.Body()))

}

type ZhiHuCrawler struct {
	Data []ZhiHuCrawlerData `json:"data"`
}

type ZhiHuCrawlerData struct {
	Created    int64  `json:"created"`
	Url        string `json:"url"`
	TitleImage string `json:"title_image"`
	Content    string `json:"content"`
	Title      string `json:"title"`
}
