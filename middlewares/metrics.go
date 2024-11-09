package middlewares

import (
    "github.com/gin-gonic/gin"
    "github.com/prometheus/client_golang/prometheus"
)

var (
    httpRequestsTotal = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Счетчик всех HTTP-запросов по методам и путям",
        },
        []string{"method", "path"},
    )
    httpRequestDuration = prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Name: "http_request_duration_seconds",
            Help: "Гистограмма для времени выполнения запросов по методам и путям",
        },
        []string{"method", "path"},
    )
)

func init() {
    prometheus.MustRegister(httpRequestsTotal, httpRequestDuration)
}

// MetricsMiddleware регистрирует количество запросов и их время выполнения
func MetricsMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        timer := prometheus.NewTimer(httpRequestDuration.WithLabelValues(c.Request.Method, c.FullPath()))
        defer timer.ObserveDuration()

        httpRequestsTotal.WithLabelValues(c.Request.Method, c.FullPath()).Inc()

        c.Next()
    }
}
