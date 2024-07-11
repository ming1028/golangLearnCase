package search

import (
	"context"
	"github.com/golangLearnCase/grpc_etcd/server/proto/pb"
	"github.com/spf13/cast"
	"log"
	"math/rand"
)

type SearchService struct {
	pb.UnimplementedSearchServiceServer
}

func (s *SearchService) Search(
	ctx context.Context,
	req *pb.SearchReq,
) (
	*pb.SearchResp,
	error,
) {
	log.Println("Search Request:", req.GetName())
	return &pb.SearchResp{
		RespName: req.GetName() + cast.ToString(rand.Intn(100)),
	}, nil
}
