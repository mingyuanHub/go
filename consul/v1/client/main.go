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

	c := proto.NewTestServerClient(client)
	r, err := c.Hello(context.Background(), &proto.HelloRequest{
		Name: "i am client",
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(r.Message)
}
