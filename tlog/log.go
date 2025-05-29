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
	if !ostools.FileExists("./timeping.log"){
		if err:=ostools.CreateFile("./timeping.log");err!=nil{
			fmt.Println("err.log 请检查权限")
			return err
		}

	}
	var err error
	logfile,err=ostools.OpenFile("./timeping.log")
	if err!=nil{
		fmt.Println("timeping.log 请检查权限")
		return fmt.Errorf("timeping.log 请检查权限")
	}
	return nil
}
func Exit_log(){
	logfile.Close()
}
func Err_in(e string, tag ...string){
	log.SetOutput(logfile)
	log.SetLevel(log.ErrorLevel)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	var out string 
	for _,v:=range tag{
		out+=v;
	}
	out+=e;
	log.Error(e)
}