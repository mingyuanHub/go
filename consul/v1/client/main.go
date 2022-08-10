package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"mingyuan/consul/v1/proto"
	"log"
)

func main() {
	client, err := grpc.Dial("localhost:6000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err.Error())
	}

	defer client.Close()

	c := proto.NewHelloClient(client)
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{
		Name: "world",
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(r.Message)
}