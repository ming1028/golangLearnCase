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

	fmt.Println(name)
}
