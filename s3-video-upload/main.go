package main

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/varun-angelone/golang-projects/s3-video-upload/handlers"
)

func main() {
	r := gin.New()
	// serve static files
	r.Static("/static", "./static")

	// endpoint
	r.POST("/submit-form", handlers.SubmitFormHandler)
	r.POST("/upload-video", handlers.VideoUploadHandler)
	r.Run(":8080")
}
