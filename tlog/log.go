package tlog

import (
	"fmt"
	"timeping/ostools"
)


func Init_log() error{
	if !ostools.FileExists("./error.log"){
		if err:=ostools.CreateFile("./error.log");err!=nil{
			fmt.Println("err.log 请检查权限")
			return err
		}
	}
	return nil
}
func Err_in(){
	
}