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
	server.ListenAndServe()
}

func Index(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)
	fmt.Println("收到请求", r.RemoteAddr, string(b))
	w.Write([]byte("ok"))
}
