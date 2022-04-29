package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func main() {
	resp, err := http.Get("http://www.baidu.com")
	fmt.Println(resp, err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body), err)

	// 带参数get请求
	paramData := url.Values{}
	paramData.Set("name", "张三")
	paramData.Set("age", "18")
	urlStruct, err := url.ParseRequestURI("http://www.baidu.com")
	fmt.Println(err)
	urlStruct.RawQuery = paramData.Encode()
	fmt.Println(urlStruct.String())

	resp, err = http.Get(urlStruct.String())
	fmt.Println(resp, err)
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	fmt.Println(string(body), err)

	// post-json
	url := "http://www.baidu.com"
	contentType := "application/json"
	postData := make(map[string]string)
	postData["name"] = "张三"
	postData["age"] = strconv.Itoa(18)
	postDataJosn, _ := json.Marshal(postData)
	resp, err = http.Post(url, contentType, strings.NewReader(string(postDataJosn)))
	// fmt.Println(resp, err)
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b), err)
}
