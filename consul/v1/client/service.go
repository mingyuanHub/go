package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"mingyuan/consul/v1/proto"
)

type ClientService struct {
	TestCli proto.TestServerClient
}

func NewClientService() *ClientService {

	client, err := grpc.Dial("localhost:6000")
	if err != nil {
		log.Fatal(err.Error())
	}
	//defer client.Close()

	c := proto.NewTestServerClient(client)

	return &ClientService{
		TestCli: c,
	}
}


func (s *ClientService) test() {

	r, err := s.TestCli.Hello(context.Background(), &proto.HelloRequest{
		Name: "i am client",
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(r.Message)
}
