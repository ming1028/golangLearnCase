package main

import (
	"fmt"
	"reflect"
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
	// map 不可比较
	/*m1 := make(map[string]int)
	m2 := make(map[string]int)
	fmt.Println(m2 == m1)*/
	// slice不可比较
	s1 := make([]int, 0, 1)
	s2 := make([]int, 0, 1)
	fmt.Println(reflect.DeepEqual(s1, s2))
	test2()
	fmt.Println("===================")
	test3()
	start := time.Now()
	defer fmt.Println(time.Now().Sub(start))
	time.Sleep(3 * time.Second)

	fmt.Println("========")
	fmt.Println(f(3))
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

	defer fmt.Println(2, recover()) //作为形参已经执行

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

func f(n int) (r int) {
	defer func() {
		r += n
		recover()
	}()

	var f func() // 声明未定义
	defer f()    // panic

	f = func() {
		r += 2
	}
	return n + 1
}
