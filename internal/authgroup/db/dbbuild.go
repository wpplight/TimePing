package db

import (
	"errors"
	"io"
	"os"
	"timeping/internal/tlog"
)

//当用户总文件没有被构造时候创造默认用户组总文件
func create_allauth() (error) { 
	var err error
	var auall *os.File
	if auall,err=os.Create(path+"/auth.tp");err!=nil{
			tlog.Common("创建失败create authdb file fail"+err.Error(),"Error","auth")
			tlog.Err_in("创建失败create authdb file fail"+err.Error(),"Error","auth")
			return errors.New("create authdb file fail")
	}
	_,err=auall.Seek(0,io.SeekStart)
	if err!=nil{
		tlog.Common("操作失败authdb file fail","Error","auth")
		tlog.Err_in("操作失败 authdb file fail"+err.Error(),"Error","auth")
		return errors.New("create authdb file fail")
	}
	_,err=auall.Write(uint16toByte(0))
	if err!=nil{
		tlog.Common("写入失败write authdb file fail","Error","auth")
		tlog.Err_in("写入失败,用户基础文件，检查权限"+err.Error(),"Error","auth")
		return errors.New("write authdb file fail")
	}
	err=auall.Sync()
	if err!=nil{
		tlog.Common("同步失败sync authdb file fail","Error","auth")
		tlog.Err_in("同步失败,用户基础文件，检查权限"+err.Error(),"Error","auth")
		return errors.New("sync authdb file fail")
	}
	auall.Close()
	return nil
}



func syncCheck() error { 
	var err error

	//打开两个文件
	auth,err=os.OpenFile(path+"auth.tp",os.O_RDWR,0644)
	if err!=nil{
		tlog.Common("打开失败open authdb file fail","Error","auth")
		tlog.Err_in("打开失败,用户基础文件，检查权限"+err.Error(),"Error","auth")
		return errors.New("open auth file fail")
	}
	shadow,err=os.OpenFile(path+"shadow.tp",os.O_RDWR,0644)
	if err!=nil{
		tlog.Common("打开失败open shadow file fail","Error","auth")
		tlog.Err_in("打开失败,密码文件，检查权限"+err.Error(),"Error","auth")
		return errors.New("open shadow file fail")
	}

	//文件校验和用户数据缓存
	err=authLoadCheck()
	if err!=nil{
		tlog.Common("用户数据文件校验失败","Error","auth")
		tlog.Err_in("用户数据文件校验失败"+err.Error(),"Error","auth")
		return errors.New("auth file check fail")
	}
	
	//依照用户表进行密码检查
	if checkPass()!=nil{
		tlog.Log_all("auth file check error","Error","authdb")
		return errors.New("check pass fail")
	}

	return nil
}

