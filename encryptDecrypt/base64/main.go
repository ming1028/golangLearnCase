package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	str := "123"
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
}
