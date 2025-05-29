package config

import (
	"fmt"
	"os"
	"timeping/tlog"
	"github.com/spf13/viper"
)

func init_config(path string) error {
	file, err := os.Create(path)
	if err != nil {
		tlog.Common("Setting creat default! Check your root setting文件创造失败", "config")
		return fmt.Errorf("setting creat default! Check your root")
	}
	defer file.Close()
	//文件名称
	viper.SetConfigName("timecnf")
	// 设置配置文件类型
	viper.SetConfigType("yaml")
	// 设置配置文件的路径（这里用当前目录）
	viper.AddConfigPath(".")
	viper.Set("TaskPoolSize", 100)
	viper.Set("TimeWheelSize", 60)
	viper.Set("Timeinterval", 100)
	viper.Set("Port", 9768)
	err = viper.WriteConfigAs(path)
	if err != nil {
		return fmt.Errorf("setting creat default! Check your root 创建setting文件错误")
	}
	tlog.Common("init successful 初始化完成", "config")
	return fmt.Errorf("初始化完成")
}
