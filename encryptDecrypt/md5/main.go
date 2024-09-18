package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/spf13/cast"
	"time"
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
	fmt.Println("https://kaiyue-files.baoyueai.com" + SignUrl("/stenography/10df6599-81dd-4ab6-a1dd-25db43add777.txt", 500))
}

func SignUrl(path string, ttl int64) string {
	authKey := "UqxhWtqhq9lpm"
	expire := time.Now().Unix() + ttl
	src := path + "-" + cast.ToString(expire) + "-0-0-" + authKey
	w := md5.New()
	w.Write([]byte(src))
	bw := w.Sum(nil) // 是否增加特定前缀
	md5str2 := hex.EncodeToString(bw)
	return path + "?auth_key=" + cast.ToString(expire) + "-0-0-" + md5str2
}
