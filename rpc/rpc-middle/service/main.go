package main

import (
	"log"
	"net"
	"net/rpc"
)


type HelloService struct {

}

func (p *HelloService) Hello(request string, reply *string) error  {
	*reply = "hello " + request
	return nil
}

func main() {
	RegisterHelloService(new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err.Error())
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf(err.Error())
		}
		go rpc.ServeConn(conn)
	}

}
