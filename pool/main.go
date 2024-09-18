package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type Pool struct {
	TaskChannel chan func()
}

func NewPool(cap ...int) *Pool {
	var n int
	if len(cap) > 0 {
		n = len(cap)
	}
	if n == 0 {
		n = runtime.NumCPU()
	}
	p := &Pool{
		TaskChannel: make(chan func()),
	}

	for i := 0; i < n; i++ {
		go func() {
			for task := range p.TaskChannel {
				task()
			}
		}()
	}
	return p
}

func (p *Pool) Submit(f func()) {
	p.TaskChannel <- f
}

func main() {
	p := NewPool()
	var wg sync.WaitGroup
	wg.Add(3)
	task1 := func() {
		fmt.Println("eat cost 3 seconds")
		time.Sleep(3 * time.Second)
		wg.Done()
	}
	task2 := func() {
		fmt.Println("wash feet cost 3 seconds")
		time.Sleep(3 * time.Second)
		wg.Done()
	}
	task3 := func() {
		fmt.Println("watch tv cost 3 seconds")
		time.Sleep(3 * time.Second)
		wg.Done()
	}

	p.Submit(task3)
	p.Submit(task2)
	p.Submit(task1)
	wg.Wait()
}
