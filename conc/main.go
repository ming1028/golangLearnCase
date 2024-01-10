package main

import (
	"errors"
	"fmt"
	"github.com/sourcegraph/conc/pool"
	"github.com/sourcegraph/conc/stream"
	"strconv"
	"time"
)

func main() {
	// 设置最大协程数量
	poolWithMaxGoroutine()
	// 获得返回值
	poolWithResult()
	// 流式
	poolStream()
}

func poolWithMaxGoroutine() {
	p := pool.New().WithMaxGoroutines(5).WithErrors()
	for i := 1; i < 6; i++ {
		j := i
		p.Go(func() error {
			time.Sleep(time.Second)
			fmt.Println(j)
			if j%2 == 0 {
				// panic(j)
				return errors.New(strconv.Itoa(j))
			}
			return nil
		})
	}
	err := p.Wait()
	fmt.Println(err.Error())
}

func poolWithResult() {
	p := pool.NewWithResults[int]()
	for i := 1; i < 6; i++ {
		j := i
		p.Go(func() int {
			return j * 2
		})
	}
	res := p.Wait()
	fmt.Println(res)
}

func poolStream() {
	times := []int{20, 45, 38, 23, 100}
	s := stream.New()
	for _, millis := range times {
		dur := time.Duration(millis) * time.Millisecond
		s.Go(func() stream.Callback {
			time.Sleep(dur)
			return func() {
				fmt.Println(dur)
			}
		})
	}
	s.Wait()
}
