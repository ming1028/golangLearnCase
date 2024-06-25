package main

import (
	mapset "github.com/deckarep/golang-set/v2"
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	// RandIdx(1, 2)
	res, err := http.Get("https://36kr.com/p/2815034915392002")
	if err != nil {
		return
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}
	err = os.MkdirAll("./uploads/html/", 0755)
	if err != nil {
		return
	}
	err = os.WriteFile("./uploads/123.html", bodyBytes, 0644)
	if err != nil {
		return
	}
	RandIdx(2, 5)
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
