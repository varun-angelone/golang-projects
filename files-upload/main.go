package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	r := gin.Default()
	r.POST("/upload", handleUpload)
	if err := r.Run(":8080"); err != nil {
		log.Fatalln("Error in Server", err)
	}
}

func handleUpload(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	files := form.File["files"]
	for _, file := range files {
		err := saveFile(file)
		if err != nil {
			fmt.Println("error in save file", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "Files uploaded successfully"})
}

func saveFile(file *multipart.FileHeader) error {
	dst := "./destination"
	fmt.Println(file.Filename)
	src, err := file.Open()
	if err != nil {
		fmt.Println("error in file destination", err)
		return err
	}
	defer src.Close()

	// ensure destination directory exists
	if err := os.MkdirAll(dst, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}

	// create a new file in the described destination folder
	dstPath := filepath.Join(dst, file.Filename)
	dstn, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dstn.Close()
	return nil
}
