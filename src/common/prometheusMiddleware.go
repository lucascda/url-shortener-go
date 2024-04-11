package common

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	RequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of http requests",
		},
		[]string{"method", "endpoint"},
	)
)

func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		path := c.FullPath()
		method := c.Request.Method
		RequestsTotal.WithLabelValues(method, path).Inc()
	}

}
