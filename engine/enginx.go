package engine

import (
	"timeping/config"
	"timeping/task"
	"timeping/tlog"
)

func Initial_engine() error {
	if err:=config.Get_setting();err!=nil{
		return err
	}
	task.InitialTaskPool()
	if err:=tlog.Init_log();err!=nil{
		return err
	}

	return nil
}
