package main

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/sd/etcdv3"
	kitTransGrpc "github.com/go-kit/kit/transport/grpc"
	etcdLog "github.com/go-kit/log"
	"github.com/golangLearnCase/grpc_etcd/server/proto/pb"
	"github.com/golangLearnCase/grpc_etcd/server/service/search"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"time"
)

const (
	ServiceAddress    = ":9100"
	EtcdServerAddress = "localhost:2379"
	Prefix            = "/services/"
	ServiceValue      = "localhost:9100"
	ServiceKey        = Prefix + "search-service"
)

func main() {
	// etcd连接参数
	option := etcdv3.ClientOptions{
		DialTimeout:   5 * time.Second,
		DialKeepAlive: 5 * time.Second,
	}
	// 连接etcd
	etcdClient, err := etcdv3.NewClient(context.Background(), []string{EtcdServerAddress}, option)
	if err != nil {
		log.Fatal(err)
	}
	// 服务名称、地址
	service := etcdv3.Service{
		Key:   ServiceKey,
		Value: ServiceValue,
	}
	// 注册
	register := etcdv3.NewRegistrar(etcdClient, service, etcdLog.NewJSONLogger(os.Stdout))
	register.Register() // 注册服务

	listener, err := net.Listen("tcp", ServiceAddress) //网络监听，注意对应的包为："net"
	if err != nil {
		fmt.Println(err)
		return
	}
	searchHandler := kitTransGrpc.NewServer(
		search.SearchEndpoint(),
		decodeRequest,
		encodeResponse,
	)
	gs := grpc.NewServer(grpc.UnaryInterceptor(kitTransGrpc.Interceptor))
	pb.RegisterSearchServiceServer(gs, &search.SearchService{
		SearchEndpointHandler: searchHandler,
	})
	gs.Serve(listener)
}

func decodeRequest(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

func encodeResponse(_ context.Context, rsp interface{}) (interface{}, error) {
	return rsp, nil
}
