package main

import (
	"log"
	"runtime/debug"
	"time"
)

func main() {
	defer recovery()
	for i := 0; i < 10; i++ {
		go go1()
	}
	time.Sleep(time.Hour)
}

func recovery() {
	log.Println("recover")
	if err := recover(); err != nil {
		log.Println("[Recovery from panic]", "Error", err, "stack", string(debug.Stack()))
		/*xzap.Info("[Recovery from panic]",
			zap.Any("Error", err),
			zap.Any("stack", debug.Stack()),
		)*/
	}
}

func protect(g func()) {
	defer func() {
		log.Println("panic")
		if x := recover(); x != nil {
			log.Printf("run time panic: %v", x)
		}
	}()
}

func go1() {
	go go2()
}

func go2() {
	defer recovery()
	panic("err")
}
