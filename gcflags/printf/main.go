package main

import "fmt"

func main() {
	foo()
}

func foo() {
	a := 6666666
	b := 77
	fmt.Printf("a = %d\n", a)
	println("addr of a in foo = ", &a)
	println("addr of b in foo =", &b)
}
