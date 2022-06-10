package main

import "fmt"

type person struct {
	name string
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
	p3.walk()
	p3.eat()

	p4 := &person{
		name: "zzz",
	}
	p4.walk()
	p4.eat()
	fmt.Printf("p1: %T, p2: %T, p3: %T, p4: %T", p1, p2, p3, p4)
}
