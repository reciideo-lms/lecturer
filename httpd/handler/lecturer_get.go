package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/reciideo-lms/lecturer/platform/lecturer"
	"net/http"
)

func LecturerGet(repo *lecturer.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := repo.GetAll()
		c.JSON(http.StatusOK, results)
	}
}
