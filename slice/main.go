package main

import (
	"fmt"
)

func main() {
	var sli1 []int
	fmt.Println(sli1, sli1 == nil, &sli1)
	fmt.Printf("nil切片地址：%p\n", &sli1)
	sli1 = append(sli1, 1)
	fmt.Println(sli1)
	sli2 := make([]int, 0)
	fmt.Printf("空切片内存地址：%p\n", &sli2)
	fmt.Println(sli2, sli2 == nil)

	sli1 = []int{1, 2, 3, 4, 5, 6}
	sli2 = make([]int, 12)
	n := copy(sli1, sli2) // sli2 => sli1
	fmt.Println(n, sli2, sli1)
	sli3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	sli4 := make([]int, 6)
	n = copy(sli4, sli3)
	fmt.Println(n, sli4, sli3)

	sli := make([]int, 3, 5)
	sli[0] = 1
	sli[1] = 2
	change(sli...)
	fmt.Println(sli)
	sli2 = sli[0:2]
	change(sli2...)
	fmt.Println(sli)
}

func change(s ...int) {
	fmt.Println("函数内：", len(s), cap(s))
	s = append(s, 3)
}
