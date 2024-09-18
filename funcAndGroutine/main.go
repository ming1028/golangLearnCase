package main

import (
	"fmt"
	"time"
)

var values = [5]int{1, 2, 3, 4, 5} // 常量不能用 :=

func main() {
	fmt.Print("闭包 顺序执行：")
	for _, value := range values {
		func() {
			fmt.Print(value, " ")
		}()
	}

	fmt.Println()
	time.Sleep(5e9)
	fmt.Print("协程：")
	// 因为协程可能在循环结束后还没有开始执行，而此时ix值是4
	for _, value := range values {
		go func() {
			fmt.Print(value, " ")
		}()
	}
	time.Sleep(5e9)

	fmt.Print("\n协程形参：")
	for _, value := range values {
		go func(val int) { // 形参栈中
			fmt.Print(val, " ")
		}(value)
	}
	time.Sleep(5e9)

	fmt.Print("\n协程变量赋值：\n")
	for _, value := range values {
		val := value // 变量声明是在循环体内部，所以在每次循环时，这些变量相互之间是不共享的，所以这些变量可以单独的被每个闭包使用。
		fmt.Printf("变量地址：%p\n", &val)
		go func() {
			fmt.Print(val, " ")
			fmt.Printf("%p\n", &val)
		}()
	}
	time.Sleep(5e9)
}
