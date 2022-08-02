package main

import (
	"errors"
	"fmt"
	err2 "github.com/pkg/errors"
	"os"
	"time"
)

func main() {
	fmt.Println(errors.New("错误"))
	fmt.Printf("current time: %v\n", time.Now())
	fmt.Println(time.Now().Year(),
		time.Now().Month(), time.Now().Day(), time.Now().Hour(),
		time.Now().Minute(), time.Now().Second(), time.Now().Weekday(),
	)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000 Mon Jan"))
	fmt.Println(time.Now().Format("2006-01-02 03:04:05.000 PM"))
	fmt.Println(time.Now().Format("2006/01/02 15:04:05.999 PM"))

	timeFormat := "2006-01-02 15:04:05.999"
	loc, _ := time.LoadLocation("Asia/Shanghai")
	fmt.Println(time.Parse(timeFormat, "2022-10-05 11:25:20"))
	fmt.Println(time.ParseInLocation(timeFormat, "2022-10-05 11:25:20", loc))

	_, err := os.Stat("stderr")
	fmt.Println(err)

	err = Error1()
	if err != nil {
		fmt.Printf("original error:%T %v\n", err2.Cause(err), err2.Cause(err))
		fmt.Printf("stack trace:\n%+v\n", err)
	}
}

func ErrorTemplate() error {
	err := errors.New("模拟错误")
	return err2.Wrap(err, "template error")
}

func Error1() error {
	err := ErrorTemplate()
	return err2.WithMessage(err, "调用")
}
