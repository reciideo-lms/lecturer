package main

import (
	"github.com/gin-gonic/gin"
	"github.com/reciideo-lms/lecturer/httpd/handler"
	"log"
)

func main() {
	r := gin.Default()

	r.GET("/health", handler.HealthGet())

	err := r.Run()
	if err != nil {
		log.Panic(err)
	}
}
