package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ryandeng/goversion"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	iconsul "mingyuan/consul/v1/consul"
	"mingyuan/consul/v1/proto"
	"mingyuan/consul/v1/v1-service1/service"
)

type Conf struct {
	Addr   string
	Consul *iconsul.ConsulConf
}

func main() {

	fmt.Println(goversion.Version())

	var err error

	if err = iconsul.DefaultClient.Start(); err != nil {
		fmt.Println("iconsul start failed,remote config unusable,err:%s", err)
	}

	var conf = &Conf{
		Consul: &iconsul.ConsulConf{
			Registration: &iconsul.RegistrationConf{
				Name: service.ServerName,
				Addr: service.ServerAddr,
				Id:   fmt.Sprintf("%s:%s", service.ServerName, service.ServerAddr),
			},
		},
	}

	srvOpts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(

		),
	}
	srv := grpc.NewServer(srvOpts...)

	var v1Service = service.NewService()

	proto.RegisterProtoServer1Server(srv, v1Service)

	// 注册服务端反射服务
	reflection.Register(srv)

	lis, err := net.Listen("tcp", service.ServerAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	iconsul.ServeWithConsul(conf.Consul, srv, lis)
}