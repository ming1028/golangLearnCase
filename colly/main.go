package main

import (
	"crypto/aes"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/gocolly/colly/v2"
	"regexp"
	"strings"
)

// todo aes ECB解密
func main() {
	collector := colly.NewCollector()
	// 在请求之前调用
	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("request", request.URL)
	})
	// 错误
	collector.OnError(func(response *colly.Response, err error) {
		fmt.Println("Error:", err)
		err = response.Request.Retry()
		if err != nil {
			fmt.Println("retry:", err)
		}
	})
	collector.OnResponse(func(response *colly.Response) {
		fmt.Println("response body length:", len(response.Body))

	})
	collector.OnHTML(".index_cententWrap__Jv8jK", func(e *colly.HTMLElement) {
		// attr := e.ChildAttr("p", "href")
		fmt.Println(e.Text)
	})
	collector.OnHTML("script[src]", func(e *colly.HTMLElement) {
		// attr := e.ChildAttr("p", "href")
		jsSrc := e.Attr("src")
		if strings.Contains(jsSrc, "app") {
			fmt.Println(strings.Replace(jsSrc, "//", "", 1))
		}
	})
	collector.OnHTML("script", func(e *colly.HTMLElement) {
		// attr := e.ChildAttr("p", "href")
		text := e.Text
		re := regexp.MustCompile(`window.initialState=`)
		matches := re.FindStringSubmatch(e.Text)
		if len(matches) >= 1 {
			initialState := strings.Replace(text, "window.initialState=", "", 1)
			fmt.Println("window.initialState:", initialState)
			encry, _ := jsonparser.GetString([]byte(initialState), "state")
			bytes, _ := DePwdCode(encry)
			fmt.Println(string(bytes))
		}
	})
	// 接收到的内容是XML ,则在之后调用
	/*collector.OnXML("//h1", func(e *colly.XMLElement) {
		fmt.Println("OnXML", e.Text)
	})*/

	// 回调后调用
	collector.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	collector.Visit("https://36kr.com/p/2815034915392002")

}

func DePwdCode(pwd string) ([]byte, error) {
	// 解密base64字符串
	pwdByte, err := base64.StdEncoding.DecodeString(pwd)
	if err != nil {
		return nil, err
	}
	// 执行AES解密
	return AesDeCrypt(pwdByte, []byte("efabccee-b754-4c"))

}
func AesDeCrypt(cypted []byte, key []byte) ([]byte, error) {
	// 创建加密算法实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// 获取块大小
	// blockSize := block.BlockSize()
	// 创建加密客户端实例
	// blockMode := cipher.NewCBCDecrypter(block, key[:blockSize]) // 使用密钥作为偏移量
	origData := make([]byte, len(cypted))
	// 这个函数也可以用来解密
	// blockMode.CryptBlocks(origData, cypted)
	for bs := 0; bs < len(cypted); bs += block.BlockSize() {
		block.Decrypt(origData[bs:bs+block.BlockSize()], cypted[bs:bs+block.BlockSize()])
	}
	block.Decrypt(origData, cypted)
	fmt.Println(string(origData))
	// 去除填充字符串
	origData, err = PKCS7UnPadding(origData)
	if err != nil {
		return nil, err
	}
	return origData, err
}

func PKCS7UnPadding(origData []byte) ([]byte, error) {
	length := len(origData)
	if length == 0 {
		return nil, errors.New("加密字符串错误")
	}
	fmt.Println("UnPadding", string(origData), origData[length:])
	unpadding := int(origData[length-1]) // 末尾填充内容，填充规则：(末尾填充3 个 3) 获取填充内容就可以得到填充长度
	return origData[:(length - unpadding)], nil
}
