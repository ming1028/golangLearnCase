package main

import (
	"fmt"
	"github.com/spf13/cast"
	"time"
)

func main() {
	for i := 1; i < 2; i++ {
		go Init(i)
	}
	time.Sleep(time.Minute * 30)
}

func createPool(
	num int,
	jobChan chan int,
) {
	for i := 0; i < 50; i++ {
		chanNum := i
		go func(
			jobChan chan int,
			chanNum int,
			proNum int,
		) {
			// 消费
			for job := range jobChan {
				fmt.Println("处理程序"+cast.ToString(num),
					"协程数：", chanNum, job)
				time.Sleep(time.Microsecond * 500)
			}
		}(jobChan, chanNum, num)
	}
}

func Init(proNum int) {
	jobChan := make(chan int, 200)
	createPool(proNum, jobChan)
	for i := 0; i < proNum*1000; i++ {
		jobChan <- i
		time.Sleep(time.Microsecond * 500)
	}
}
