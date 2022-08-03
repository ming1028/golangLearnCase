package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"time"
)

func main() {
	timer1 := time.NewTimer(time.Second * 2)
	time1 := time.Now()
	fmt.Printf("t1:%v\n", time1)
	time.Sleep(time.Second)
	// timer1.Stop() 定时器暂停
	fmt.Printf("t2:%v\n", <-timer1.C)

	// Ticker
	ticker := time.NewTicker(1 * time.Second)
	i := 0
	group := errgroup.Group{}
	group.Go(func() error {
		for {
			i++
			fmt.Println(<-ticker.C)
			if i == 5 {
				ticker.Stop()
				return nil
			}
		}
	})
	if err := group.Wait(); err != nil {
		log.Fatalln("err:", err)
	}
}
