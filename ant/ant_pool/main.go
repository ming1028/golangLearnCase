package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
	"sync/atomic"
)

var sum int32

func main() {
	runTimes := 1000
	var wg sync.WaitGroup

	p, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
		myFunc(i.(int32))
		wg.Done()
	})
	defer p.Release()

	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = p.Invoke(int32(i))
	}

	wg.Wait()
	fmt.Printf("running goroutines: %d\n", p.Running())
}

func myFunc(i int32) {
	n := i
	atomic.AddInt32(&sum, n)
	fmt.Printf("run with %d\n", n)
}
