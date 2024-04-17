package main

import (
	"github.com/gin-gonic/gin"
	"node_exporter_custom/src/router"
)

func main() {
	r := gin.Default()
	r.GET("/health", router.CheckHealth)
	r.GET("/metrics", router.PrometheusHandler())
	_ = r.Run(":18080")
}
