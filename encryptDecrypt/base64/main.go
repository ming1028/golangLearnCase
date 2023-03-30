package main

import (
	"encoding/base64"
	"fmt"
	"github.com/spf13/cast"
)

func main() {
	d := 32
	fmt.Println(base64.StdEncoding.EncodeToString([]byte(cast.ToString(d))))
	str := "BWvOF1WtUcsnnl+wM5lO6A=="
	input := []byte(str)
	fmt.Printf("[]byte : %v\n", input)

	encodeString := base64.StdEncoding.EncodeToString(input)
	fmt.Printf("encode base64 : %v\n", encodeString)

	decodeBytes, _ := base64.StdEncoding.DecodeString(encodeString)
	fmt.Printf("decode base64: %v\n", string(decodeBytes))

	input = []byte("342+34dfsdfsf")
	urlEncode := base64.URLEncoding.EncodeToString(input)
	fmt.Printf("urlencode:%v\n", urlEncode)

	urlDecode, _ := base64.URLEncoding.DecodeString(urlEncode)
	fmt.Printf("urldecode:%v\n", string(urlDecode))

	str = "MTcy"
	s, _ := base64.StdEncoding.DecodeString(str)
	fmt.Println(string(s))
}
