package main

import (
	"fmt"
	"github.com/google/uuid"
	"os"
	"sync/atomic"
	"time"
)

func main() {
	var num int64
	fmt.Println(uuid.New().String())
	now := time.Now()
	//日期
	date := now.Format("20060102150405")
	//毫秒
	m := now.UnixNano()/1e6 - now.UnixNano()/1e9*1e3
	ms := sup(m, 3)
	//进程id
	p := os.Getpid() % 1000
	ps := sup(int64(p), 3)
	//随机数
	i := atomic.AddInt64(&num, 1)
	r := i % 10000
	rs := sup(r, 4)
	n := fmt.Sprintf("%s%s%s%s", date, ms, ps, rs)
	fmt.Println(n)
}

func sup(i int64, n int) string {
	m := fmt.Sprintf("%d", i)
	for len(m) < n {
		m = fmt.Sprintf("0%s", m)
	}
	return m
}
