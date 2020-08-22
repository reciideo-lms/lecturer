package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/reciideo-lms/lecturer/platform/lecturer"
	"net/http"
)

func LecturerGetSingle() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		item, err := lecturer.GetSingle(id)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.JSON(http.StatusOK, item)
		}
	}
}
