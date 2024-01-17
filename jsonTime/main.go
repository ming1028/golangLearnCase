package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	p1 := Post{
		CreatedTime: time.Now(),
	}
	p1j, _ := json.Marshal(p1)
	fmt.Printf("p1j:%s\n", p1j)

	p2j := `{"createdTime":"2020-04-05 12:25:42"}`
	var p2 Post
	json.Unmarshal([]byte(p2j), &p2)
	fmt.Printf("p2:%#v\n", p2)
}

type Post struct {
	CreatedTime time.Time `json:"createdTime"`
}
