package main

import (
	"encoding/json"
	"fmt"
)

type hub struct {
	c map[*connection]bool
	b chan []byte
	r chan *connection
	u chan *connection
}

var h = hub{
	c: make(map[*connection]bool),
	u: make(chan *connection),
	b: make(chan []byte),
	r: make(chan *connection),
}

func (h *hub) run() {
	for {
		select {
		case c := <-h.r: // 连接保存？
			// socket连接信息
			h.c[c] = true
			c.data.Ip = c.ws.RemoteAddr().String()
			c.data.Type = "handshake"
			c.data.UserList = user_list
			dataB, _ := json.Marshal(c.data)
			fmt.Println(2, string(dataB))
			c.sc <- dataB // 用户连接信息压入用户接收socket？
		case c := <-h.u: // 删除连接？
			if _, ok := h.c[c]; ok {
				delete(h.c, c)
			}
		case data := <-h.b:
			for c := range h.c { // 用户socket连接信息map
				select {
				case c.sc <- data: // 将登录用户数据发送给各个连接 c.sc 无缓存channel
				default:
					delete(h.c, c)
					close(c.sc)
				}
			}
		}
	}
}
