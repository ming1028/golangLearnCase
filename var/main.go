package main

import (
	"fmt"
	"reflect"
)

func main() {
	var k = 9
	for k = range []int{} {
		fmt.Println("range", k)
	}
	var m map[int]bool
	// m[123] = false // 声明没有分配内存空间
	_ = m[123]

	var p *[5]string
	for range p {
		_ = len(p)
	}

	var s []int
	// s[0] = 1 声明没有分配内存空间
	_ = s[:]
	// s, s[0] = []int{1, 2}, 9 // 先左边 在右边

	var ss = (*[3]int)(nil)
	fmt.Println(reflect.TypeOf(ss), reflect.ValueOf(ss), reflect.ValueOf(ss).Kind())
	fmt.Printf("%#v\n", ss)
	fmt.Println(len(ss))
	nil := 123
	fmt.Println(nil, reflect.TypeOf(nil))

	var x int8 = -128 // -128~127
	var y = x / -1    // y int8类型 128溢出
	fmt.Println(y)
}
