package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	str := "md5加密"
	data := []byte(str)
	has := md5.Sum(data)
	md5Str1 := fmt.Sprintf("%x", has) // 转换成16进制
	// fmt.Sprintf("%x", md5.Sum([]byte("")))
	fmt.Println(md5Str1)

	w := md5.New()
	// io.WriteString(w, str) // 将str写入到w中 方式1
	w.Write(data)
	// 方式2
	bw := w.Sum(nil) // 是否增加特定前缀
	fmt.Println(bw)
	// hex.EncodeToString 转换成十六进制的
	md5str2 := hex.EncodeToString(bw)
	fmt.Println(md5str2)
}
