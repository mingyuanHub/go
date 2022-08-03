package main

import (
	"fmt"
	"github.com/ryandeng/goversion"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"mingyuan/consul/v1/proto"
	"net"
)

func main() {
	fmt.Println(goversion.Version())

	lis, err := net.Listen("tcp", ":6000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srvOpts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(

		),
	}
	srv := grpc.NewServer(srvOpts...)

	var people = &People{}

	proto.RegisterHelloServer(srv, people)

	// 注册服务端反射服务
	reflection.Register(srv)

	srv.Serve(lis)
}