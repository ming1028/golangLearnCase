package main

import (
	"fmt"
	"time"
)

func main() {
	ec, rc, done := make(chan error, 1), make(chan interface{}), make(chan struct{})
	go ResponseHandle(rc, ec, done)

	rc <- 1
	time.Sleep(time.Second)
	close(rc)
	time.Sleep(time.Hour)
	return
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

func ResponseHandle(
	ch <-chan interface{},
	ec <-chan error,
	done <-chan struct{},
) {
	defer func() {
		fmt.Println("response return")
	}()
	for {
		select {
		case d, ok := <-ch:
			if !ok {
				fmt.Println("ch 关闭")
				return
			}
			fmt.Println(d)
		case e, ok := <-ec:
			if !ok {
				fmt.Println("ec 关闭")
				return
			}
			fmt.Println(e)
		case <-done:
			return
		}
	}
}
