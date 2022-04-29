package main

import (
	"errors"
	"fmt"
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
}
