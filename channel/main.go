package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 9; i++ {
		Retry(func() {
			defer fmt.Println(1)
		})
	}

	chan1 := make(chan int, 1)
	stop2(chan1)
	go func() {
		for j := 1; j < 3; j++ {
			jc := make(chan int, 1)
			chanPool(jc, j)
			for i := j * 10; i < (j+1)*10; i++ {
				jc <- i
			}
			close(jc)
			//time.Sleep(time.Second * 5)
		}
	}()

	time.Sleep(time.Minute)
}

// 只能输出
func stop(stop <-chan int) {
	// close(stop)
}

// 输入
func stop2(stop chan<- int) {
	close(stop)
}

func chanPool(jobChan chan int, num int) {
	for i := 0; i < num; i++ {
		go func(n, m int) {
			for msg := range jobChan {
				fmt.Println(msg, n, "协程", m)
			}
			fmt.Println("关闭", n, "协程", m)
		}(num, i)
	}
}

func Retry(f func()) {
	f()
}
