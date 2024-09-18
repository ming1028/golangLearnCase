package main

import (
	"context"
	"fmt"
	"github.com/wuqinqiang/easyio"
	"os"
	"os/signal"
	"syscall"
)

var _ easyio.EventHandler = (*Handler)(nil)

type EasyioKey struct{}

type Handler struct{}

var CtxKey EasyioKey

func (h Handler) OnOpen(c easyio.Conn) context.Context {
	return context.WithValue(context.Background(), CtxKey, Message{Msg: "helloword"})
}

func (h Handler) OnRead(ctx context.Context, c easyio.Conn) {
	_, ok := ctx.Value(CtxKey).(Message)
	if !ok {
		return
	}
	var b = make([]byte, 100)
	_, err := c.Read(b)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("[Handler] read data:", string(b))
	if _, err = c.Write(b); err != nil {
		panic(err)
	}
}

func (h Handler) OnClose(_ context.Context, c easyio.Conn) {
	fmt.Println("[Handler] closed", c.Fd())
}

type Message struct {
	Msg string
}

func main() {
	e := easyio.New("tcp", ":8080", easyio.WithNumPoller(4),
		easyio.WithEventHandler(Handler{}))
	if err := e.Start(); err != nil {
		panic(err)
	}
	defer e.Stop()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	<-c
}
