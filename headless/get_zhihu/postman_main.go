package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://www.zhihu.com/api/v4/columns/shuhangli/items?ws_qiangzhisafe=1"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Cookie", "_xsrf=5add50ca-a3ae-40a8-950e-ddd76a2d4ab8;__zse_ck=001_4FoRacN1WsVVVhYN8Qhl1yyxW580obQ7JSHs3a/gOpNVy4fT1IG94vo5mT6ZizHtkz16j2BMe5V8M9NbN2LIwd/ATJtksLzAXE=DYj9TO2I1p3toBcqSci2Q8Uwvl1Ef; _xsrf=g4B6Q88WysU7SZRUEOiRgbAnh2cnEeZT; BEC=f7bc18b707cd87fca0d61511d015686f")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
