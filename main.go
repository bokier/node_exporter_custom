package main

import (
	"github.com/gin-gonic/gin"
	"node_exporter_custom/src/router"
)

func main() {
	r := gin.New()

	r.GET("/health", router.CheckHealth)
	_ = r.Run(":18080")
}
