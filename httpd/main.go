package main

import (
	"github.com/gin-gonic/gin"
	"github.com/reciideo-lms/lecturer/config"
	"github.com/reciideo-lms/lecturer/httpd/handler"
	"github.com/reciideo-lms/lecturer/platform/lecturer"
	"log"
)

func main() {
	r := gin.Default()

	err := config.InitDatabase()
	if err != nil {
		log.Panic(err)
	}

	err = lecturer.New()
	if err != nil {
		log.Panic(err)
	}

	r.GET("/health", handler.HealthGet())

	l := r.Group("/lecturer")
	l.GET("/", handler.LecturerGet())
	l.GET("/:id", handler.LecturerGetSingle())
	l.POST("/", handler.LecturerPost())
	l.DELETE("/:id", handler.LecturerDelete())

	err = r.Run()
	if err != nil {
		log.Panic(err)
	}
}
