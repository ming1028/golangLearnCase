package main

import (
	mapset "github.com/deckarep/golang-set/v2"
	"math/rand"
	"time"
)

func main() {
	RandIdx(1, 2)
}
func RandIdx(l, need int) []int {
	ret := make([]int, 0, need)
	ms := mapset.NewSet[int]()
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < need; i++ {
		for {
			r := rand.Intn(l)
			if !ms.Contains(r) {
				ms.Add(r)
				break
			}
		}

	}
	ret = ms.ToSlice()
	return ret
}
