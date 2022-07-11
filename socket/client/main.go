package main

import (
	"socket/config"
	"socket/util"
	"fmt"
	"net"
)

func main() {
	fmt.Printf("client connect %v\n", config.ServerAddress)
	conn, err := net.Dial(config.ServerNetworkType, config.ServerAddress)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Printf("client %v connected\n", conn.LocalAddr())

	util.Write(conn, "hello111")
	if _, err := util.Read(conn); err != nil {
		fmt.Println(err)
	}

	util.Write(conn, "hello222")
	if _, err := util.Read(conn); err != nil {
		fmt.Println(err)
	}


}