package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
)

const (
	region     = "us-east-1"
	accessKey  = ""
	secretKey  = ""
	bucketName = ""
)

var uploader *s3manager.Uploader

func init() {
	// init aws session
	awsSession, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String(region),
			Credentials: credentials.NewStaticCredentials(
				accessKey,
				secretKey,
				"",
			),
		},
	})
	if err != nil {
		panic("error in awsSession initialisation" + err.Error())
	}

	uploader = s3manager.NewUploader(awsSession)
}

func main() {
	r := gin.Default()
	r.POST("/upload", uploadFile)
	r.Run(":8080")
}

func uploadFile(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		panic("file upload api" + err.Error())
	}
	var errors []string
	var uploadedURLs []string
	files := form.File["files"]
	for _, file := range files {
		fileHeader := file
		f, err := fileHeader.Open()
		if err != nil {
			errors = append(errors, err.Error())
			continue
		}
		defer f.Close()
		uploadedURL, err := saveFile(f, fileHeader)
		if err != nil {
			errors = append(errors, err.Error())
		} else {
			uploadedURLs = append(uploadedURLs, uploadedURL)
		}
	}
	if len(errors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": strings.Join(errors, ",")})
	} else {
		c.JSON(http.StatusOK, gin.H{"urls": strings.Join(uploadedURLs, ",")})
	}

}

func saveFile(f io.Reader, fileHeader *multipart.FileHeader) (string, error) {
	// upload the file to s3 using fileReader
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileHeader.Filename),
		Body:   f,
	})
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucketName, fileHeader.Filename)
	return url, nil
}
