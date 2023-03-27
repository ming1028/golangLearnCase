package main

import (
	"bytes"
	"fmt"
	"html/template"
	"math"
)

var str = `abc,lsdkls,{{.}}`

type name struct {
	Nickname string
	Age      int
}

func main() {
	fmt.Println(math.Ceil(1 / 50))
	tpl, err := template.New("tmpl").
		Delims("{{", "}}").
		Parse(str)
	fmt.Println(err)
	var b bytes.Buffer
	n := name{
		Age:      12,
		Nickname: "王八蛋,",
	}
	err = tpl.Execute(&b, n.Nickname)
	fmt.Println(err)
	fmt.Println(b.String())
}
