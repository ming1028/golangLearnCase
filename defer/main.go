package main

import (
	"fmt"
	"time"
)

// recover 当前goroutine是否有panic行为
// defer必须在panic前面
func main() {
	/*defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover:", err)
		}
	}()*/
	// test()
	/*m1 := make(map[string]int)
	m2 := make(map[string]int)
	fmt.Println(m2 == m1)*/
	test2()
	fmt.Println("===================")
	test3()
	start := time.Now()
	time.Sleep(3 * time.Second)
	defer func() {
		fmt.Println(time.Now().Sub(start))
	}()
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

	defer fmt.Println(2, recover()) // revocer已执行

	defer func() {
		// 函数调用栈 recover会判断是否在同一个goroutine、是否panic、函数是否退出，是否已经被修复（已修复总返回nil）
		// 当前参数和当前goroutine的函数指针判断
		func() {
			fmt.Println("3", recover()) // ??为什么defer中的闭包不能recover
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
