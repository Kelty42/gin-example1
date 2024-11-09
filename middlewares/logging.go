package middlewares

import (
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
    "fmt"
    "time"
)

func JSONLoggerMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        startTime := time.Now()

        // Обработка запроса
        c.Next()

        // Вычисление времени выполнения
        latency := time.Since(startTime)

        // Преобразуем латенцию в секунды
        latencyInSeconds := latency.Seconds()

        // Логирование в формате JSON
        logrus.WithFields(logrus.Fields{
            "status":     fmt.Sprintf("%d", c.Writer.Status()), // Преобразуем статус в строку
            "method":     c.Request.Method,
            "path":       c.Request.URL.Path,
            "ip":         c.ClientIP(),
            "latency":    latencyInSeconds, // Латенция в секундах
            "user-agent": c.Request.UserAgent(),
        }).Info("request completed")
    }
}
