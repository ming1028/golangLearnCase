package main

import (
	"fmt"
	"sync"
)

const N = 10

var wg = &sync.WaitGroup{}

func main() {
	for i := 0; i < N; i++ {
		// i := i 复用i 大概率出现都是N的结果
		wg.Add(1)
		go func() {
			// wg.Add(1) 还未执行 wg.Wait()已执行 main函数退出
			fmt.Println(i)
			defer wg.Done()
		}()
	}
	wg.Wait()
}
