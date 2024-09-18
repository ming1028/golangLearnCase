package main

import "fmt"

type HandleFunc func(int, int)

type MiddleFunc func(handleFunc HandleFunc) HandleFunc

func main() {
	ms := []MiddleFunc{f1(), f2()}
	h := func(i, j int) {
		fmt.Println(i + j)
	}
	for _, m := range ms {
		h = m(h)
	}
	h(1, 2)
}

func f1() MiddleFunc {
	return func(next HandleFunc) HandleFunc {
		return func(i, j int) {
			fmt.Println("f1", i, j)
			next(j, i)
		}
	}
}

func f2() MiddleFunc {
	return func(next HandleFunc) HandleFunc {
		return func(i, j int) {
			fmt.Println("f2", i, j)
			next(j, i)
		}
	}
}
