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
	BubbleSort(arr)
	fmt.Println(arr)
	BubbleSort2(arr)
	fmt.Println(arr)
	fmt.Println(arr[len(arr)-1])
}

func BubbleSort(arr []int32) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func BubbleSort2(arr []int32) {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
	}
}
