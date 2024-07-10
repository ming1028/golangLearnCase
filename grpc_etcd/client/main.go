package main

import (
	"context"
	"github.com/golangLearnCase/grpc_etcd/server/proto/pb"
	"github.com/spf13/cast"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const PORT = 9099

func main() {
	conn, err := grpc.NewClient(":"+cast.ToString(PORT), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	searchClient := pb.NewSearchServiceClient(conn)
	resp, err := searchClient.Search(context.Background(), &pb.SearchReq{
		Name: "grpc search",
	})
	if err != nil {
		panic(err)
	}
	log.Printf("resp:%s", resp.GetRespName())
}
