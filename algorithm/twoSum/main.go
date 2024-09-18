package main

import "fmt"

func main() {
	type lsit struct {
		Val int
	}
	nums := []int{2, 7, 11, 15}
	target := 9
	res := twoSum(nums, target)
	fmt.Println(res)
}

func twoSum(nums []int, target int) []int {
	diff := make(map[int]int)
	for i, v := range nums {
		idx, ok := diff[target-v]
		if ok {
			return []int{idx, i}
		}
		diff[v] = i
	}
	return []int{}
}
