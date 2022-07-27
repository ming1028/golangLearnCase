package main

import (
	"fmt"
	r2 "github.com/juju/ratelimit"
	"go.uber.org/ratelimit"
	"time"
)

// 限流策略
func main() {
	r1 := ratelimit.New(500)

	prev := time.Now()
	for i := 0; i < 10; i++ {
		now := r1.Take()
		fmt.Println(i, now.Sub(prev).Seconds(), now.Second())
		prev = now
	}

	prev = time.Now()
	for i := 0; i < 10; i++ {
		bucket := r2.NewBucket(time.Second, 1)
		if bucket.TakeAvailable(1) > 0 {
			fmt.Println(i, time.Now().Sub(prev).Seconds())
		}
		prev = time.Now()
	}
}
