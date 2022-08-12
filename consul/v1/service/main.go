package main

import (
	"fmt"
	"github.com/ryandeng/goversion"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"mingyuan/consul/v1/common"
	iconsul "mingyuan/consul/v1/consul"
	"mingyuan/consul/v1/proto"
	"net"
)

const ServerName = "GrpcService"

type Conf struct {
	Addr string
	Consul *iconsul.ConsulConf
}

func main() {
	var err error

	fmt.Println(goversion.Version())

	var conf = &Conf{
		Addr: "127.0.0.1:6000",
		Consul:&iconsul.ConsulConf{
			Registration: &iconsul.Registration{
			},
		},
	}

	conf.Consul.Registration.Name = ServerName

	if conf.Consul.Registration.Port, err = common.GetPort(conf.Addr); err != nil {
		panic(err)
	}

	conf.Consul.Registration.Id = fmt.Sprintf("%s-%s-%s", ServerName, common.GetLANHost(), conf.Addr)

	lis, err := net.Listen("tcp", ":6000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srvOpts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(

		),
	}
	srv := grpc.NewServer(srvOpts...)

	var service = &Service{}

	proto.RegisterHelloServer(srv, service)

	// 注册服务端反射服务
	reflection.Register(srv)


	iconsul.ServeWithConsul(conf.Consul, srv, lis)
}