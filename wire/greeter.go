package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Message string

type Greeter struct {
	Message Message // <- adding a Message field
}

type Event struct {
	Greeter Greeter // <- adding a Greeter field
}

func NewMessage(ctx *gin.Context) Message {
	return Message("Hi there!")
}

func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

func (g Greeter) Greet() Message {
	return g.Message
}

func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}
