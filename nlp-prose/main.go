package main

import (
	"fmt"
	"github.com/jdkato/prose/v2"
)

func main() {
	doc, _ := prose.NewDocument("I'll be right with you.")
	for _, tok := range doc.Tokens() {
		fmt.Println(tok.Text, tok.Tag, tok.Label)
	}
	fmt.Println("======")
	// Iterate over the doc's named-entities:
	for _, ent := range doc.Entities() {
		fmt.Println(ent.Text, ent.Label)
	}
	fmt.Println("======")
	// Iterate over the doc's sentences:
	for _, sent := range doc.Sentences() {
		fmt.Println(sent.Text)
	}
}
