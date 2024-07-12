package search

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/transport/grpc"
	"github.com/golangLearnCase/grpc_etcd/server/proto/search"
	"github.com/spf13/cast"
	"log"
	"math/rand"
)

type SearchService struct {
	search.UnimplementedSearchServiceServer
	SearchEndpointHandler grpc.Handler
}

func (s *SearchService) Search(
	ctx context.Context,
	req *search.SearchReq,
) (
	*search.SearchResp,
	error,
) {
	log.Println("Search Request:", req.GetName())
	return &search.SearchResp{
		RespName: req.GetName() + cast.ToString(rand.Intn(100)),
	}, nil
}

func SearchEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*search.SearchReq)
		return &search.SearchResp{
			RespName: req.GetName() + cast.ToString(rand.Intn(100)),
		}, nil
	}
}
