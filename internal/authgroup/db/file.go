package db

import (
	"errors"
	"os"
	"timeping/internal/tlog"
)

//顺序读取校验总auth文件,使用os read
func Init_authdb(Fpath string) (error) {
	

	//检验路径合法性
	if Fpath!="./" {
		if Fpath[len(Fpath)-1]=='/'{
			Fpath=Fpath[:len(Fpath)-1]
		}
		if info, err:=os.Stat(Fpath);os.IsNotExist(err){
			tlog.Common("文件路径不存在", "Error","AuthDb")
			tlog.Err_in("文件路径不存在"+err.Error(),"Error","AuthDb")
			return errors.New("file path not exist")
		}else{
			if !info.IsDir(){
				tlog.Common("文件路径不是文件夹", "Error","AuthDb")
				tlog.Err_in("文件路径不是文件夹"+err.Error(),"Error","AuthDb")
				return errors.New("file path not dir")
			}
		}
		Fpath+="/"
	}
	path=Fpath

	//句柄初始化
	auth=nil
	authindex=nil
	datafile=nil
	shadow=nil

	//校验文件是否存在
    err:=checkFileExit()
	if err!=nil{
		tlog.Common("持久化文件错误"+err.Error(),"Warning","auth")
		tlog.Err_in("缺少持久化文件，如果不是初次启动请校验数据完整性"+err.Error(),"Warning","auth")	
		return err
	}
	err=syncCheck()
	if err!=nil{
		tlog.Common("持久化文件错误"+err.Error(),"Warning","auth")
		tlog.Err_in("持久化文件错误"+err.Error(),"Warning","auth")
		return err
	}

    return nil
}

func Release() {
	if auth != nil {
		tlog.Common("释放auth资源错误，该资源未被初始化")
		tlog.Err_in("释放auth资源错误，该资源未被初始化,auth文件句柄传入错误","Warning","auth")
	}else{
		auth.Close()
	}
	if shadow != nil {
		tlog.Common("释放shadow资源错误，该资源未被初始化")
		tlog.Err_in("释放shadow资源错误，该资源未被初始化,shadow文件句柄传入错误","Warning","auth")
	}else{
		shadow.Close()
	}
	if datafile != nil {
		tlog.Common("释放datafile资源错误，该资源未被初始化")
		tlog.Err_in("释放datafile资源错误，该资源未被初始化,datafile文件句柄传入错误","Warning","auth")
	}else{
		datafile.Close()
	}
	if authindex != nil {
		tlog.Common("释放authindex资源错误，该资源未被初始化")
		tlog.Err_in("释放authindex资源错误，该资源未被初始化,authindex文件句柄传入错误","Warning","auth")
	}else{
		authindex.Close()
	}
}

