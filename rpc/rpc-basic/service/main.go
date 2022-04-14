package main

import (
	"log"
	"net"
	"net/rpc"
)

func main()  {
	rpc.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err.Error())
	}

	con, err := listener.Accept()
	if err != nil {
		log.Println(err.Error())
	}

	rpc.ServeConn(con)

}
