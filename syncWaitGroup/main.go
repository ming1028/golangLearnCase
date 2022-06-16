package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"sync"
)

func main() {
	g := errgroup.Group{}
	g.Go(func() error {
		fmt.Println("234")
		return nil
	})
	if g.Wait() != nil {
		fmt.Println("err")
	}

	group := sync.WaitGroup{}
	group.Add(1)
	go func() {
		fmt.Println("888")
		group.Done()
	}()
	group.Wait()
}
