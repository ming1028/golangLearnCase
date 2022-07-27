package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
)

func main() {
	group := new(errgroup.Group)
	var urls = []string{
		"http://pkg.go.dev",
		"http://www.liwenzhou.com",
		"http://www.yixieqitawangzhi.com",
	}

	for _, url := range urls {
		url := url // 注意此处声明新的变量
		// 启动一个goroutine去获取url内容
		group.Go(func() error {
			resp, err := http.Get(url)
			if err == nil {
				fmt.Printf("获取%s成功\n", url)
				resp.Body.Close()
			}
			return err // 返回错误
		})
	}
	if err := group.Wait(); err != nil {
		fmt.Println(err)
	}

}
