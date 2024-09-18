package main

import "fmt"

func main() {
	sli3 := []int{1, 2, 3}
	var res []int
	sliTemp := sli3[1:]
	res = append(res, sli3[0:1]...)
	res = append(res, 4)
	res = append(res, sliTemp...)
	sli3 = res
	fmt.Println(sli3)

	fmt.Println(append([]int{4}, sli3...))
}
