package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JSONSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": data})
}

func JSONError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"status": "error", "message": message})
}
