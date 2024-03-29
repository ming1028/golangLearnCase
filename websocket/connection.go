package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

type connection struct {
	ws   *websocket.Conn
	sc   chan []byte
	data *Data
}

var user_list = []string{}

var wu = &websocket.Upgrader{
	ReadBufferSize:  512,
	WriteBufferSize: 512,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func myws(w http.ResponseWriter, r *http.Request) {
	ws, err := wu.Upgrade(w, r, nil) // connect连接对象
	if err != nil {
		return
	}
	c := &connection{
		sc:   make(chan []byte, 256),
		ws:   ws,
		data: &Data{},
	}
	fmt.Println(1, c)
	h.r <- c
	go c.writer()
	c.reader()
	defer func() {
		fmt.Println("断开退出")
		c.data.Type = "logout"
		user_list = del(user_list, c.data.User)
		c.data.UserList = user_list
		c.data.Content = "退出" + c.data.User
		data_b, _ := json.Marshal(c.data)
		h.b <- data_b
		h.r <- c
	}()
}

func (c *connection) writer() {
	for message := range c.sc { // 读取当前连接channel返回的信息
		fmt.Println(3, string(message))
		c.ws.WriteMessage(websocket.TextMessage, message)
	}
	c.ws.Close()
}

func (c *connection) reader() {
	for {
		_, message, err := c.ws.ReadMessage()
		fmt.Println(4, string(message), err)
		if err != nil {
			c.data.Content = "链接断开退出"
			fmt.Println("异常退出")
			h.r <- c
			break
		}
		json.Unmarshal(message, &c.data)
		switch c.data.Type {
		case "login":
			c.data.User = c.data.Content
			c.data.From = c.data.User
			user_list = append(user_list, c.data.User)
			data_b, _ := json.Marshal(c.data)
			fmt.Println("登录用户信息：", string(data_b))
			h.b <- data_b // 全局socket中用户信息channel 无缓存
		case "user":
			c.data.Type = "user"
			data_b, _ := json.Marshal(c.data)
			h.b <- data_b
		case "logout":
			c.data.Type = "logout"
			user_list = del(user_list, c.data.User)
			data_b, _ := json.Marshal(c.data)
			h.b <- data_b
			h.r <- c
		default:
			fmt.Println("=====default=====")
		}
	}
}

func del(slice []string, user string) []string {
	count := len(slice)
	if count == 0 {
		return slice
	}
	if count == 1 && slice[0] == user {
		return []string{}
	}

	var n_slice = []string{}
	for i := range slice {
		if slice[i] == user && i == count {
			return slice[:count]
		} else if slice[i] == user {
			n_slice = append(slice[:i], slice[i+1:]...)
		}
	}
	fmt.Println(n_slice)
	return n_slice
}
