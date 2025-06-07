package config

import (
	"timeping/internal/tlog"
	"timeping/pkg/ostools"
	"github.com/spf13/viper"
)

func Load_setting() error {
	path:="./timecnf.yaml"
	if !ostools.FileExists(path){
		if err:=init_config(path);err!=nil{
			return err
		}
	}
	
	Conf.TaskPoolSize=uint16(viper.GetInt("TaskPoolSize"))
	Conf.TimeWheelSize=uint16(viper.GetInt("TimeWheelSize"))
	Conf.Timeinterval=uint16(viper.GetInt("Timeinterval"))
	Conf.Timelevel=viper.GetInt("Timelevel")
	Conf.Port=uint16(viper.GetInt("Port"))

	tlog.Common("Setting is OK 配置读取成功","Setting")
	return nil
}

