package engine

import (
	"container/list"
	"os"
	"os/signal"
	"syscall"
	"timeping/config"
	"timeping/task"
	"timeping/tlist"
	"timeping/tlog"
)
var(
	run_chan chan os.Signal
	Engine_kernal struct{
		TaskPool *list.List
		UnuseQueue *tlist.Tlist
	}
)

func Initial_engine() error {
	if err:=config.Get_setting();err!=nil{
		return err
	}
	Engine_kernal.TaskPool,Engine_kernal.UnuseQueue=task.InitialTaskPool()
	if err:=tlog.Init_log();err!=nil{
		return err
	}
	run_chan = make(chan os.Signal, 1)
	signal.Notify(run_chan, syscall.SIGINT, syscall.SIGTERM)
	return nil
}

func Run(){
	tlog.Common("start successfully 启动成功","engine")
	<- run_chan
	tlog.Common("Exit 已停止")
}