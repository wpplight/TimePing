package engine

import (
	"os"
	"os/signal"
	"syscall"
	"timeping/internal/config"
	"timeping/internal/tlog"
)

var (
	run_chan chan os.Signal
)

func Initial_engine() error {
	//初始化日志,后面模块报错都会用到日志
	if err := tlog.Init_log(); err != nil {
		return err
	}
	//初始化config模块，读取参数，后面模块都需要参数进行参考
	if err := config.Load_setting(); err != nil {
		return err
	}
	
	run_chan = make(chan os.Signal, 1)
	signal.Notify(run_chan, syscall.SIGINT, syscall.SIGTERM)
	return nil
}

func Run() {
	tlog.Common("start successfully 启动成功", "engine")
	<-run_chan
	tlog.Common("Exit 已停止","engine")
	tlog.Exit_log()
}
