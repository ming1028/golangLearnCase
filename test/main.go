package main

import (
	"fmt"
	"strings"
)

func main() {
	strings.Contains("abc", " ")
	m := 1
	m <<= 1 // 向左位移赋值  >> 向右位移
	fmt.Println("位移", m)
	sli := []int{1, 2, 3, 4}
	fmt.Println(cap(sli), len(sli))
	map1 := map[string]string{
		"name": "张三",
		"addr": "上海",
	}
	map2 := make(map[string]string, 10)
	fmt.Println(len(map2), len(map1))

	fmt.Printf("%f, %T\n", 5.0/3.0, 5.0/3.0)

	var s1 []int
	if s1 == nil {
		fmt.Println("s1 is nil")
	}
	s2 := make([]int, 0, 0)
	if s2 == nil {
		fmt.Println("s2 is nil")
	}

	var a int
	var intPtr *int
	intPtr = &a
	*intPtr = 20
	fmt.Println(a)

	var p *string
	fmt.Println(p)
	fmt.Printf("p的值是%v， 内存地址%p\n", p, p)
	if p == nil {
		fmt.Println("p is nil")
	}
	p = new(string)
	fmt.Printf("p的值是%v， 内存地址%p\n", p, p)
	*p = "引用类型声明不会分配内存"
	var map3 map[string]string
	map3["引用类型声明"] = "不会分配内存"
	fmt.Println(map3)
}

func init() {
	fmt.Println(1)
}
func init() {
	fmt.Println(2)
}
func init() {
	fmt.Println(3)
}
func init() {
	fmt.Println(5)
}
func init() {
	fmt.Println(4)
}
