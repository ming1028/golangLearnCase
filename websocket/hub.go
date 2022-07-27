package main

import (
	"encoding/json"
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
			h.c[c] = true
			c.data.Ip = c.ws.RemoteAddr().String()
			c.data.Type = "handshake"
			c.data.UserList = user_list
			dataB, _ := json.Marshal(c.data)
			c.sc <- dataB
		case c := <-h.u: // 删除连接？
			if _, ok := h.c[c]; ok {
				delete(h.c, c)
			}
		case data := <-h.b:
			for c := range h.c {
				select {
				case c.sc <- data: // 数据发送
				default:
					delete(h.c, c)
					close(c.sc)
				}
			}
		}
	}
}
