package service

import (
	"context"
	"errors"

	"google.golang.org/grpc"
)

type AuthRequest struct {
	UserName string
}

type AuthReply struct {
	Msg string
}

type UserClient interface {
	Auth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthReply, error)
}

type userClient struct{}

func (u *userClient) Auth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthReply, error) {
	if in.UserName == "admin" {
		return &AuthReply{Msg: "success"}, nil
	}
	return nil, errors.New("error userName")
}

func newUserClient(conn interface{}) UserClient {
	return &userClient{}
}
