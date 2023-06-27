package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type WechatResp struct {
	Errcode  int32  `json:"errcode"`
	Errmsg   string `json:"errmsg"`
	Msgid    int64  `json:"msgid"`
	RespInfo string `json:"resp_info"`
}

func main() {
	wechatResp := &WechatResp{}
	sj := `{"errcode":43101,"errmsg":"user refuse to accept the msg rid: 649a34ad-71d15d93-47ff8411"}`
	err := json.Unmarshal([]byte(sj), wechatResp)
	if err != nil {

	}
	t := time.Now().Format("01-02 15:04:05")
	fmt.Println(t)
	u := "http://www.baidu.com?a=1&b=2"
	us, err := url.Parse(u)
	if err != nil {
		return
	}
	fmt.Println(us)
	queryValues, _ := url.ParseQuery(us.RawQuery)
	fmt.Println(queryValues)
	queryValues.Add("sendId", "23")
	us.RawQuery = queryValues.Encode()
	fmt.Println(us.String())
	resp, err := http.Get("http://www.baidu.com")
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(len(string(body)), err)

	// 带参数get请求
	paramData := url.Values{}
	paramData.Set("name", "张三")
	paramData.Set("age", "18")
	urlStruct, err := url.ParseRequestURI("http://www.baidu.com")
	fmt.Println(err)
	urlStruct.RawQuery = paramData.Encode()
	fmt.Println(urlStruct, urlStruct.String())

	resp, err = http.Get(urlStruct.String())
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	fmt.Println(len(string(body)), err)

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
	fmt.Println(len(string(b)), err)

	http.HandleFunc("/go", myHandler)
	http.ListenAndServe("127.0.0.1:8000", nil)

}

// handler函数
func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr, "连接成功")
	// 请求方式：GET POST DELETE PUT UPDATE
	fmt.Println("method:", r.Method)
	// /go
	fmt.Println("url:", r.URL.Path)
	fmt.Println("header:", r.Header)
	fmt.Println("body:", r.Body)
	// 回复
	w.Write([]byte("www.5lmh.com"))
}
