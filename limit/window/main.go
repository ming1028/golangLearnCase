package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var (
	QPS        int64 = 2
	TimeWindow int64 = 1000
	TimeCnt    int64
	TimeStart  = time.Now().UnixMilli()
)

func main() {
	for i := 0; i < 20; i++ {
		go func() {
			timeNow := time.Now().UnixMilli()
			time.Sleep(time.Millisecond * 200)
			if timeNow-TimeStart > TimeWindow {
				atomic.StoreInt64(&TimeCnt, 0)
				TimeStart = time.Now().UnixMilli()
			}
			atomic.AddInt64(&TimeCnt, 1)
			if atomic.LoadInt64(&TimeCnt) >= QPS {
				fmt.Println("限流了")
			} else {
				fmt.Println("放行")
			}
		}()
	}
	time.Sleep(time.Hour)
}
