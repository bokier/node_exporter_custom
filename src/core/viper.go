package core

/*
	viper 文件处理模块
*/
import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"node_exporter_custom/src/bash"
)

// InitViper 初始化Viper， 参数为 config.yaml 地址
func InitViper(filepath string) (err error) {
	viper.SetConfigFile(filepath)
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}
	// 把读取到的配置信息反序列化到 Conf 变量中
	if err := viper.Unmarshal(&bash.Conf); err != nil {
		fmt.Printf("[error]  InitViper().viper.Unmarshal failed, err:%v\n", err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("[system info] config is be changed ...")
		if err := viper.Unmarshal(&bash.Conf); err != nil {
			fmt.Printf("[error] InitViper().viper.Unmarshal failed, err:%v\n", err)
		}
	})
	return
}
