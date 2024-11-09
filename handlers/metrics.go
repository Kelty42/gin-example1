package handlers

import "github.com/gin-gonic/gin"

// Пример для маршрута /metrics
func MetricsHandler(c *gin.Context) {
    c.JSON(200, gin.H{"message": "metrics endpoint"})
}
