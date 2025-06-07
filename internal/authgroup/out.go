package authgroup

import (
	
)

func Init() error {
	//用户组缓存载入
	if err:=init_authdb();err!=nil{
		return err
	}
	

	return nil
}

func Kill() error {
	//用户组缓存释放
	if err:=kill_authdb();err!=nil{
		return err
	}
	return nil
}