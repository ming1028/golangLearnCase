package main

import (
	"fmt"
	"sort"
)

func main() {
	sli := []int{12, 3, 65, 12}
	sort.SliceStable(sli, func(i, j int) bool { // 严格模式 保持原有顺序
		return sli[i] < sli[j]
	})
	fmt.Println(sli)
	sort.Slice(sli, func(i, j int) bool {
		return sli[i] > sli[j]
	})
	fmt.Println(sli)
}
