package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func APIHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "api endpoint"})
}
