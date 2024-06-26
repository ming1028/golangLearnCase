package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("开始")
	resp, err := http.Get("http://localhost:8088/streamGet")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	reader := bufio.NewReader(resp.Body)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from stream:", err)
			break
		}
		fmt.Print(line)
	}

}
