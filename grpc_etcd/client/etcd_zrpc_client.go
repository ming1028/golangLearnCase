package main

import (
	"context"
	"fmt"
	"github.com/golangLearnCase/grpc_etcd/server/proto/search"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
)

func main() {
	conn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{ // 通过 etcd 服务发现时，只需要给 Etcd 配置即可
			Hosts: []string{"127.0.0.1:2388"},
			Key:   "/services/search", // 服务名前缀
		},
	})

	searchClient := search.NewSearchServiceClient(conn.Conn())
	resp, err := searchClient.Search(context.Background(), &search.SearchReq{
		Name: "zrpc search:",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.GetRespName())
}
