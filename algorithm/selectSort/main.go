package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var arr []int32
	rand.Seed(time.Now().Unix())
	for i := 0; i < 20; i++ {
		arr = append(arr, rand.Int31n(100))
	}
	fmt.Println(arr)
	selectSort(arr)
	fmt.Println(arr)
}

func selectSort(arr []int32) {
	length := len(arr)
	if length <= 1 {
		return
	}

	for i := 0; i < length; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		if i != min {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
}
