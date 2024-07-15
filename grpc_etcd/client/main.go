package main

import (
	"context"
	"fmt"
	"github.com/golangLearnCase/grpc_etcd/server/proto/search"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const PORT = 9099

const (
	etcdEndpoints = "localhost:2388" // etcd 服务地址
	serviceName   = "search"
	serviceKey    = "/services/" + serviceName
	serviceValue  = "127.0.0.1:9099"
)

func main() {
	cfg := clientv3.Config{
		Endpoints:         []string{etcdEndpoints},
		DialKeepAliveTime: 5 * time.Second,
	}
	client, err := clientv3.New(cfg)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	addrs := DiscoverService(client)
	if len(addrs) == 0 {
		panic("no discovery service")
	}
	etcdResolver, err := resolver.NewBuilder(client)
	if err != nil {
		panic(err)
	}
	// grpcResolver.Register(etcdResolver)
	svcCfg := `
{
    "loadBalancingConfig": [
        {
            "round_robin": {}
        }
    ]
}
`
	conn, err := grpc.NewClient(
		"etcd:///"+serviceKey,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithResolvers(etcdResolver),
		grpc.WithDefaultServiceConfig(svcCfg),
	)
	// conn, err := grpc.NewClient(addrs[0], grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	searchClient := search.NewSearchServiceClient(conn)
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := searchClient.Search(ctxTimeout, &search.SearchReq{
		Name: "grpc search",
	})
	if err != nil {
		panic(err)
	}
	log.Printf("resp:%s", resp.GetRespName())
}

func DiscoverService(client *clientv3.Client) []string {
	/*cfg := clientv3.Config{
		Endpoints:         []string{etcdEndpoints},
		DialKeepAliveTime: 5 * time.Second,
	}
	client, err := clientv3.New(cfg)
	if err != nil {
		panic(err)
	}*/

	resp, err := client.Get(context.Background(), serviceKey, clientv3.WithPrefix())
	if err != nil {
		panic(err)
	}
	var addrs []string
	for _, kv := range resp.Kvs {
		fmt.Println("Discovered service:", string(kv.Key), string(kv.Value))
		addrs = append(addrs, string(kv.Value))
	}
	return addrs
}
