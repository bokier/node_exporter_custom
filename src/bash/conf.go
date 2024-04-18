package bash

type Config struct {
	Name          string `mapstructure:"name"`
	Port          int    `mapstructure:"port"`
	NodeName      string `mapstructure:"nodeName"`
	NodeIp        string `mapstructure:"nodeIp"`
	Version       string `mapstructure:"version"`
	MetricsEnable bool   `mapstructure:"metricsEnable"`
}

var Conf *Config
