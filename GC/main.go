package main

import (
	"os"
	"runtime/trace"
)

func main() {
	f, _ := os.Create("trace.out")
	defer f.Close()
	trace.Start(f)
	defer trace.Stop()
	for n := 1; n < 100000; n++ {
		allocate()
	}
}

func allocate() {
	_ = make([]byte, 1<<20)
}
