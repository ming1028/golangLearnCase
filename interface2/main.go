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
}

func (stu *Student) say() {
	fmt.Println(123)
}
