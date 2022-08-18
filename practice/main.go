package main

import "fmt"

type Student struct {
	Age int
}

func main() {
	kv := map[string]Student{"m": {Age: 22}}
	// kv["m"].Age = 21 map无法取址
	s := []Student{{Age: 21}}
	s[0].Age = 22
	fmt.Println(kv, s)
	s1 := make(map[string]int)
	delete(s1, "h")
	fmt.Println(s1["h"])
	sli := []int{1, 2, 3}
	m := make(map[int]*int)
	for k, v := range sli {
		m[k] = &v
	}
	for k, v := range m {
		fmt.Println(k, "->", *v)
	}

	a, b := 1, 1
	defer func(x, y int) {
		fmt.Println(x + y)
	}(a, b)
	a, b = 2, 2
	defer func() {
		fmt.Println(a + b)
	}()
	a, b = 3, 3
	fmt.Println(a + b)
	var i interface{}
	if i == nil {
		fmt.Println("nil")
		return
	}
	fmt.Println("not nil")
}
