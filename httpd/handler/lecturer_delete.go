package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/reciideo-lms/lecturer/platform/lecturer"
	"net/http"
)

func LecturerDelete(repo *lecturer.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		err := repo.Delete(id)
		if err == nil {
			c.Writer.WriteHeader(http.StatusNoContent)
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
	}
}
