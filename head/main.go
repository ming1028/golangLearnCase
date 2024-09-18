package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func getFileSize(url string) (int64, error) {
	// 创建一个 HTTP 客户端
	client := &http.Client{}

	// 创建一个 HTTP HEAD 请求
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create request: %v", err)
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to retrieve file information. status code: %d", resp.StatusCode)
	}

	// 获取 Content-Length 字段
	contentLength := resp.Header.Get("Content-Length")
	if contentLength == "" {
		return 0, fmt.Errorf("could not determine file size from headers")
	}

	// 将 Content-Length 转换为 int64
	fileSize, err := strconv.ParseInt(contentLength, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse content length: %v", err)
	}

	return fileSize, nil
}

func main() {
	url := "https://pdfs.cir.cn/JiXieDianZi/50/%E5%85%89%E8%8A%AF%E7%89%87%E8%A1%8C%E4%B8%9A%E7%8E%B0%E7%8A%B6%E5%8F%8A%E5%89%8D%E6%99%AF_3285050.pdf"
	fileSize, err := getFileSize(url)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("The file size is %d bytes.\n", fileSize)
	}
}
