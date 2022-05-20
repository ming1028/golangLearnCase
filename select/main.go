package main

import "fmt"

// 当ch1和ch2同时达到就绪状态时，优先执行任务1，在没有任务1的时候再去执行任务2呢
func main() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	for {
		select {
		case job1 := <-ch1:
			fmt.Println(job1)
		case job2 := <-ch2:
		priority:
			for {
				select {
				case job1 := <-ch1:
					fmt.Println(job1)
				default:
					break priority
				}
			}
			fmt.Println(job2)
		}
	}
}
