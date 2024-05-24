package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	data := `line 1
line 2
line 3
`

	// 创建一个字符串读取器
	sr := strings.NewReader(data)

	// 创建一个新的 Scanner
	scanner := bufio.NewScanner(sr)

	// 使用 Scanner 逐行读取数据
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error:", err)
	}
}
