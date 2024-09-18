package main

import (
	"fmt"
	"github.com/spf13/cast"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9001")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()

	for i := 0; i < 20; i++ {
		msg := " hello, world " + cast.ToString(i)
		conn.Write([]byte(msg))
	}
	time.Sleep(time.Second * 5)
	conn.Write([]byte("Q"))
	buf := [512]byte{}
	n, err := conn.Read(buf[:])
	if err != nil {
		fmt.Println("recv failed, err:", err)
		return
	}
	fmt.Println(string(buf[:n]))
}
