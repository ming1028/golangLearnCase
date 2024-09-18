package main

import (
	"fmt"
	"strings"
)

func main() {
	var stringBuilder strings.Builder
	stringBuilder.Reset() // 重置
	stringBuilder.WriteString("hello ")
	stringBuilder.WriteString("world")
	for _, r := range `bbaa` {
		stringBuilder.WriteRune(r)
	}
	// 直接通过指针转换输出string bytes.Buffer申请一块内存存放string
	fmt.Println(stringBuilder.String())
}
