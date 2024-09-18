package main

import (
	"context"
	"errors"
	"golang.org/x/sync/singleflight"
	"sync"
	"time"
)

func main() {

}

type Call struct {
	wg   sync.WaitGroup
	resp string
	err  error
}

var (
	mu sync.Mutex
	m  = make(map[string]*Call)
)

func Search(ctx context.Context, word string) (string, error) {
	mu.Lock()
	if c, ok := m[word]; ok {
		mu.Unlock()
		c.wg.Wait() // 阻塞等待结果返回
		return c.resp, c.err
	}
	c := new(Call)
	m[word] = c
	c.wg.Add(1)
	mu.Unlock() // 释放锁，相同请求阻塞在wg.Wait中
	// 结果返回
	c.resp, c.err = "response", nil
	c.wg.Done()
	mu.Lock()
	delete(m, word) // 避免map读写并发(sync.Map)
	mu.Unlock()
	return c.resp, c.err
}

var sf = singleflight.Group{}

func SearchSingleFlight(ctx context.Context, word string) (string, error) {
	// 超时处理
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	/*resp, err, _ := sf.Do(word, func() (interface{}, error) {
		val, err := "response", errors.New("")
		if err != nil {
			// 重试
			sf.Forget(word)
			return "", nil
		}
		return val, err
	})*/
	result := sf.DoChan(word, func() (interface{}, error) {
		val, err := "response", errors.New("")
		if err != nil {
			// 重试
			sf.Forget(word)
			return "", nil
		}
		return val, err
	})
	select {
	case r := <-result:
		return r.Val.(string), r.Err
	case <-ctx.Done():
		return "", ctx.Err()
	}
}
