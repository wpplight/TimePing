package tlog

import (
	"fmt"
	"os"
	"timeping/ostools"

	log "github.com/sirupsen/logrus"
)

var(
	logfile *os.File
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
func Exit_log(){
	logfile.Close()
}
func Err_in(e string){
	var err error
	logfile,err=ostools.OpenFile("./error.log")
	if err!=nil{
		fmt.Println("error.log 请检查权限")
		return
	}
	log.SetOutput(logfile)
	log.SetLevel(log.ErrorLevel)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	log.Error(e)
}