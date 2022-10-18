package main

func main() {
	chan1 := make(chan int, 1)
	stop2(chan1)
}

// 只能输出
func stop(stop <-chan int) {
	// close(stop)
}

// 输入
func stop2(stop chan<- int) {
	close(stop)
}
