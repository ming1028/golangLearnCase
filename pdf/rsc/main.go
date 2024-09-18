package main

import (
	"fmt"
	"rsc.io/pdf"
)

func main() {
	file, err := pdf.Open("./colly/files/2.pdf")
	if err != nil {
		panic(err)
	}
	fmt.Println(file.NumPage())
	content := file.Page(file.NumPage()).Content().Text
	text := ""
	for _, v := range content {
		text += v.S
	}
	fmt.Println(text)
}
