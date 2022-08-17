package service

import (
	"context"
	"mingyuan/consul/v1/proto"
)

const (
	ServerName = "v1-service2"
	ServerAddr = "127.0.0.1:6002"
)

type Service struct{
	Name string
}

func (p *Service) Hello(c context.Context, s *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{
		Message: "hello, " + s.Name,
	}, nil
}
