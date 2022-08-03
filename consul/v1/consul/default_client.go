package iconsul

import (
	"github.com/hashicorp/consul/api"
)

type defaultClient struct {

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