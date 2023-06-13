package main

import "fmt"

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
}

func (stu *Student) say() {
	fmt.Println(123)
}

type student struct {
	name string
	age  int
}
