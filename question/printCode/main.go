package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
)

// 交替打印数字字母
func main() {
	letter, number := make(chan struct{}), make(chan struct{})
	g := new(errgroup.Group)
	g.Go(func() error {
		i := 1
		for {
			select {
			case <-number:
				if i >= 28 {
					fmt.Println("数字退出")
					return nil
				}
				fmt.Print(i, " ")
				i++
				fmt.Print(i, " ")
				i++
				letter <- struct{}{}
			}
		}
	})

	g.Go(func() error {
		i := 'A'
		for {
			select {
			case <-letter:
				if i >= 'Z' {
					fmt.Println("字母退出")
					number <- struct{}{}
					return nil
				}
				fmt.Print(string(i), " ")
				i++
				fmt.Print(string(i), " ")
				i++
			}
			number <- struct{}{}
		}
	})
	number <- struct{}{}
	if err := g.Wait(); err != nil {
		fmt.Println("err")
	}
}
