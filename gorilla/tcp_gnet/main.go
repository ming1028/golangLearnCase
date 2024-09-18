package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {
	for i := 0; i < 1000; i++ {
		go func(num int) {
			conn, err := net.Dial("tcp", "127.0.0.1:9000")
			if err != nil {
				fmt.Println(err)
				return
			}
			for {
				conn.Write([]byte(`client:` + strconv.Itoa(num)))
				var b = make([]byte, 100)
				n, _ := conn.Read(b)
				fmt.Println(string(b[:n]))
				time.Sleep(time.Second)
			}
		}(i)
	}
	time.Sleep(time.Hour)
}
