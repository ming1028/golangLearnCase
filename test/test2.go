package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"unicode"
)

func SplitWords(s string) []string {
	return strings.FieldsFunc(s, func(r rune) bool {
		return unicode.IsSpace(r)
	})
}

func main() {
	s := SplitWords(" ")
	fmt.Println(s)
	return
	url := "http://127.0.0.1:8080/v1/word/mark"
	method := "POST"

	payload := strings.NewReader(`{"word": "the", "state": 2, "source": "sentence", "sentence": "I think I did the right thing."}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IiIsInVzZXJJZCI6MSwiaXNzIjoibGFuZ3VhZ2VSZWFjdG9yIiwiZXhwIjoxNzE0MDQyMDc4fQ.DvtPGh6LdtIbVTzse5pKCMuwA6GlJ-4mIqx_Eg9YhW8")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
