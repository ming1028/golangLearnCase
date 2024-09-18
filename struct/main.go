package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	name string
}

type person3 struct {
	Name string
}

func (p person) walk() {
	fmt.Println(p.name, " walk")
}

func (p *person) eat() {
	fmt.Println(p.name, " eat")
}
func main() {
	var p1 person
	p1.walk()
	p1.eat()

	p2 := new(person)
	p2.walk()
	p2.eat()

	p3 := person{
		name: "sss",
	}
	p3NoTag := person3{
		Name: "noJsonTag",
	}
	p3Json, _ := json.Marshal(p3NoTag)
	fmt.Println("没有tagJson:", string(p3Json))
	p3.walk()
	p3.eat()

	p4 := &person{
		name: "zzz",
	}
	p4.walk()
	p4.eat()
	fmt.Printf("p1: %T, p2: %T, p3: %T, p4: %T\n", p1, p2, p3, p4)

	// 匿名结构体
	var user struct {
		Name string
		Age  int
	}
	user.Name = "sss"
	user.Age = 33
	fmt.Printf("%#v\n", user)
	var fra Fragment = new(GetPodAction)
	var fra2 Fragment = &GetPodAction{}
	var fra3 Fragment = GetPodAction{}
	fmt.Println(fra3, fra2, fra)
}

type Fragment interface {
	Exec(transInfo *person) error
}
type GetPodAction struct {
}

func (g GetPodAction) Exec(transInfo *person) error {
	return nil
}
