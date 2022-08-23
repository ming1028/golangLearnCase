package main

import "fmt"

// 闭包复制的是原对象的指针
func main() {
	tmp1 := add(10)
	fmt.Println(tmp1(1), tmp1(2))

	tmp2 := add(100)
	fmt.Println(tmp2(1), tmp2(2))

	tmp3, tmp4 := test1(10)
	fmt.Println(tmp3(1), tmp4(2))
}

func add(base int) func(int) int {
	fmt.Printf("base in func (%p)\n", &base)
	return func(i int) int {
		fmt.Printf("func base in closure (%p)\n", &base)
		base += i
		return base
	}
}

func test1(base int) (
	func(int) int,
	func(int) int,
) {
	fmt.Printf("base in func (%p)\n", &base)
	add := func(i int) int {
		fmt.Printf("base in closure1 (%p)\n", &base)
		base += i
		return base
	}

	add2 := func(i int) int {
		fmt.Printf("base in closure2 (%p)\n", &base)
		base -= i
		return base
	}
	return add, add2
}
