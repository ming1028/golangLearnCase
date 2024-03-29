package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	type stackTracer interface {
		StackTrace() errors.StackTrace
	}
	err := err1()
	e := errors.Cause(err)
	er, _ := e.(stackTracer)
	fmt.Printf("%+v", er.StackTrace())
	fmt.Println(e.Error())
	fmt.Printf("%T %v\n", e, e)
}

func err1() error {
	err := err2()
	return errors.WithMessagef(err, "err1")
}

func err2() error {
	err := err3()
	return errors.WithMessage(err, "err2 调用")
}

func err3() error {
	return errors.Wrapf(errors.New("err3错误发生"), "err3")
}
