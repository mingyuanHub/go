package iconsul

import (
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"net"
)

type ConsulConf struct {
	Registration *Registration
}

type Registration struct {
	Name string
	Id   string
	Addr string
	Port int
}


func ServeWithConsul(conf *ConsulConf, srv *grpc.Server, lister net.Listener) error {

	// 注册服务
	// consulClient.Agent() 先获取当前机器上的consul agent节点
	iconsul.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID:      conf.Registration.Id,
		Name:    conf.Registration.Name,
		Address: "127.0.0.1",
		Port:    conf.Registration.Port,
		Check: &api.AgentServiceCheck{
			CheckID:  "CheckId111111",
			TCP:      "127.0.0.1:5050",
			Interval: "10s",
			Timeout:  "1s",
		},
	})

	// 运行完成后注销服务
	defer iconsul.Agent().ServiceDeregister("MyService")

	return srv.Serve(lister)
}