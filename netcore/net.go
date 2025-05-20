package netcore

import (
	"net"
	"strconv"
	"timeping/global"
	"timeping/tlog"
)

type NetCore struct{
	listen net.Listener
}

func NewNet() (*NetCore,error){

	listen,err:=net.Listen("tcp","127.0.0.1:"+strconv.Itoa(int(global.Conf.Port)))
	if err!=nil{
		tlog.Common("Listen error","netcore")
		return nil,err
	}
	n:=new(NetCore)
	n.listen=listen
	return n,nil
}

