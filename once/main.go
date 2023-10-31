package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var (
	tokenOnce Once
	token     *tokenStruct
)

type Once struct {
	done uint32
	m    sync.Mutex
}

type tokenStruct struct {
	Expire time.Time
	Token  string
}

func main() {
	rand.Seed(time.Now().UnixMicro())
	for i := 0; i < 10; i++ {
		v := i
		go func() {
			rd := rand.Intn(3) + 1
			fmt.Println("协程", v, rd)
			time.Sleep(time.Second * time.Duration(rd))
			fmt.Println(getToken())
		}()
	}
	time.Sleep(time.Hour)
}

func getToken() *tokenStruct {
	if token != nil && token.Expire.Before(time.Now()) {
		tokenOnce.doReset()
	}
	tokenOnce.Do(func() {
		fmt.Println("once get")
		token = &tokenStruct{
			Expire: time.Now().Add(time.Second * 2),
			Token:  "333",
		}
	})
	return token
}

func (o *Once) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 0 {
		o.doSlow(f)
	}
}

func (o *Once) doSlow(f func()) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}

func (o *Once) doReset() {
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 1 {
		atomic.StoreUint32(&o.done, 0)
	}
}
