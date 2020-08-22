package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/reciideo-lms/lecturer/platform/lecturer"
	"net/http"
)

type lecturerPostRequest struct {
	Forename    string                         `json:"forename"`
	Surname     string                         `json:"surname"`
	Description string                         `json:"description,omitempty"`
	Platforms   []lecturerPlatformsPostRequest `json:"platforms,omitempty"`
}

type lecturerPlatformsPostRequest struct {
	Platform string `json:"platform"`
	URL      string `json:"url"`
}

func LecturerPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := lecturerPostRequest{}

		err := c.BindJSON(&requestBody)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var platforms []lecturer.Platform
		for _, platform := range requestBody.Platforms {
			platform := lecturer.Platform{
				Platform: platform.Platform,
				URL:      platform.URL,
			}
			platforms = append(platforms, platform)
		}

		item := lecturer.Lecturer{
			Forename:    requestBody.Forename,
			Surname:     requestBody.Surname,
			Description: requestBody.Description,
			Platforms:   platforms,
		}

		result, err := lecturer.Add(item)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, result)
		}
	}
}
