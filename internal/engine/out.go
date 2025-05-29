package engine

import (
	"os"
	"os/signal"
	"syscall"
	"timeping/internal/config"
	"timeping/internal/task"
	"timeping/internal/tlog"
)

var (
	run_chan chan os.Signal
)

func Initial_engine() error {
	if err := tlog.Init_log(); err != nil {
		return err
	}
	if err := config.Load_setting(); err != nil {
		return err
	}
	task.InitialTaskPool()

	run_chan = make(chan os.Signal, 1)
	signal.Notify(run_chan, syscall.SIGINT, syscall.SIGTERM)
	return nil
}

func Run() {
	tlog.Common("start successfully 启动成功", "engine")
	<-run_chan
	tlog.Common("Exit 已停止")
}
