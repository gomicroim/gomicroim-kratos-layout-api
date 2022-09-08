package service

import (
	"context"

	v1 "github.com/go-kratos/kratos-layout/api/helloworld/v1"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	userGrpcClient UserClient
}

// NewGreeterService new a greeter service.
func NewGreeterService(client UserClient) *GreeterService {
	return &GreeterService{userGrpcClient: client}
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.userGrpcClient.Auth(ctx, &AuthRequest{UserName: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: g.Msg}, nil
}
