package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":20000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil && err != io.EOF {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil && err != io.EOF {
			fmt.Println("read from client failed, err:", err)
			break
		}
		if err == io.EOF {
			fmt.Println("客户端断开")
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到客户端发来的数据：", recvStr)
		conn.Write([]byte(recvStr))
	}
}
