package main

import (
	"socket/config"
	"socket/util"
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Printf("client connect %v\n", config.ServerAddress)
	conn, err := net.Dial(config.ServerNetworkType, config.ServerAddress)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Printf("client %v connected\n", conn.LocalAddr())

	for {
		util.Write(conn, "hello")
		//如果服务断开，则报错：read tcp 127.0.0.1:59262->127.0.0.1:9090: wsarecv: An existing connection was forcibly closed by the remote host.
		if _, err := util.Read(conn); err != nil {
			fmt.Println(err)
		}

		time.Sleep(5 * time.Second)
	}

}