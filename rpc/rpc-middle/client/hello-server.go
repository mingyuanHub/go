package main

import "net/rpc"

const HelloServiceName = "rpc-basic/rpc-middle/HelloService"

type HelloServiceInterface = interface {
	Hello(request string, reply *string) error
}

//1. 创建一个HelloServiceInterface地址，但不会分配内存的,并且如果给字段赋值会报错。
//2. 在代码中判断HelloServiceClient这个struct是否实现了HelloServiceInterface这个interface。
var _ HelloServiceInterface = (*HelloServiceClient)(nil)

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

type HelloServiceClient struct {
	*rpc.Client
}

func DiaHelloService(network string, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {

	}
	return &HelloServiceClient{Client: c}, nil
}

func (p *HelloServiceClient) Hello(request string, reply *string) error{
	return p.Client.Call(HelloServiceName+".Hello", request, reply)
}