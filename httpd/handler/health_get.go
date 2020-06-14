package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "up",
		})
	}
}
