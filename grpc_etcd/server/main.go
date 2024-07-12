package main

import (
	"github.com/golangLearnCase/grpc_etcd/server/pkg/xetcd"
	searchPb "github.com/golangLearnCase/grpc_etcd/server/proto/search"
	"github.com/golangLearnCase/grpc_etcd/server/service/search"
	"github.com/spf13/cast"
	"google.golang.org/grpc"
	"log"
	"net"
)

const PORT = 9099

func main() {
	grpcServer := grpc.NewServer()
	searchPb.RegisterSearchServiceServer(grpcServer, &search.SearchService{})

	lis, err := net.Listen("tcp", ":"+cast.ToString(PORT))
	if err != nil {
		log.Fatalf("net listen err: %v", err)
	}

	// 服务注册
	xetcd.RegisterEtcdService()

	log.Fatal(grpcServer.Serve(lis))
}
