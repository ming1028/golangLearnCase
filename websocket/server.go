package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// go run server.go hub.go data.go connection.go
func main() {
	router := mux.NewRouter()
	go h.run()                     // select 监听
	router.HandleFunc("/ws", myws) // 连接路由
	if err := http.ListenAndServe("127.0.0.1:8080", router); err != nil {
		fmt.Println("err:", err)
	}
}
