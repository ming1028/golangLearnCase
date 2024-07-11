package main

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/etcdv3"
	"github.com/go-kit/kit/sd/lb"
	etcdLog "github.com/go-kit/log"
	"github.com/golangLearnCase/grpc_etcd/server/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"os"
	"time"
)

const (
	EtcdServerAddress = "localhost:2379"
	Prefix            = "/services/"
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
	logger := etcdLog.NewJSONLogger(os.Stdout)

	// 创建实例管理器, 此管理器会Watch监听etcd中prefix的目录变化更新缓存的服务实例数据
	instancer, err := etcdv3.NewInstancer(etcdClient, ServiceKey, logger)
	if err != nil {
		log.Fatal(err)
	}

	endpointer := sd.NewEndpointer(instancer, ReqFactory(), logger)

	// 创建负载均衡器
	balancer := lb.NewRoundRobin(endpointer)

	/*// 方式1 直接获得请求的endpoint
	reqEndpoint, err := balancer.Endpoint()
	if err != nil {
		log.Fatal(err)
	}*/
	// 方式2 定义重试次数的请求
	reqEndpoint := lb.Retry(3, time.Second, balancer)
	req := &pb.SearchReq{
		Name: "kit etcd request:",
	}
	response, err := reqEndpoint(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	resp := response.(*pb.SearchResp)
	log.Println(resp.GetRespName())
}

func ReqFactory() sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		conn, err := grpc.NewClient(instance, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return nil, nil, err
		}
		client := pb.NewSearchServiceClient(conn) // 生成服务客户端
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			req := request.(*pb.SearchReq)
			return client.Search(ctx, req)
		}, conn, nil
	}
}
