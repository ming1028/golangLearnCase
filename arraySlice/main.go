package main

import (
	"fmt"
	"github.com/elliotchance/testify-stats/assert"
	"testing"
)

const (
	name = 2
	name2
)

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
	sli2 := sli1 // 同一底层数组
	// sli2 扩充 底层数组改变
	sli2[1] = 988
	fmt.Println(sli1, sli2, arr2)
	sli2 = append(sli2, []int{5, 6, 7, 8, 2, 3, 23, 23, 23, 23}...)
	fmt.Println(sli1, sli2)

	fmt.Printf("%+d\n", -2)
	fmt.Println(name, name2)

	// assert equal
	fmt.Println("slice比较", assert.Equal(&testing.T{}, sli1, sli2, "比较"))
}

func editArr(arr *[5]int) {
	arr[1] = 66
}
