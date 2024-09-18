package main

import (
	"fmt"
	"github.com/beefsack/go-rate"
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

	// 阻塞
	rl := rate.New(3, time.Second) // 3 times per second
	begin := time.Now()
	for i := 1; i <= 10; i++ {
		rl.Wait()
		fmt.Printf("%d started at %s\n", i, time.Now().Sub(begin))
	}

	rl = rate.New(2, time.Second)
	for i := 1; i <= 5; i++ {
		if ok, remaining := rl.Try(); ok {
			fmt.Printf("通过\n")
		} else {
			fmt.Printf("阻塞, please wait %s\n", remaining)
			time.Sleep(time.Millisecond * 500)
		}
	}
}
