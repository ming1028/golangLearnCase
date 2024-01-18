package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// 00001100 00000100 00000001 ｜ 按位或
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	fmt.Println(log.Flags())
	log.SetPrefix("[log]")
	logFile, _ := os.OpenFile("./log/xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	log.SetOutput(logFile)
	log.Println("日志输出")
}
