package bash

var Conf *Config

type Config struct {
	Name          string            `mapstructure:"name"`
	Port          int               `mapstructure:"port"`
	NodeName      string            `mapstructure:"nodeName"`
	NodeIp        string            `mapstructure:"nodeIp"`
	Version       string            `mapstructure:"version"`
	MetricsEnable bool              `mapstructure:"metricsEnable"`
	Metrics       map[string]string `mapstructure:"metrics"`
}

type MetricsConfig struct {
}
