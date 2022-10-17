package main

import "fmt"

func main() {
	m := map[string]int{"uno": 1}
	p := &m
	fmt.Println(m, p)
	fmt.Printf("%v, %#v\n", p, p)
	// *p = 2
	fmt.Println(m["uno"])

	var m1 map[string]int
	delete(m1, "oh")
	fmt.Println(m1)
}
