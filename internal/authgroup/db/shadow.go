package db

import (
	"errors"
	"os"
	"timeping/internal/tlog"
	"timeping/pkg/ostools"
)

//密码加密文件

//shadow 文件 RowSize = 16

func createShadow() error{
	p:=path+"shadow.tp"
	err:=ostools.CreateFile(p)
	if err!=nil{
		return err
	}
	return nil
}

//核验密码账户一致性
func firstCheck() error {
	//安全性检测
	if shadow==nil{
		tlog.Log_all("shadow file is nil","Error","auth")
		return errors.New("shadow file is nil")
	}
	if auth==nil{ 
		tlog.Log_all("auth file is nil","Error","auth")
		return errors.New("auth file is nil")
	}
	Asize,err:=os.Stat(path+"auth.tp")
	if err!=nil{
		tlog.Common("auth file is not exist","Error","auth")
		tlog.Err_in("auth file is not exist"+err.Error(),"Error","auth")
		return errors.New("auth file is not exist")
	}
	Ssize,err:=os.Stat(path+"shadow.tp")
	if err!=nil{
		tlog.Common("shadow file is not exist","Error","auth")
		tlog.Err_in("shadow file is not exist"+err.Error(),"Error","auth")
		return errors.New("shadow file is not exist")
	}
	asize=0
	//检测是否init过
	if Asize.Size()==0{
		if err=setAuthLen(0);err!=nil{
			tlog.Common("auth file init error","Error","auth")
			tlog.Err_in("auth file init error"+err.Error(),"Error","auth")
			return err
		}
		if Ssize.Size() == 0{
			//该软件还没有初始化
			return nil
		}
	}
	if Ssize.Size()==0{
		tlog.Log_all("文件数据不匹配，密码文件为空，数据缺失","Error","authdb")
		return errors.New("no use")
	}
	asize,err=getAuthLen()
	if err!=nil{
		tlog.Log_all("auth file get length error","Error","authdb")
		return err
	}
	usrtable:=make([]UsrItem,max(80,asize+40))
	
	
}