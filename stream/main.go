package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8088/streamGet", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	fmt.Println("开始")
	client := &http.Client{}
	resp, err := client.Do(req)
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
