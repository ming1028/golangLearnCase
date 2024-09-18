package main

import (
	"fmt"
	"math"
)

type ints interface {
	Say()
}

type People interface {
	Speak(string) string
}

type Student struct{}

// Speak 指针实现
func (stu *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "sb"
	} else {
		talk = "您好"
	}
	return
}

func main() {
	var peo People = &Student{} // 指针接受者实现了People接口，是People类型
	think := "bitch"
	fmt.Println(peo.Speak(think))
	m := make(map[string]*student)
	stus := []student{
		{name: "pprof.cn", age: 18},
		{name: "测试", age: 23},
		{name: "博客", age: 28},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
	fmt.Println(math.Sqrt(4))

	tmp1 := add(10)
	fmt.Println(tmp1(1), tmp1(2))
	// 此时tmp1和tmp2不是一个实体了
	tmp2 := add(100)
	fmt.Println(tmp2(1), tmp2(2))

	f1, f2 := test01(10)
	// base一直是没有消
	fmt.Println(f1(1), f2(2))
	// 此时base是9
	fmt.Println(f1(3), f2(4))
}

func (stu *Student) say() {
	fmt.Println(123)
}

type student struct {
	name string
	age  int
}

// 外部引用函数参数局部变量
func add(base int) func(int) int {
	return func(i int) int {
		base += i
		return base
	}
}

// 返回2个函数类型的返回值
func test01(base int) (func(int) int, func(int) int) {
	// 定义2个函数，并返回
	// 相加
	add := func(i int) int {
		base += i
		return base
	}
	// 相减
	sub := func(i int) int {
		base -= i
		return base
	}
	// 返回
	return add, sub
}
