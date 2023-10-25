package main

import (
	"fmt"
	"github.com/spf13/cast"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 1000; i++ {
		fmt.Println(rand.Int63n(8) + 3)
	}
	for i := 1; i <= 3; i++ {
		buildData(i)
	}
	time.Sleep(time.Hour)
}

func buildData(i int) {
	msgChan := make(chan int, 100)
	msgPool(i, 5, msgChan)
	for i := 1; i < 5; i++ {
		msgChan <- i
	}
	close(msgChan)
	fmt.Println("任务号："+cast.ToString(i), "结束")
}

func msgPool(j, num int, msgChan chan int) {
	for i := 0; i < num; i++ {
		chanNum := i
		bNum := j
		go func(bNum, chanNum int, msgChan chan int) {
			for msg := range msgChan {
				fmt.Println(msg, "队列号", chanNum, "任务号：", bNum)
				time.Sleep(time.Second * 2)
			}
			fmt.Println("任务号："+cast.ToString(bNum), "队列号："+cast.ToString(chanNum), "结束")
		}(bNum, chanNum, msgChan)
	}
}
