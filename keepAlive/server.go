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
	// 使用http.NewServeMux 多路复用器 Handle 或者HandleFunc  赋值给server.Handler
	http.HandleFunc("/index", Index)
	// 如果server的Handler为nil,将使用DefaultServeMux作为handler
	server.ListenAndServe()
}

func Index(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)
	fmt.Println("收到请求", r.RemoteAddr, string(b))
	w.Write([]byte("ok"))
}
