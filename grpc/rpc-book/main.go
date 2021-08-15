package main

import (
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
	//rpc.RegisterName("HelloService", new(HelloService))

	RegisterHelloService(new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {

	}
	for {
		conn, err := listener.Accept()
		if err != nil {

		}
		go rpc.ServeConn(conn)
	}

}
