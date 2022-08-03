package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
			// time.Sleep(time.Second)
		}
	}("world")

	for i := 0; i < 2; i++ {
		fmt.Println("hello")
		runtime.Gosched()
	}
	runtime.GOMAXPROCS(2)
	go a()
	go b()
	time.Sleep(time.Second)
}

func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}
