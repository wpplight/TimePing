package netcore

import (
	"net"
	"strconv"
	"timeping/tlog"
	"timeping/utype"
)

type NetCore struct{
	port int

}

func NewNet() error{

	listen,err:=net.Listen("tcp","127.0.0.1:"+strconv.Itoa(int(utype.Conf.Port)))
	if err!=nil{
		tlog.Common("Listen error","netcore")
		return err
	}
	
}