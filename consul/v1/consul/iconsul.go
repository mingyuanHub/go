package iconsul

import (
	"google.golang.org/grpc"
	"net"
)


func ServeWithConsul(conf *ConsulConf, srv *grpc.Server, lister net.Listener) error {

	Register(conf)

	return srv.Serve(lister)
}