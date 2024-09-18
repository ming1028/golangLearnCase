package main

import (
	"context"
	"fmt"
	"github.com/spf13/cast"
	"sync"
	"time"
)

var (
	wg    sync.WaitGroup
	syMap sync.Map
)

func main() {
	/*wg.Add(1)
	go worker()
	wg.Wait()
	fmt.Println("over")*/
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	go doSomething(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("超时")
		cancel()
	}
	time.Sleep(time.Second * 20)
	for i := 0; i < 3; i++ {
		var Wat WxAccToken
		val, ok := AppIdAccToken.Load(cast.ToString(1))
		if ok {
			Wat = val.(WxAccToken)
		}
		Wat.ExpiresAt = time.Now().Add(100 * time.Hour)
		Wat.AccessToken = "dsfsdfsf"
		Wat.ExpiresIn = 20
		AppIdAccToken.Store(cast.ToString(1), Wat)
	}

	AppIdAccToken.Range(func(k, value any) bool {
		fmt.Println(k.(string), value)
		return true
	})
}

func worker() {
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
	}
	wg.Done()
}

var AppIdAccToken sync.Map

type WxAccToken struct {
	ExpiresAt   time.Time `json:"expires_at"`
	AccessToken string    `json:"access_token"`
	ExpiresIn   int32     `json:"expires_in"`
}

func doSomething(ctx context.Context) {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}
