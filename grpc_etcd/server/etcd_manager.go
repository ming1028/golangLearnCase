package main

import (
	"context"
	"fmt"
	searchPb "github.com/golangLearnCase/grpc_etcd/server/proto/search"
	"github.com/golangLearnCase/grpc_etcd/server/service/search"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"google.golang.org/grpc"
	"net"
	"time"
)

const (
	etcdEndpoints = "localhost:2388" // etcd 服务地址
	serviceName   = "search"
	serviceKey    = "/services/" + serviceName
	serviceValue  = "127.0.0.1:9099"
)

func main() {
	l, err := net.Listen("tcp", serviceValue)
	if err != nil {
		panic(err)
	}

	cfg := clientv3.Config{
		Endpoints:         []string{etcdEndpoints},
		DialKeepAliveTime: 5 * time.Second,
	}
	client, err := clientv3.New(cfg)
	if err != nil {
		panic(err)
	}

	em, err := endpoints.NewManager(client, serviceKey)
	if err != nil {
		panic(err)
	}

	var ttl int64 = 30
	leaseResp, err := client.Grant(context.Background(), ttl)
	if err != nil {
		panic(err)
	}
	err = em.AddEndpoint(context.Background(), serviceKey+"/"+serviceValue, endpoints.Endpoint{
		Addr: serviceValue,
		Metadata: map[string]string{
			"region": "shanghai",
		},
	}, clientv3.WithLease(leaseResp.ID))
	if err != nil {
		panic(err)
	}

	kaCtx, kaCancel := context.WithCancel(context.Background())
	go func() {
		// 在这里操作续约
		ch, err1 := client.KeepAlive(kaCtx, leaseResp.ID)
		if err1 != nil {
			fmt.Println(err1)
		}
		for kaResp := range ch {
			// 正常就是打印一下 DEBUG 日志啥的
			fmt.Println(kaResp.String(), time.Now().String())
		}
	}()

	grpcServer := grpc.NewServer()
	searchPb.RegisterSearchServiceServer(grpcServer, &search.SearchService{})
	err = grpcServer.Serve(l)
	fmt.Println(err)
	// 你要退出了，正常退出
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// 我要先取消续约
	kaCancel()
	// 退出阶段，先从注册中心里面删了自己
	err = em.DeleteEndpoint(ctx, serviceKey)
	// 关掉客户端
	client.Close()
	grpcServer.GracefulStop()
}
