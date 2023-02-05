package client

import (
	iconsul "mingyuan/consul/v1/consul"
	"mingyuan/consul/v1/v1-service2/service"
)

func GetService2Client() {
	var conf = &iconsul.ConsulConf{
		Discovery: &iconsul.DiscoveryConf{
			Name: service.ServerName,
		},
	}

	iconsul.Discovery(conf)
}