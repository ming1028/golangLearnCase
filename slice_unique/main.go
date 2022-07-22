package main

import (
	"fmt"
	"sort"
)

func main() {
	sli := []int{12, 3, 65, 12}
	sort.SliceStable(sli, func(i, j int) bool {
		return sli[i] < sli[j]
	})
	fmt.Println(sli)
}
