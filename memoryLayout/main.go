package main

import (
	"fmt"
	"unsafe"
)

// Foo 结构体是占用一块连续的内存，一个结构体变量的大小是由结构体中的字段决定
// 1、起始的存储地址 必须是 内存对齐边界 的倍数。
//
//	2、整体占用字节数 必须是 内存对齐边界 的倍数。
type Foo struct {
	A int8
	B int8
	C int8
}

type Bar struct {
	x int64 // 8字节 一字节8位
	y *Foo  // 8??
	z bool  // 1
}

type Bar2 struct {
	x int32 // 4
	z bool  // 1
	y *Foo  // 8??
}

type Bar3 struct {
	x int32 // 4
	y Foo   // 3
	z bool  // 1
}

type Bar4 struct {
	m struct{}
	n int8
}

type Bar5 struct {
	n int8
	m struct{}
}

type Bar6 struct {
	n bool   // 1
	m int32  // 4
	d string // 16
	h int64  // 8
}

/**
 *如果结构体或数组类型不包含大小大于零的字段（或元素），则其大小为0。两个不同的0大小变量在内存中可能有相同的地址。
 */
func main() {
	var b6 Bar6
	fmt.Println("Bar6", unsafe.Sizeof(b6), unsafe.Alignof(b6.n), unsafe.Alignof(b6.m),
		unsafe.Alignof(b6.d), unsafe.Sizeof(b6.d), unsafe.Alignof(b6.h), unsafe.Sizeof(b6.h))
	var f Foo
	var in int
	// uintptr 指针占用8字节
	fmt.Println("Foo", unsafe.Sizeof(f), unsafe.Sizeof(&f), unsafe.Sizeof(&in))

	var b Bar
	fmt.Println("Bar", unsafe.Sizeof(b))

	var b2 Bar2
	fmt.Println("Bar2", unsafe.Sizeof(b2))

	var b3 Bar3
	fmt.Println("Bar3", unsafe.Sizeof(b3))

	/**
	 *由于空结构体struct{}的大小为 0，所以当一个结构体中包含空结构体类型的字段时，通常不需要进行内存对齐
	 */
	var b4 Bar4
	fmt.Println("Bar4", unsafe.Sizeof(b4))

	/**
	 *当空结构体类型作为结构体的最后一个字段时，如果有指向该字段的指针，
	 *那么就会返回该结构体之外的地址。为了避免内存泄露会额外进行一次内存对齐。
	 */
	var b5 Bar5
	fmt.Println("Bar5", unsafe.Sizeof(b5))

	/**
	 *将常用字段放置在结构体的第一个位置上减少CPU要执行的指令数量，从而达到更快的访问效果。
	 *
	 */
}
