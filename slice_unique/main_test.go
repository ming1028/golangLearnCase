package main

import (
	"math/rand"
	"sort"
	"testing"
)

func main() {

}

func RemoveDuplicates(userIDs []int64) []int64 {
	processed := make(map[int64]struct{})

	uniqUserIDs := make([]int64, 0)
	for _, userID := range userIDs {
		if _, ok := processed[userID]; ok {
			continue
		}
		uniqUserIDs = append(uniqUserIDs, userID)
		processed[userID] = struct{}{}
	}
	return uniqUserIDs
}

// RemoveDuplicatesInPlace 限制：需要先排序
func RemoveDuplicatesInPlace(userIDs []int64) []int64 {
	if len(userIDs) < 2 {
		return userIDs
	}

	// 升序排序
	sort.SliceStable(userIDs, func(i, j int) bool {
		return userIDs[i] < userIDs[j]
	})

	uniqPointer := 0

	for i := 0; i < len(userIDs); i++ {
		// 与uniqPointer位置数比较，不相等将数值，放在uniqPointer位置下一位置，uniqPointer往后偏移
		if userIDs[i] != userIDs[uniqPointer] {
			uniqPointer++
			userIDs[uniqPointer] = userIDs[i]
		}
	}
	return userIDs[:uniqPointer]
}

func generateSlice() []int64 {
	s := make([]int64, 0, 1)

	for i := 0; i < 1_00_00; i++ {
		s = append(s, rand.Int63())
	}
	return s
}

func BenchmarkRemoveDuplicates(b *testing.B) {
	s := generateSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveDuplicates(s)
	}
}

func BenchmarkRemoveDuplicatesInPlace(b *testing.B) {
	s := generateSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveDuplicatesInPlace(s)
	}
}
