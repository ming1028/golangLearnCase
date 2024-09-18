package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// 只读方式打开文件
	file, _ := os.Open("./os/main.go")
	defer file.Close()

	var content []byte
	var buf [128]byte
	for {
		n, err := file.Read(buf[:])
		fmt.Println(n)
		if err == io.EOF {
			break
		}
		content = append(content, buf[:n]...)
	}
	// fmt.Println(string(content))
	file2, _ := os.OpenFile("./os/xx.log", os.O_CREATE|os.O_WRONLY, 0666)
	defer file2.Close()

	writer := bufio.NewWriter(file2)
	for i := 0; i < 10; i++ {
		writer.WriteString("hello\n")
	}
	// 刷新缓冲区
	writer.Flush()

	file3, _ := os.Open("./os/xx.log")
	defer file3.Close()

	reader := bufio.NewReader(file3)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(string(line))
	}
}
