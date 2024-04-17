package router

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
)

type handler struct {
	unfilteredHandler       http.Handler
	exporterMetricsRegistry *prometheus.Registry
}

func newHandler() *handler {
	h := &handler{
		exporterMetricsRegistry: prometheus.NewRegistry(),
	}
	return h
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

func PrometheusHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		newHandler().ServeHTTP(ctx.Writer, ctx.Request)
	}
}
