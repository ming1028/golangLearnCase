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

	s1 := s{
		Data: map[string]interface{}{},
	}
	s1.Data["count"] = int32(12)
	s1MsgPack, _ := msgpack.Marshal(s1)
	fmt.Println(s1MsgPack)
	s2 := new(s)
	s2.Data = make(map[string]interface{})
	_ = msgpack.Unmarshal(s1MsgPack, s2)
	for i, v := range s2.Data {
		fmt.Printf("%s %T %v\n", i, v, v)
	}
}

type s struct {
	Data map[string]interface{}
}
