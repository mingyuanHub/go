### consule

```html

v1-client ---[grpc]--->>> v1-service1  
    
v1-service1  ---[consul 服务发现， 获取service2 地址] --- [grpc] --->>>  v1-service2

```