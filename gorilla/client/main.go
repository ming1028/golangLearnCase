package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"time"
)

func main() {
	ws, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ws.Close()

	go func() {
		for {
			_ = ws.WriteMessage(websocket.BinaryMessage, []byte("ping"))
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		_, data, err := ws.ReadMessage()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("receive message:", string(data))
	}
}
