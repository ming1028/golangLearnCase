package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var arr []int32
	for i := 0; i < 20; i++ {
		arr = append(arr, rand.Int31n(100))
	}
	fmt.Println(arr)
	fmt.Println(QuickSort(arr))
}

func QuickSort(arr []int32) []int32 {
	if len(arr) < 1 {
		return arr
	}
	firstData := arr[0]
	left := make([]int32, 0, len(arr))
	right := make([]int32, 0, len(arr))
	for i := 1; i < len(arr); i++ {
		if arr[i] < firstData {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}
	left, right = QuickSort(left), QuickSort(right)
	return append(append(left, firstData), right...)
}
