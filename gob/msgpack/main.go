package main

import (
	"fmt"
	"github.com/vmihailenco/msgpack"
)

type Person struct {
	Name   string
	Age    int
	Gender string
}

func main() {
	person1 := Person{
		Name:   "是",
		Age:    22,
		Gender: "难",
	}
	/*p1 := new(Person)
	p1.Age = 33*/
	fmt.Printf("%#v\n", person1)

	person1Pack, _ := msgpack.Marshal(person1)
	fmt.Println(person1Pack)
	var person2 Person
	_ = msgpack.Unmarshal(person1Pack, &person2)
	fmt.Printf("%#v\n", person2)
}
