package main

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	mapset "github.com/deckarep/golang-set/v2"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"
)

var (
	HttpClient *http.Client
	str        = `{"extra_info":{"app_id":4,"template_id":19,"tag_id":[7],"corp_id":"wwc5c5dbbc5c2ccf25","agent_id":1000033}}`
)

const (
	MaxIdleConnections int = 3
	RequestTimeOut     int = 30
)

func main() {
	getUrl("https://36kr.com/p/2814959596870537")
	return
	html := `<html><body><h1>My Heading</h1><p>This is a paragraph.</p><p>This is another paragraph.</p></body></html>`
	re := regexp.MustCompile("<.*?>")
	text := re.ReplaceAllString(html, "")
	fmt.Println(text)
	return
	HttpClient = &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: MaxIdleConnections,
		},
		Timeout: time.Duration(RequestTimeOut) * time.Second,
	}
	url := "http://10.68.30.129:8080"

	for i := 0; i < 10; i++ {
		go func() {
			resp, err := HttpClient.Post(
				url,
				"application/json",
				bytes.NewReader([]byte(str)),
			)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(resp)
			resp.Body.Close()
		}()
	}

	time.Sleep(time.Hour)
}

func getUrl(url string) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return
	}

	// 使用 io.ReadAll 获取响应体内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	// goquery解析
	doc, _ := goquery.NewDocumentFromReader(bytes.NewReader(body))
	doc.Find("*").Each(func(index int, element *goquery.Selection) {
		// 打印标签名
		tag := goquery.NodeName(element)
		need := []string{"div", "p", "title"}
		ms := mapset.NewSet[string](need...)
		if !ms.Contains(tag) {
			return
		}
		fmt.Printf("Tag: %s\n", goquery.NodeName(element))
		// 打印属性
		for _, attr := range element.Nodes[0].Attr {
			fmt.Printf("  Attribute: %s=\"%s\"\n", attr.Key, attr.Val)
		}

		// 打印文本内容
		if text := strings.TrimSpace(element.Text()); text != "" {
			fmt.Printf("  Text: %s\n", text)
		}
	})
	/*doc, err := html.Parse(strings.NewReader(string(body)))
	if err != nil {
		return
	}
	traverse(doc, 0)*/
}

func traverse(node *html.Node, depth int) {
	if node.Type == html.ElementNode {
		fmt.Printf("%s<%s>\n", strings.Repeat("  ", depth), node.Data)
		if node.Data == "script" {
			fmt.Println("script")
		}
		for _, attr := range node.Attr {
			fmt.Printf("%s  %s=\"%s\"\n", strings.Repeat("  ", depth), attr.Key, attr.Val)
		}
	} else if node.Type == html.TextNode {
		fmt.Printf("%s%s\n", strings.Repeat("  ", depth), node.Data)
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		traverse(child, depth+1)
	}

	if node.Type == html.ElementNode {
		fmt.Printf("%s</%s>\n", strings.Repeat("  ", depth), node.Data)
	}
}

func extractText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	var text string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text += extractText(c)
	}
	/*var c *html.Node
	if n.FirstChild != nil {
		c = n.FirstChild
	} else if n.NextSibling != nil {
		c = n.NextSibling
	}
	if c != nil {
		text = extractText(c)
	}*/
	return text
}
