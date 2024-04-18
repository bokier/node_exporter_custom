package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"node_exporter_custom/src/bash"
)

// CheckHealth 健康检测
func CheckHealth(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

// CheckVersion 获取版本
func CheckVersion(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"version": bash.Conf.Version,
	})
}
