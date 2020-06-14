package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/reciideo-lms/lecturer/platform/lecturer"
	"net/http"
)

func LecturerGet(repo *lecturer.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		results, err := repo.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, results)
		}
	}
}
