package main

import (
	"fmt"
	"sync"
)

var syncMap = sync.Map{}

func main() {
	syncMap.Store("name", "张三")
	syncMap.Store("age", 20)
	val, ok := syncMap.Load("name")
	fmt.Println(val, ok)
	syncMap.Delete("name")
	val, loaded := syncMap.LoadAndDelete("age")
	fmt.Println(val, loaded)
	syncMap.LoadOrStore("address", "shanghai")
	syncMap.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})
}
