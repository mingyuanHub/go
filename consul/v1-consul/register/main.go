package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
	"net"
)

func main() {
	// 使用默认配置创建consul客户端
	consulClient, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		log.Fatal(err)
		fmt.Println(11111111111, err.Error())
	}

	// 注册服务
	// consulClient.Agent()先获取当前机器上的consul agent节点
	consulClient.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID:      "MyService 111",
		Name:    "My Service",
		Address: "127.0.0.1",
		Port:    5050,
		Check: &api.AgentServiceCheck{
			CheckID:  "MyService Check 111111",
			TCP:      "127.0.0.1:5050",
			Interval: "10s",
			Timeout:  "1s",
		},
	})

	// 运行完成后注销服务
	defer consulClient.Agent().ServiceDeregister("MyService")

	l, err := net.Listen("tcp", ":5050")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go func() {
			log.Printf("Ip: %s connected", conn.RemoteAddr().String())
		}()
	}
}
