package main

import "fmt"

func main() {
	m1 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
	}
	for idx, v := range m1 {
		m1[idx+"12"] = v + 1
	}
	for v := range m1 {
		fmt.Println(v)
	}
	fmt.Println(m1)
}