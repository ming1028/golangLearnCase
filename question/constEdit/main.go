package main

func main() {
	const c = 8 // 常量不能取址
	/*a := &c
	*a = 12
	fmt.Println(*a, c)*/
}

type p struct {
	name string `json:"name"`
	age  int    `json:"age"`
}
