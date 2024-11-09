package handlers

import (
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
)

func NoRouteHandler(c *gin.Context) {
    logrus.WithFields(logrus.Fields{
        "status":     "404",
        "method":     c.Request.Method,
        "path":       c.Request.URL.Path,
        "ip":         c.ClientIP(),
        "latency":    0,
        "user-agent": c.Request.UserAgent(),
    }).Error("404 page not found")

    c.JSON(404, gin.H{"error": "page not found"})
}
