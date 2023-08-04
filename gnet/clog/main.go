package main

import (
	"fmt"
	"github.com/panjf2000/gnet"
	"github.com/panjf2000/gnet/pkg/pool/goroutine"
	"log"
	"time"
)

type echoServer struct {
	*gnet.EventServer
	pool *goroutine.Pool
}

func (es *echoServer) React(
	frame []byte,
	c gnet.Conn,
) (
	out []byte,
	action gnet.Action,
) {
	data := append([]byte{'a'}, frame...)
	fmt.Println(string(data))

	_ = es.pool.Submit(func() {
		time.Sleep(time.Second * 1)
		c.AsyncWrite(data)
	})
	return
}

func main() {
	p := goroutine.Default()
	defer p.Release()

	echo := &echoServer{}
	log.Fatal(gnet.Serve(echo, "tcp://:9001", gnet.WithMulticore(true)))
}
