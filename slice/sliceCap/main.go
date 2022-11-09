package main

import (
	"fmt"
	"unsafe"
)

func main() {
	tmp := make([]int, 2, 5)
	tmp[0] = 1
	tmp[1] = 2
	fmt.Printf("%v %d %d %p\n", tmp, len(tmp), cap(tmp), tmp)
	change(tmp)
	fmt.Printf("调用函数后：%v %d %d %p\n", tmp, len(tmp), cap(tmp), tmp)

	p := unsafe.Pointer(&tmp[1])
	q := uintptr(p) + 8
	t := (*int)(unsafe.Pointer(q))
	fmt.Println(*t)
}

func change(tmp []int) {
	// 实参和形参内存空间中的len和cap是独立的
	fmt.Printf("append之前：%p\n", tmp)
	tmp = append(tmp, 3) // append返回新的slice覆盖原slice
	fmt.Printf("append之后：%v %d %d %p\n", tmp, len(tmp), cap(tmp), tmp)
}
