package main

import (
	"context"
	"mingyuan/consul/v1/proto"
)

type Service struct{
	Name string
}

func (p *Service) Hello(c context.Context, s *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{
		Message: "hello, " + s.Name,
	}, nil
}
