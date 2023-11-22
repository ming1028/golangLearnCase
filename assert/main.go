package main

import "fmt"

func main() {
	var n interface{} = 55
	fmt.Println(n.(int))
	switch n.(type) {
	case int:
		fmt.Println("int")
	}
}
