package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"node_exporter_custom/src/collector"
	"node_exporter_custom/src/core"
	"node_exporter_custom/src/router"
)

func Router() {
	r := gin.Default()
	r.GET("/health", router.CheckHealth)
	r.GET("/version", router.CheckVersion)
	r.GET("/metrics", router.PrometheusHandler())
	_ = r.Run(":18080")
}

func init() {
	err := core.InitViper("./custom.yaml")
	if err != nil {
		return
	}
	fmt.Println("[init] init viper..")
}

func main() {
	collector.InitBashCollector()
	//fmt.Println("version:", bash.Conf.Version)
	Router()
}
