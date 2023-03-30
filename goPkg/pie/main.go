package main

import (
	"fmt"
	"github.com/elliotchance/pie/pie"
	"strings"
)

func main() {
	var names pie.Strings
	names = []string{"Bob", "Sally", "John", "Jane"}
	name := names.FilterNot(func(s string) bool {
		return strings.HasPrefix(s, "J")
	}).Map(strings.ToUpper).Last()

	fmt.Println(name, names)

	var abs pie.Ints
	abs = []int{1, 3, -1, -33}
	abs = abs.Abs()
	fmt.Println(abs)

	// 是否全都满足callback
	t := abs.All(func(value int) bool {
		return value > 0
	})
	fmt.Println(t)

	// 任意一个满足
	t = abs.Any(func(value int) bool {
		return value > 10
	})
	fmt.Println(t)

	aa := []int{12, 33}
	var aaa pie.Ints
	aaa = aa
	t = aaa.Any(func(value int) bool {
		return value > 10
	})
	fmt.Println(t)

	aver := aaa.Average()
	fmt.Println(aver)

	bottom := abs.Bottom(2)
	fmt.Println(bottom)

	fmt.Println(abs.Contains(44))
	fmt.Println(abs)
	// 过滤相同元素
	sli1, sli2 := abs.Diff(pie.Ints{1, 55})
	fmt.Println(sli1, sli2)

	fmt.Println(abs.DropTop(3), abs.DropTop(44))

	ss1 := abs.DropWhile(func(s int) bool {
		return s < 2
	})
	fmt.Println(ss1)

	fmt.Println(abs, aaa)
	// 交集
	inter := abs.Intersect(aaa)
	fmt.Println(inter)
	// 连接
	fmt.Println(abs.Join(","))

	// 相邻值操作
	fmt.Println(abs.Reduce(func(i int, i2 int) int {
		fmt.Println(i, i2)
		return i - i2
	}))
}
