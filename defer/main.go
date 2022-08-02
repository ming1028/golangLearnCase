package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover:", err)
		}
	}()
	// test()
	test2()
	fmt.Println("===================")
	test3()
}

func test() {
	var run func() = nil
	defer run()

	fmt.Println("test")
}

// 捕获函数recover只有在延迟调用内直接调用才会终止错误，否则总是返回nil
func test2() {
	defer func() {
		fmt.Println("1", recover(), "1-1")
	}()

	defer fmt.Println(2, recover())

	defer func() {
		func() {
			fmt.Println("3")
			recover()
		}()
	}()
	panic("test pannic")
}

// 延迟调用的引发的错误，可被后续延迟调用捕获，但仅最后一个错误可被捕获
func test3() {
	defer func() {
		fmt.Println("1", recover())
	}()

	defer func() {
		panic("defer panic")
	}()
	panic("test panic")
}
