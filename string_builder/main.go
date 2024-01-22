package main

import (
	"fmt"
	"strings"
)

func main() {
	var stringBuilder strings.Builder
	stringBuilder.WriteString("hello ")
	stringBuilder.WriteString("world")
	// 直接通过指针转换输出string bytes.Buffer申请一块内存存放string
	fmt.Println(stringBuilder.String())
}
