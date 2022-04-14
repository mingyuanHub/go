package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
	"test1/services"
)

func main() {
	creds, err := credentials.NewServerTLSFromFile("keys/server_no_passwd.crt", "keys/server_no_passwd.key")
	if err != nil {
		fmt.Println(err)
	}
	rpcServer := grpc.NewServer(grpc.Creds(creds))

	//rpcServer := rpc-middle.NewServer()
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))
	services.RegisterMaServiceServer(rpcServer, new(services.MaService))

	listen , err := net.Listen("tcp", ":8084")

	if err != nil {
		fmt.Println(err)
	}

	rpcServer.Serve(listen)

}
