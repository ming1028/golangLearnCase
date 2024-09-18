package main

import (
	"fmt"
	v8 "rogchap.com/v8go"
)

func main() {
	ctx := v8.NewContext()                                          // creates a new V8 context with a new Isolate aka VM
	v, _ := ctx.RunScript("const add = (a, b) => a + b", "math.js") // executes a script on the global context
	fmt.Println(v)
	v, _ = ctx.RunScript("const result = add(3, 4)", "main.js") // any functions previously added to the context can be called
	fmt.Println(v)
	val, _ := ctx.RunScript("result", "value.js") // return a value in JavaScript back to Go
	fmt.Printf("addition result: %s", val)
}
