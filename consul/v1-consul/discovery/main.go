package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
)

func main() {
	//使用默认配置创建consul客户端
	consulClient, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		log.Fatal(err)
	}

	catalogService, meta, err := consulClient.Catalog().Service("My Service", "", nil)

	fmt.Println("1111", len(catalogService))

	for _, c := range catalogService{
		fmt.Println(c.ID, c.Node, c.ServiceID, c.ServiceName,c.ServiceID)
	}

	fmt.Println(meta)
}
