package main

import (
	"github.com/gin-gonic/gin"
	"yourprojectname/internal/handler"
)

func main() {
	router := gin.Default()

	// Serve the static files for your website
	router.Static("/", "./public")

	// Register the form handler
	router.POST("/submit", handler.ProcessForm)

	router.Run(":8080")
}
