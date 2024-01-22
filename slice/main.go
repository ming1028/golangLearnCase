package main

import (
	"fmt"
)

func main() {
	var sli1 []int // nil slice 可以append
	fmt.Println(sli1, sli1 == nil, &sli1)
	fmt.Printf("nil切片地址：%p\n", &sli1)
	sli1 = append(sli1, 1)
	fmt.Println(sli1)
	sli2 := make([]int, 0) // 内存已分配
	fmt.Printf("空切片内存地址：%p\n", &sli2)
	fmt.Println(sli2, sli2 == nil)

	sli1 = []int{1, 2, 3, 4, 5, 6}
	sli2 = []int{11, 21, 31, 41, 51, 6, 7, 8}
	n := copy(sli1, sli2) // sli2 => sli1
	fmt.Println("23333", n, sli2, sli1)
	sli3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	sli4 := make([]int, 6)
	n = copy(sli4, sli3)
	fmt.Println(n, sli4, sli3)

	arr := [5]int{}
	sli := arr[0:3]
	// sli := make([]int, 3, 5) // 0 0 0
	sli[0] = 1
	sli[1] = 2
	// sli: 1 2 0
	change(sli...) // 第一次append之后 底层数组为：1 2 0 3 未发生扩容
	fmt.Println(sli)
	sli2 = sli[0:2]
	change(sli2...)
	fmt.Println(sli2, sli)
	x := make([]int, 2, 10)
	_ = x[6:10]
	// _ = x[6:]
	_ = x[2:]
}

func change(s ...int) {
	fmt.Println("函数内：", len(s), cap(s))
	s = append(s, 3)
}
