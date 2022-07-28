package main

import "fmt"

func main() {
	sli := []int{1, 2, 3, 4}
	fmt.Println(cap(sli), len(sli))
	map1 := map[string]string{
		"name": "张三",
		"addr": "上海",
	}
	map2 := make(map[string]string, 10)
	fmt.Println(len(map2), len(map1))

	fmt.Println(5.0 / 3.0)

	var s1 []int
	if s1 == nil {
		fmt.Println("s1 is nil")
	}
	s2 := make([]int, 0, 0)
	if s2 == nil {
		fmt.Println("s2 is nil")
	}

}

func init() {
	fmt.Println(1)
}
func init() {
	fmt.Println(2)
}
func init() {
	fmt.Println(3)
}
func init() {
	fmt.Println(5)
}
func init() {
	fmt.Println(4)
}
