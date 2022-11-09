package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "/hello"
	fmt.Println(str[0] == '/') // char/unit8
	for _, u := range str {
		fmt.Print(u, string(u), " ")
	}
	fmt.Println()
	// 修改string转为byte，然后修改
	strByte := []byte(str)
	fmt.Println(strByte)
	for _, s := range strByte {
		fmt.Print(string(s), " ")
	}
	fmt.Println()
	// 截取
	strSli := strByte[:3]
	fmt.Println(string(strSli))
	fmt.Println("字节数：", len(str), "字符数：", utf8.RuneCountInString(str))

	fmt.Println(utf8.RuneCountInString("中国hello"), len("中国hello"))
	// 字符串连接
	b := bytes.NewBufferString(str)
	b.WriteString(" ")
	b.WriteString(" world")
	fmt.Println(b.String())

	i := 65
	fmt.Println(string(i))

	buffStr := bytes.NewBufferString("New")
	buffStr.WriteString("Buffer")
	buffStr.WriteString("String")
	fmt.Println(buffStr.String())
}
