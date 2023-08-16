package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"github.com/spf13/cast"
	"math/rand"
	"sync"
	"time"
)

func main() {
	defer ants.Release()

	runTimes := 1000
	var wg sync.WaitGroup
	syncCalculateSum := func() {
		demoFunc()
		wg.Done()
	}

	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = ants.Submit(syncCalculateSum)
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", ants.Running())
}

func demoFunc() {
	time.Sleep(10 * time.Millisecond)
	randInt := rand.Int()
	fmt.Println(cast.ToString(randInt) + " Hello world")
}
