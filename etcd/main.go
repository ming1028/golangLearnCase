package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"log"
	"time"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatalf("connect etcd err:%+v\n", err)
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = cli.Put(ctx, "keyName", "value")

	resp, err := cli.Get(ctx, "keyName")
	for _, v := range resp.Kvs {
		fmt.Printf("%s:%s\n", v.Key, v.Value)
	}

	rch := cli.Watch(ctx, "keyName")
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Println(ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}
