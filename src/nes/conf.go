package bash

var Conf *Config

type Config struct {
	Name       string            `mapstructure:"name"`
	Port       string            `mapstructure:"port"`
	NodeName   string            `mapstructure:"nodeName"`
	NodeIp     string            `mapstructure:"nodeIp"`
	Version    string            `mapstructure:"version"`
	BashEnable bool              `mapstructure:"bashEnable"` // 是否开启bash指标获取
	ApiEnable  bool              `mapstructure:"apiEnable"`  // 是否开启api指标获取
	Metrics    map[string]string `mapstructure:"metrics"`
}
