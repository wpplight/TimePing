package authgroup

import (
	"os"
	"timeping/internal/tlog"
)

//当用户总文件没有被构造时候创造默认用户组总文件
func create_allauth() error { 
	var err error
	if auall,err=os.Create("./auth/auth.tp");err!=nil{
			tlog.Common("create authdb file fail","Error","auth")
			tlog.Err_in("create authdb file fail","Error","auth")
			return err
	}
	auall.Close()
	auall,err=os.OpenFile("./auth/auth.tp",os.O_APPEND,0644)
	if err!=nil{
		auall.Close()
		tlog.Common("create authdb file fail","Error","auth")
		tlog.Err_in("create authdb file fail","Error","auth")
		return err
	}
	

	return nil
}