package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 10
	fmt.Println(reflect.TypeOf(i))
}

func classifier(items ...interface{}) {
	for _, x := range items {
		switch x.(type) { // 只能用在switch中
		case bool:
			fmt.Println("bool")
		case float64:
			fmt.Println("float64")
		}
	}
}
