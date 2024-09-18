package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
	/*dsn := "root@root@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})*/

	// resp := make(map[string]string)
	// apiUrl := "https://48.newspush.eastmoney.com/sse?cb=icomet_cb_0&cname=bdc02c361aab973818f3583fb8b5e6d5&seq=0&noop=0&token=&_=" + cast.ToString(time.Now().UnixMilli())

}

func http2(apiUrl string) {
	httpReq, err := http.NewRequest("GET", apiUrl, nil)
	httpReq.Header.Set("Content-Type", "application/json")
	if err != nil {
		return
	}
	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return
	}
	// todo 连接是否被关掉？
	defer resp.Body.Close()

	b := bytes.Buffer{} // 重置错误返回导致的脏数据
	if resp.StatusCode != http.StatusOK {
		return
	}
	for {
		buf := make([]byte, 55)
		n, err := resp.Body.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("respBodyReadError", err)
			return
		}
		b.Write(buf[:n])
		if err == io.EOF && b.Len() > 0 {
			return
		}
	}
}
