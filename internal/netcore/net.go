package netcore

import (
	"net"
	"strconv"
	"timeping/api/proto/netsend/netgrpc"
	"timeping/internal/config"
	"timeping/internal/tlog"
	"google.golang.org/grpc"
)

type NetCore struct{
	listen net.Listener
	server *grpc.Server
}
type netserver struct{
	netgrpc.UnimplementedTpServiceServer
}

func NewNet() (*NetCore,error){
	n:=new(NetCore)
	var err error
	n.listen,err=net.Listen("tcp","127.0.0.1:"+strconv.Itoa(int(config.Conf.Port)))
	if err!=nil{
		tlog.Common("Listen error","netcore")
		return nil,err
	}
	return n,nil
}
func(n *NetCore) Run() error {
	s:=grpc.NewServer()
	netgrpc.RegisterTpServiceServer(s,&netserver{})
	if err:=s.Serve(n.listen);err!=nil{
		return err
	}

	return nil
}

