package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"strings"
	"unicode"
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
	// getUrl(`https://www.etc.org.cn/UserFiles/Article/file/6384005594877351547983703.pdf`)
	CleanHtml("https://www.researchgate.net/profile/Ping-Guo-24/publication/341980871_The_First_Principles_for_Artificial_Intelligence/links/5f4cc0d3458515a88b96bb40/The-First-Principles-for-Artificial-Intelligence.pdf")
	return
	/*html := `<html><body><h1>My Heading</h1><p>This is a paragraph.</p><p>This is another paragraph.</p></body></html>`
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

	time.Sleep(time.Hour)*/
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
	doc, err := html.Parse(strings.NewReader(string(body)))
	if err != nil {
		return
	}

	txt := extractText(doc)
	fmt.Println(txt)
	return
	// goquery解析
	/*doc, _ := goquery.NewDocumentFromReader(bytes.NewReader(body))
	doc.Find("*").Each(func(index int, element *goquery.Selection) {
		// 打印标签名
		tag := goquery.NodeName(element)
		need := []string{"div", "p", "title"}
		ms := mapset.NewSet[string](need...)
		if !ms.Contains(tag) {
			return
		}
		fmt.Printf("Tag: %s\n", goquery.NodeName(element))
		// 打印属性 onclick等行内属性
		for _, attr := range element.Nodes[0].Attr {
			fmt.Printf("  Attribute: %s=\"%s\"\n", attr.Key, attr.Val)
		}

		// 打印文本内容
		if text := strings.TrimSpace(element.Text()); text != "" {
			fmt.Printf("  Text: %s\n", text)
		}
	})*/
	/*body = []byte(`<html>
	  <body>
	    <h1>Welcome</h1>
	    <p onclick="alert('Hi!')">This is an example paragraph.</p>
	    <a href="javascript:void(0)">Click me</a>
	    <img src="example.jpg" onerror="alert('Error!')">
	    <style>body {font-size: 12px;}</style>
	  </body>
	</html>`)
		doc, err := html.Parse(bytes.NewReader(body))
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
	if n.Type == html.TextNode && (n.Parent.Data != "script" && n.Parent.Data != "style") {
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

func CleanHtml(url string) int64 {
	pdfParseUrl := `https://kaiyue-papi.baoyueai.com/pdf_content`
	client := &http.Client{}
	reqData := map[string]string{
		"file_url": url,
	}
	jsonData, err := json.Marshal(reqData)
	if err != nil {
		zap.L().Error("failed to marshal data", zap.Error(err))
		return 0
	}
	// 创建一个 POST 请求
	req, err := http.NewRequest("POST", pdfParseUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		zap.L().Error("failed to create request", zap.Error(err))
		return 0
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		zap.L().Error("failed to send request", zap.Error(err))
		return 0
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		zap.L().Error("failed to read response body", zap.Error(err))
		return 0
	}

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		zap.L().Error(fmt.Sprintf("unexpected status code: %d", resp.StatusCode))
		return 0
	}
	retData := make(map[string]interface{})
	_ = json.Unmarshal(body, &retData)
	data, ok := retData["data"]
	if !ok {
		return 0
	}
	code := retData["code"].(float64)
	fmt.Println(code == 200)
	contentSli := strings.FieldsFunc(data.(string), func(r rune) bool {
		return (r != '.' || r != '。') && (unicode.IsPunct(r) || unicode.IsSpace(r))
	})
	retStr := strings.Join(contentSli, "")
	dotCnt := strings.Count(retStr, "。")
	zap.L().Info("cleanHtmlContent", zap.Any("url", url), zap.Any("content", retStr), zap.Any("dotCnt", dotCnt))
	letterCnt := countEnglishLetters(retStr)
	if dotCnt > 10 && letterCnt < 20 {
		return 0
	}
	return int64(len(data.(string)))
}

func countEnglishLetters(str string) int {
	count := 0
	for _, char := range str {
		if unicode.IsLetter(char) && (char >= 'A' && char <= 'Z' || char >= 'a' && char <= 'z') {
			count++
		}
	}
	return count
}
