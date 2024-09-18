package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

type WsMessageReq struct {
	MsgType int32           `json:"msgType"`
	Data    json.RawMessage `json:"data"`
}

type VoiceToTextReq struct {
	Id         int64  `json:"id"`
	DataStream []byte `json:"dataStream"` // 数据流
	Index      int32  `json:"index"`      // 片段索引
	IsEnd      int32  `json:"isEnd"`      // 是否结束
}

func main() {
	header := http.Header{}
	header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IiIsInVzZXJJZCI6MSwiaXNzIjoibGFuZ3VhZ2VSZWFjdG9yIiwiZXhwIjoxNzI3MDczMzU4fQ.-NSXPQ3pFA-0O_0IOCC_Ulfh8pwInDeW388azxB8IJQ")
	ws, _, err := websocket.DefaultDialer.Dial("ws://localhost:8099/ws", header)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ws.Close()
	/*ws.SetPongHandler(func(appData string) error {
		return ws.WriteMessage(websocket.PongMessage, []byte(appData))
	})*/
	/*go func() {
		i := 1
		for {
			jsonRaw, _ := json.Marshal(&VoiceToTextReq{
				DataStream: []byte{1, 2, 3, 3, 3, 3, 33},
				Index:      int32(i),
			})
			wdjson, _ := json.Marshal(&WsMessageReq{
				MsgType: 1,
				Data:    jsonRaw,
			})
			err := ws.WriteMessage(websocket.PingMessage, wdjson)
			if err != nil {
				fmt.Println("writeError", err)
			}
			time.Sleep(time.Second * 10)
			i++
		}
	}()*/

	go func() {
		for {
			_, message, err := ws.ReadMessage()
			if err != nil {
				log.Println("Read error:", err)
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Printf("Unexpected close error: %v", err)
				}
				time.Sleep(time.Second * 2)
				break
			}
			log.Printf("Received message: %s", message)
		}
	}()

	jsonRaw, _ := json.Marshal(&VoiceToTextReq{
		DataStream: []byte{1, 2, 3, 3, 3, 3, 33},
		Index:      int32(1),
	})
	wdjson, _ := json.Marshal(&WsMessageReq{
		MsgType: 1,
		Data:    jsonRaw,
	})
	err = ws.WriteMessage(websocket.TextMessage, wdjson)
	if err != nil {
		fmt.Println("writeError", err)
		return
	}
	i := 1
	for {
		jsonRaw, _ := json.Marshal(&VoiceToTextReq{
			DataStream: []byte{1, 2, 3, 3, 3, 3, 33},
			Index:      int32(i),
			Id:         8,
		})
		wdjson, _ := json.Marshal(&WsMessageReq{
			MsgType: 1,
			Data:    jsonRaw,
		})
		err := ws.WriteMessage(websocket.TextMessage, wdjson)
		if err != nil {
			fmt.Println("writeError", err)
			return
		}
		time.Sleep(time.Second * 30)
		i++
	}
}
