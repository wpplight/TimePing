package engine

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"timeping/config"
	"timeping/task"
	"timeping/tlog"
)
var(
	run_chan chan os.Signal
)

func Initial_engine() error {
	if err:=config.Get_setting();err!=nil{
		return err
	}
	task.InitialTaskPool()
	if err:=tlog.Init_log();err!=nil{
		return err
	}
	run_chan = make(chan os.Signal, 1)
	signal.Notify(run_chan, syscall.SIGINT, syscall.SIGTERM)
	return nil
}

func Run(){
	<- run_chan
	fmt.Println("TimePing Exit")
}