package main

import (
	"fmt"
	"golang.org/x/sync/singleflight"
	"time"
)

var sf singleflight.Group

func main() {
	g := new(singleflight.Group)
	go func() {
		// 相同的key视为同一请求
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

// 解决阻塞问题
func singleFlightTimeout(req int64) {
	ch := sf.DoChan("getData", func() (interface{}, error) {
		return getData(req), nil
	})

	select {
	case <-time.After(500 * time.Millisecond):
		return
	case ret := <-ch:
		fmt.Println("结果输出", ret)
	}
}

// singleflight.Forget(key) 删除key，重新执行
