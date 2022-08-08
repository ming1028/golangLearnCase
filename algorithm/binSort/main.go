package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	var arr []int32
	rand.Seed(time.Now().Unix())
	for i := 0; i < 20; i++ {
		arr = append(arr, rand.Int31n(100))
	}
	fmt.Println(arr)
	sort.Slice(arr, func(i, j int) bool {
		if arr[i] < arr[j] {
			return true
		}
		return false
	})
	findData := rand.Intn(20)
	fmt.Println(arr, arr[findData])
	fmt.Println(binSort(arr[findData], arr))
}

func binSort(find int32, arr []int32) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		fmt.Println(mid)
		if arr[mid] > find {
			high = mid - 1
		} else if arr[mid] < find {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
