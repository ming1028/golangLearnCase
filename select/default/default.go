package main

import "fmt"

func main() {
	ch := make(chan int)
	for {
		select {
		case <-ch:
			fmt.Println("case1")
		default:
			fmt.Println("default")
		}
	}
}
