package service

import (
	"context"
	v1 "github.com/raw34/petstore/api/greeter/v1"
	"github.com/raw34/petstore/internal/biz"
)

type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc *biz.GreeterUsecase
}

func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

func (s *GreeterService) GreeterSayHello(ctx context.Context, req *v1.GreeterSayHelloRequest) (*v1.GreeterV1HelloReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: req.Name})
	if err != nil {
		return nil, err
	}
	return &v1.GreeterV1HelloReply{Message: "Hello " + g.Hello + ", From " + req.GetFrom() + ", Device " + req.GetDevice()}, nil
}
