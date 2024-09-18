package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"net/http"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	http.HandleFunc("/ws", wsUpGrader)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func wsUpGrader(w http.ResponseWriter, r *http.Request) {
	conn, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	for {
		// 接受消息
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("server receive messageType", messageType, "message", string(message))
		// 发送消息
		err = conn.WriteMessage(messageType, []byte("pong"))
		if err != nil && err != io.EOF {
			fmt.Println(err)
			return
		}
	}
}
