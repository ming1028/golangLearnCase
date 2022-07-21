package main

import "fmt"

func main() {
	arr1 := new([5]int) // new 值类型 make引用类型
	fmt.Println(arr1)
	editArr(arr1)
	fmt.Println(arr1)

	arr2 := [...]int{4, 4, 5}
	fmt.Println(arr2)

	sli1 := arr2[:len(arr2)]
	fmt.Println(sli1)
	sli1 = sli1[:len(arr2)-1]
	fmt.Println(sli1)
}

func editArr(arr *[5]int) {
	arr[1] = 66
}
