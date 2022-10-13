package main

import (
	"fmt"
	"reflect"
)

type Direction int

const (
	a = iota
	b = iota
	c
)

const (
	d = "dd"
	e
	f
	g = iota
)

const (
	North Direction = iota
	East
	South
	West
)

// iota 在 const 关键字出现时将被重置为0，const中每新增一行常量声明将使 iota 计数一次。
func main() {
	fmt.Println(a, b, c, e, f, g)
	fmt.Println(reflect.TypeOf(East), reflect.TypeOf(West))
}
