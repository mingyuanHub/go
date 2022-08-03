package main

import (
	"context"
	"mingyuan/consul/v1/proto"
)

type People struct{
	Name string
}

func (p *People) SayHello(c context.Context, s *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{
		Message: "hello " + s.Name,
	}, nil
}
