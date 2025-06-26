package db

import (
	"errors"
	"timeping/internal/tlog"
)

func setAuthLen(n uint16) error {
	if _,err:=auth.Seek(0,0);err!=nil{
		tlog.Common("查找auth文件错误","Error","auth")
		tlog.Err_in("auth file seek fail"+err.Error(),"Error","auth")
		return errors.New("auth file seek fail")
	}
	_,err:=auth.Write(uint16toByte(n))
	if err!=nil{
		tlog.Common("写入auth文件错误","Error","auth")
		tlog.Err_in("auth file write fail"+err.Error(),"Error","auth")
		return errors.New("auth file write fail")
	}
	if err=auth.Sync() ;err!=nil{
		tlog.Common("同步auth文件错误","Error","auth")
		tlog.Err_in("auth file sync fail"+err.Error(),"Error","auth")
		return errors.New("auth file sync fail")
	}
	return nil

}

func getAuthLen() (uint16,error) { 
	if auth==nil{
		tlog.Log_all("auth file is nil","Error","auth")
		return 0,errors.New("auth file is nil")
	}
	_,err:=auth.Seek(0,0)
	if err!=nil{
		tlog.Common("auth file seek fail","Error","auth")
		tlog.Err_in("auth file seek fail"+err.Error(),"Error","auth")
		return 0,errors.New("auth file seek fail")
	}
	b:=make([]byte,2)
	_,err=auth.Read(b)
	if err!=nil{
		tlog.Common("auth file read fail","Error","auth")
		tlog.Err_in("auth file read fail"+err.Error(),"Error","auth")
		return 0,errors.New("auth file read fail")
	}
	return byteToUint16(b),nil
}

//将用户表单载入到内存中
func loadUsr() error{ 
	if usrtable==nil{
		tlog.Log_all("usr table is nil","Error","auth")
		return errors.New("usr table is nil")
	}
	if asize==0{
		tlog.Log_all("auth size is zero","Error","auth")
		return errors.New("auth size is zero")
	}
	for i:=0;i<int(asize);i++{
		
	}
}