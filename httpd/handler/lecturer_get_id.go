package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/reciideo-lms/lecturer/platform/lecturer"
	"net/http"
)

func LecturerGetSingle(repo *lecturer.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		result, err := repo.GetSingle(id)
		if err != nil && err.Error() == "NotFound" {
			c.AbortWithStatus(http.StatusNotFound)
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, result)
		}
	}
}
