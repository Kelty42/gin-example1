package handlers

import "github.com/gin-gonic/gin"

func APIHandler(c *gin.Context) {
    c.JSON(200, gin.H{"message": "api endpoint"})
}
