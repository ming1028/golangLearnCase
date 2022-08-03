package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	x      int64
	wg     sync.WaitGroup
	lock   sync.Mutex
	rwlock sync.RWMutex
)

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock()
		x += 1
		lock.Unlock()
	}
	wg.Done()
}

func atomicAdd() {
	atomic.AddInt64(&x, 1)
	wg.Done()
}
