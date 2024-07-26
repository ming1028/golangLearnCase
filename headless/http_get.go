package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	httpClient := &http.Client{}
	request, err := http.NewRequest(http.MethodGet, "https://www.zhihu.com/api/v4/columns/designero/items?ws_qiangzhisafe=1", nil)
	if err != nil {
		return
	}
	cookie := &http.Cookie{
		Name:  "__zse_ck",
		Value: "001_UGqdg2wOk6SQZMLDzuFtpuk3a=ewKrLNDGg22Tq2rTa0/KTd=jGTIJmHEe11+jxoY2pQ9OsMw474eIT9u6jfzJwO3Y908j4A0Yp7ZrAiu1ayMP1IkOvQdW6=cK6eH6l+",
	}
	request.AddCookie(cookie)
	resp, err := httpClient.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	fmt.Println(string(b))
}
