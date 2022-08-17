package iconsul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"mingyuan/consul/v1/common"
	"net"
)

type ConsulConf struct {
	Registration *Registration
}

type Registration struct {
	Name string
	Id   string
	Addr string
}


func ServeWithConsul(conf *ConsulConf, srv *grpc.Server, lister net.Listener) error {

	ip, err := common.GetIp(conf.Registration.Addr)
	if err != nil {
		return err
	}

	port, err := common.GetPort(conf.Registration.Addr)
	if err != nil {
		return err
	}

	fmt.Println(ip, port)

	// 注册服务
	// consulClient.Agent() 先获取当前机器上的consul agent节点
	iconsul.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID:      conf.Registration.Id,
		Name:    conf.Registration.Name,
		Address: ip,
		Port:    port,
		Check: &api.AgentServiceCheck{
			CheckID:  "check-" + conf.Registration.Id,
			TCP:      conf.Registration.Addr,
			Interval: "10s",
			Timeout:  "1s",
		},
	})

	// 运行完成后注销服务
	defer iconsul.Agent().ServiceDeregister(conf.Registration.Id)

	return srv.Serve(lister)
}