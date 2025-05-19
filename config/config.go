package config

import (
	"fmt"
	"timeping/ostools"
	"timeping/tlog"
	"timeping/utype"

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
	utype.Conf.TaskPoolSize=uint16(viper.GetInt("TaskPoolSize"))
	utype.Conf.TimeWheelSize=uint16(viper.GetInt("TimeWheelSize"))
	utype.Conf.Timeinterval=uint16(viper.GetInt("Timeinterval"))
	utype.Conf.Port=uint16(viper.GetInt("Port"))
	tlog.Common("Setting is OK 配置读取成功","Setting")
	return nil
}

