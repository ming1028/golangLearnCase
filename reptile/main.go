package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

var (
	reEmail = `\w+@\w+\.\w+` // w代表大小写字母、数字、下划线
	// s?有或者没有s
	// +代表出1次或多次
	// \s\S各种字符
	// +?代表贪婪模式
	reLinke = `href="(https?://[\s\S]+?)"`
)

func main() {
	resp, err := http.Get("https://tieba.baidu.com/p/6051076813?red_tag=1573533731")
	if err != nil {
		log.Fatalf("http get err:%#v\n", err)
	}
	defer resp.Body.Close()

	pageBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("ioutil ReadAll err:%#v\n", err)
	}

	rege := regexp.MustCompile(`(\d+)@qq.com`)               // 传入正则表达式，得到正则表达式的struct
	res := rege.FindAllStringSubmatch(string(pageBytes), -1) // n 匹配数量
	for _, qq := range res {
		fmt.Println(qq)
		fmt.Println("email:", qq[0])
		fmt.Println("qq:", qq[1])
	}
}
