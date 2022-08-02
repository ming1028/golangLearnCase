package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover:", err)
		}
	}()
	test()
}

func test() {
	var run func() = nil
	defer run()

	fmt.Println("test")
}
