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
					/**
					 *break语句还可以在语句后面添加标签，
					 *表示退出某个标签对应的代码块，
					 *标签要求必须定义在对应的for、switch和 select的代码块上。
					 */
					break priority
				}
			}
			fmt.Println(job2)
		}
	}
}
