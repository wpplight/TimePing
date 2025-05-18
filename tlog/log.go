package tlog

import (
	"fmt"
	"timeping/ostools"
	log "github.com/sirupsen/logrus"
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
func Err_in(e string){
	f,err:=ostools.OpenFile("./error.log")
	if err!=nil{
		fmt.Println("error.log 请检查权限")
		return
	}
	defer f.Close()
	log.SetOutput(f)
	log.SetLevel(log.ErrorLevel)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	log.Error(e)
}