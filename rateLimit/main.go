package main

import (
	"fmt"
	"go.uber.org/ratelimit"
	"time"
)

// 限流策略
func main() {
	r1 := ratelimit.New(100)

	prev := time.Now()
	for i := 0; i < 10; i++ {
		now := r1.Take()
		fmt.Println(i, now.Sub(prev))
		prev = now
	}
}
