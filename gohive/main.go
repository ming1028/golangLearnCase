package main

import (
	"fmt"
	"github.com/loveleshsharma/gohive"
	"net"
	"sync"
	"time"
)

var wg sync.WaitGroup

var addressChan = make(chan string, 100)

type works struct{}

func main() {
	begin := time.Now()

	var (
		ip       = "10.68.99.4"
		poolSize = 70000
		pool     = gohive.NewFixedPool(poolSize)
	)
	go func() {
		for port := 1; port <= 65535; port++ {
			address := fmt.Sprintf("%s:%d", ip, port)
			addressChan <- address
		}
		close(addressChan)
	}()

	for work := 9; work < poolSize; work++ {
		wg.Add(1)
		pool.Submit(works{})
	}
	wg.Wait()
	fmt.Println("耗时：", time.Now().Sub(begin))
}

func (w works) Run() {
	defer wg.Done()

	for {
		address, ok := <-addressChan
		if !ok {
			break
		}
		// fmt.Println("address:", address)
		conn, err := net.Dial("tcp", address)
		// conn, err := net.DialTimeout("tcp", address, 10)
		if err != nil {
			fmt.Println("close:", address, err)
			continue
		}
		conn.Close()
		fmt.Println("open:", address)
	}
}
