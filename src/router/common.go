package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CheckHealth 健康检测
func CheckHealth(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
