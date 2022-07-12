package main

import (
	"socket/config"
	"socket/util"
	"fmt"
	"io"
	"net"
	"time"
)

func handleConnect(conn net.Conn) {
	fmt.Printf("client %v connected\n", conn.RemoteAddr())
	for {
		conn.SetReadDeadline(time.Now().Add(time.Second * time.Duration(2))) //设置读取超时时间
		if str, err := util.Read(conn); err != nil {
			if err == io.EOF {
				fmt.Printf("client %v closed\n", conn.RemoteAddr())
				break
			} else {
				fmt.Printf("read error:%v\n", err.Error())
			}
		} else {
			util.Write(conn, "welcome " + str, )
		}
	}
}

func main() {
	fmt.Printf("server start %v\n", config.ServerAddress)
	listener, err := net.Listen(config.ServerNetworkType, config.ServerAddress)

	////第二种监听方式 ListenTCP
	//tcpServer, _ := net.ResolveTCPAddr("127.0.0.1","9090")
	//listener, _ = net.ListenTCP("tcp", tcpServer)

	if err != nil {
		panic(err)
	}
	defer listener.Close()
	fmt.Printf("waiting client connect...\n")
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnect(conn)
	}
}