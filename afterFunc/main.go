package main

import (
	"fmt"
	"time"
)

var (
	RetryFlag chan int = make(chan int)
	retryTime []int    = []int{100, 300, 800} // 毫秒
)

func main() {
	retry(func() int {
		fmt.Println("1111")
		time.Sleep(time.Second * 2)
		return 2
	})
	time.Sleep(time.Minute)
}

func retry(f func() int) {
	idx := 0
	for {
		go time.AfterFunc(time.Duration(retryTime[idx])*time.Millisecond, func() {
			RetryFlag <- f()
		})
		err := <-RetryFlag
		if err > 0 {
			return
		}
		fmt.Println(err, "调用")
		if idx == len(retryTime)-1 {
			return
		}
		idx++
	}
}
