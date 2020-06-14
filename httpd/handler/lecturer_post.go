package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/reciideo-lms/lecturer/platform/lecturer"
	"net/http"
)

type lecturerPostRequest struct {
	Forename string `json:"forename"`
	Surname  string `json:"surname"`
}

func LecturerPost(repo *lecturer.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := lecturerPostRequest{}

		err := c.BindJSON(&requestBody)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		item := lecturer.Lecturer{
			Forename: requestBody.Forename,
			Surname:  requestBody.Surname,
		}
		repo.Add(item)

		c.JSON(http.StatusOK, item)
	}
}
