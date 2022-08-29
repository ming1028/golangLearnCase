package main

import (
	"encoding/json"
	"fmt"
	"runtime"
)

type People struct {
	name string `json:"name"`
}

func main() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"
	select {
	case val := <-int_chan:
		fmt.Println(val)
	case val := <-string_chan:
		fmt.Println(val)
	}
	js := `{
		"name":"11"
	}`
	var p People // 属性小写
	_ = json.Unmarshal([]byte(js), &p)
	fmt.Println(p)
	fmt.Println(&People{name: "11"} == &People{name: "11"})
	fmt.Println(People{name: "11"} == People{name: "11"})
	str1 := []string{"a", "b", "c"}
	str2 := str1[1:]
	str2[1] = "new"
	fmt.Println(str1)
	fmt.Printf("%p  %p\n", &str1, &str2)
	str2 = append(str2, "x", "y", "z") // 扩容后底层数组发生变化
	fmt.Printf("%p  %p\n", &str1, &str2)
	fmt.Println(str1)
}
