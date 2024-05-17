package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"node_exporter_custom/src/collector"
	"node_exporter_custom/src/core"
	nes "node_exporter_custom/src/nes"
	"node_exporter_custom/src/router"
)

func init() {
	err := core.InitViper("./custom.yaml")
	if err != nil {
		return
	}
	fmt.Println("[init] init viper..")
}

func NoLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 直接将请求传递到下一个处理器
		c.Next()
	}
}

func Router() {
	r := gin.New()

	// 设置不记录访问日志, 需要配合gin.New和middleware使用
	r.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: []string{"/metrics", "/health", "/version"},
	}))
	r.Use(gin.Recovery())

	r.Use(NoLogMiddleware())
	{
		r.GET("/metrics", router.PrometheusHandler())
		r.GET("/health", router.CheckHealth)
		r.GET("/version", router.CheckVersion)
	}

	runPort := ":" + nes.Conf.Port
	_ = r.Run(runPort)
}

func main() {

	if nes.Conf.BashEnable {
		collector.InitBashCollector()
	}
	Router()
}
