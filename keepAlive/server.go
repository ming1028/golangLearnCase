package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	server := http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(Index),
	}
	// 如果server的Handler为nil,将使用DefaultServeMux(全局变量)作为handler
	http.HandleFunc("/index", Index) // 使用默认DefaultServeMux
	// 方式2:使用http.NewServeMux 多路复用器中的Handle或者HandleFunc定义各种路由，赋值给server.Handler
	server.ListenAndServe()
}

func Index(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)
	fmt.Println("收到请求", r.RemoteAddr, string(b))
	w.Write([]byte("ok"))
}
