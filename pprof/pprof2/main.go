package main

import (
	"fmt"
	"github.com/arl/statsviz"
	"github.com/shirou/gopsutil/process"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/debug"
	"time"
)

func main() {
	fmt.Println("初始goroutine数量：", runtime.NumGoroutine())

	statsviz.RegisterDefault()

	// 打印GC信息
	go printGCStats()

	// pprof
	go func() {
		fmt.Println(http.ListenAndServe(fmt.Sprintf("127.0.0.1:%s", "6060"), nil))
	}()

	// 打印内存和cpu信息
	go func() {
		pid := os.Getpid()
		fmt.Println("当前程序pid：", pid)

		pro, _ := process.NewProcess(int32(pid))
		for {
			v, _ := pro.CPUPercent()
			if v < 1 {
				continue
			}
			memPercent, _ := pro.MemoryPercent()
			fmt.Printf("该进程的cpu占用率:%v,内存占用:%v, 时间:%v\n", v, memPercent, time.Now().Format(time.RFC3339Nano))
			println("---------------分割线------------------")
			time.Sleep(5 * time.Second)
		}
	}()

	fmt.Printf("最初！程序中goroutine的数量为:%d\n", runtime.NumGoroutine())

	for i := 0; i < 1000000; i++ {
		go func() {
			time.Sleep(time.Second * 10)
		}()
	}
	fmt.Println("for循环结束后，goroutine数量：", runtime.NumGoroutine())

	time.Sleep(5e9)

	fmt.Printf("5s后程序中goroutine的数量为:%d\n", runtime.NumGoroutine())

	time.Sleep(5e9)

	fmt.Printf("10s后程序中goroutine的数量为:%d\n", runtime.NumGoroutine())

	time.Sleep(5e9)
	fmt.Printf("15s后程序中goroutine的数量为:%d\n", runtime.NumGoroutine())

	time.Sleep(5e9)
	fmt.Printf("20s后程序中goroutine的数量为:%d\n", runtime.NumGoroutine())

	time.Sleep(time.Hour)
}

func printGCStats() {
	t := time.NewTicker(time.Second)
	s := debug.GCStats{}
	for {
		select {
		case <-t.C:
			debug.ReadGCStats(&s)
			fmt.Printf("gc %d last@%v, PauseTotal %v\n", s.NumGC, s.LastGC, s.PauseTotal)
		}
	}
}
