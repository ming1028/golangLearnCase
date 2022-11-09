package main

import (
	"fmt"
	"reflect"
)

func main() {
	s1 := []int{}
	fmt.Println(len(s1), cap(s1), s1 == nil)
	var s2 []int // 声明
	fmt.Println(len(s2), cap(s2), s2 == nil)

	i := 20
	t := T{
		10,
		&i,
	}
	p := &t.x
	fmt.Println(p)
	*p++
	*p--
	t.y = p
	fmt.Println(*t.y)

	x := make([]int, 2, 10)
	_ = x[6:10]
	// _ = x[6:] // [i:j] 省略j j为切片或者数组的长度，并且i<=j
	_ = x[2:]

	var m map[int]bool
	a := m[123]
	fmt.Println(a)
	var ps *[5]string
	fmt.Printf("%#v\n", ps)
	fmt.Println(len(ps), reflect.TypeOf(ps), reflect.ValueOf(ps))
	for range ps {
		fmt.Println(len(ps))
	}

	nil := 123
	fmt.Println(nil)

	ts := [2]Ts{}
	for i, t := range &ts { // for-range 数组，是数组的副本 数组指针，循环变量t是数组元素的副本（结构体值）
		switch i {
		case 0:
			t.n = 3
			ts[1].n = 9
		case 1:
			fmt.Print(t.n, " ")
		}
	}
	fmt.Print(ts)
}

type T struct {
	x int
	y *int
}

type Ts struct {
	n int
}
