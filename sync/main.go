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
	val, loaded := syncMap.LoadAndDelete("age") // 获取然后删除
	fmt.Println(val, loaded)
	val, loaded = syncMap.LoadOrStore("address", "shanghai") // 获得值 没有则保存
	fmt.Println(val, loaded)
	syncMap.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})
}
