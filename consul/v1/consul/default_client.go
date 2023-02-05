package iconsul

import (
	"errors"
	"fmt"
	"github.com/hashicorp/consul/api"
	"mingyuan/consul/v1/common"
)

type defaultClient struct {

}

type ConsulConf struct {
	Registration *RegistrationConf
	Discovery *DiscoveryConf
}

type RegistrationConf struct {
	Name string
	Id   string
	Addr string
}

type DiscoveryConf struct {
	Name string
}

var (
	DefaultClient = &defaultClient{}

	iconsul *api.Client
)

func (c *defaultClient) Start() (err error) {
	// 使用默认配置创建consul客户端
	iconsul, err = api.NewClient(api.DefaultConfig())

	return
}


func Register(conf *ConsulConf) (err error) {

	if DefaultClient == nil {
		err = errors.New("unStart defaultClient")
		return
	}

	if conf.Registration == nil {
		err = errors.New("invalid conf.RegistrationConf")
		return
	}

	ip, err := common.GetIp(conf.Registration.Addr)
	if err != nil {
		return err
	}

	port, err := common.GetPort(conf.Registration.Addr)
	if err != nil {
		return err
	}

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
	//defer iconsul.Agent().ServiceDeregister(conf.RegistrationConf.Id)

	return
}

func Discovery(conf *ConsulConf) (err error) {
	if DefaultClient == nil {
		err = errors.New("unStart defaultClient")
		return
	}

	if conf.Discovery == nil {
		err = errors.New("invalid conf.DiscoveryConf")
		return
	}

	catalogService, _, err := iconsul.Catalog().Service(conf.Discovery.Name, "", nil)

	if err != nil {
		return err
	}

	for _, c := range catalogService{
		fmt.Println(c.ID, c.Node, c.ServiceID, c.ServiceName,c.ServiceID)
	}

	return nil
}