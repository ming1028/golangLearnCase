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
	// hmac.Sum 将哈希值追加到提供的切片后面
	// hamc.Equal 对比两个哈希值是否相等 值比较之外还会比对时间
	fmt.Println(hex.EncodeToString(hmac.Sum([]byte(""))))
	/*hmac2.New(md5.New, []byte(key))
	hmac.Write([]byte(data))
	hex.EncodeToString(hmac.Sum([]byte("")))*/
}
