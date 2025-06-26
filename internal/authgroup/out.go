package authgroup

import (
	"timeping/internal/authgroup/db"
)

func Init(Fpath string)  error {
	var err error
	//用户组缓存载入
	if err=db.Init_authdb(Fpath);err!=nil{
		return err
	}
	initusr()

	return nil
}

func Kill()  {
	//用户组缓存释放
	db.Release()

}