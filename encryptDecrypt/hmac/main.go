package main

import (
	hmac2 "crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	key := "hello world"
	data := "234242hello"
	hmac := hmac2.New(md5.New, []byte(key))
	hmac.Write([]byte(data))
	fmt.Println(hex.EncodeToString(hmac.Sum([]byte(""))))
	/*hmac2.New(md5.New, []byte(key))
	hmac.Write([]byte(data))
	hex.EncodeToString(hmac.Sum([]byte("")))*/
}
