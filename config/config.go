package config

import (
	"fmt"
	"timeping/ostools"
	"timeping/tlog"
	"timeping/global"
	"github.com/spf13/viper"
)

func Get_setting() error {
	path:="./timecnf.yaml"
	if !ostools.FileExists(path){
		if err:=init_config(path);err!=nil{
			fmt.Println(err)
			return err
		}
	}
	global.Conf.TaskPoolSize=uint16(viper.GetInt("TaskPoolSize"))
	global.Conf.TimeWheelSize=uint16(viper.GetInt("TimeWheelSize"))
	global.Conf.Timeinterval=uint16(viper.GetInt("Timeinterval"))
	global.Conf.Port=uint16(viper.GetInt("Port"))
	tlog.Common("Setting is OK 配置读取成功","Setting")
	return nil
}

