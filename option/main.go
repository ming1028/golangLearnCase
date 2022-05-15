package main

import "fmt"

type Option struct {
	A string
	B string
	C int
}

type OptionFunc func(*Option)

var defaultOption = &Option{
	A: "a",
	B: "b",
	C: 1,
}

func main() {
	x := newOption("张三", "李四", 5)
	fmt.Println(x)
	x = newOption2()
	fmt.Println(x)
	x = newOption2(WithA("王五"), WithB("赵四"))
	fmt.Println(x)
}

func newOption2(opts ...OptionFunc) (opt *Option) {
	opt = defaultOption
	for _, o := range opts {
		// o是 withA("").... 返回一个参数为option指针的OptionFunc,
		o(opt)
	}
	return
}

func newOption(a, b string, c int) *Option {
	return &Option{
		A: a,
		B: b,
		C: c,
	}
}

func WithA(a string) OptionFunc {
	return func(option *Option) {
		option.A = a
	}
}

func WithB(b string) OptionFunc {
	return func(option *Option) {
		option.B = b
	}
}

func WithC(c int) OptionFunc {
	return func(option *Option) {
		option.C = c
	}
}
