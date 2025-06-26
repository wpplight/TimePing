package db

import "timeping/pkg/ostools"

//data文件 pagesize=2kb

func createData() error{
	p:=path+"data.tp"
	err:=ostools.CreateFile(p)
	if err!=nil{
		return err
	}
	return nil
}