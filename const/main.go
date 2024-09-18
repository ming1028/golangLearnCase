package main

import "fmt"

const name = 123

func main() {
	// &name  常量不能寻址
	fmt.Printf("%T\n", name)
	// fmt.Printf("%p\n", &name)
}
