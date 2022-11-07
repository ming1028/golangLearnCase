package main

import "fmt"

func main() {
	s1 := []int{}
	fmt.Println(len(s1), cap(s1), s1 == nil)
	var s2 []int // 声明
	fmt.Println(len(s2), cap(s2), s2 == nil)
}
