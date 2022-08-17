package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"mingyuan/consul/v1/proto"
	"log"
)

func main() {
	client, err := grpc.Dial("127.0.0.1:6001",grpc.WithInsecure())
	if err != nil {
		log.Fatal(err.Error())
	}

	defer client.Close()

	c := proto.NewProtoServer1Client(client)
	r, err := c.Hi(context.Background(), &proto.HiRequest{
		Name: "i am client",
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(r.Message)
}