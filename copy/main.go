package main

import "fmt"

type T struct {
	A string
	B []string
}

func main() {
	// struct是深拷贝，y拷贝x所有的内容，slice是浅拷贝，y、x中B相等都指向同一片内存地址
	x := T{"煎鱼", []string{"上班"}}

	y := x
	y.A = "咸鱼"
	y.B[0] = "下班"

	fmt.Println(x)
	fmt.Println(y)
}
