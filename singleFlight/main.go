package main

import (
	"fmt"
	"golang.org/x/sync/singleflight"
	"time"
)

func main() {
	g := new(singleflight.Group)
	go func() {
		v1, _, shared := g.Do("getData", func() (interface{}, error) {
			ret := getData(1)
			return ret, nil
		})
		fmt.Printf("1st call: v1:%v, shared:%v\n", v1, shared)
	}()
	time.Sleep(2 * time.Second)
	v2, _, shared := g.Do("getData", func() (interface{}, error) {
		ret := getData(1)
		return ret, nil
	})
	fmt.Printf("2st call: v2:%v, shared:%v\n", v2, shared)
}

func getData(id int64) string {
	fmt.Println("getData")
	time.Sleep(10 * time.Second)
	return "getData"
}
