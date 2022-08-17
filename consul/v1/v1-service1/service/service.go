package service

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"mingyuan/consul/v1/proto"
)

const (
	ServerName = "v1-service1"
	ServerAddr = "127.0.0.1:6001"
)

type Service struct {
	Name string

	Service2Cli proto.ProtoServer2Client
}

func NewService() *Service {

	client, err := grpc.Dial("127.0.0.1:6002", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err.Error())
	}
	//defer client.Close()

	c := proto.NewProtoServer2Client(client)

	return &Service{
		Service2Cli: c,
	}
}

func (s *Service) Hi(c context.Context, r *proto.HiRequest) (*proto.HiReply, error) {

	msg := s.Hello(r.Name)

	return &proto.HiReply{
		Message: "Hi, " + msg,
	}, nil
}

func (s *Service) Hello(msg string) string {

	r, err := s.Service2Cli.Hello(context.Background(), &proto.HelloRequest{
		Name: "i am service1 with msg=" + msg,
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	return r.Message
}
