package main

import (
	"github.com/gin-gonic/gin"
	"github.com/reciideo-lms/lecturer/httpd/handler"
	"github.com/reciideo-lms/lecturer/platform"
	"github.com/reciideo-lms/lecturer/platform/lecturer"
	"log"
)

func main() {
	r := gin.Default()
	db := platform.InitDB()

	repo := lecturer.New(db)

	r.GET("/health", handler.HealthGet())

	l := r.Group("/lecturer")
	l.GET("/", handler.LecturerGet(repo))
	l.POST("/", handler.LecturerPost(repo))

	err := r.Run()
	if err != nil {
		log.Panic(err)
	}
}
