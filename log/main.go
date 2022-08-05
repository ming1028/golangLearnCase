package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	fmt.Println(log.Flags())
	log.SetPrefix("[log]")
	logFile, _ := os.OpenFile("./log/xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	log.SetOutput(logFile)
	log.Println("日志输出")
}
