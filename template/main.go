package main

import (
	"bytes"
	"fmt"
	"html/template"
	"math"
)

// .整个对象  .field 结构体具体值，渲染时要传结构题
var str = "{{.}}"

type name struct {
	Nickname string
	Age      int
}

func main() {
	fmt.Println(math.Ceil(1 / 50))
	tpl, err := template.New("tmpl").Parse(str)
	fmt.Println(err)
	var b bytes.Buffer
	n := name{
		Age:      12,
		Nickname: "333",
	}
	err = tpl.Execute(&b, n.Nickname)
	fmt.Println(err)
	fmt.Println(b.String())
}
