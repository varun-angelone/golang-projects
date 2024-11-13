package handlers

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
	"os"
)

var (
	region     string
	accessKey  string
	secretKey  string
	bucketName string
	uploader   *s3manager.Uploader
)

func init() {
	// init env variables
	region = os.Getenv("AWS_REGION")
	accessKey = os.Getenv("AWS_ACCESS_KEY")
	secretKey = os.Getenv("AWS_SECRET_KEY")
	bucketName = os.Getenv("AWS_BUCKET_NAME")

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

func VideoUploadHandler(c *gin.Context) {
	formID := c.PostForm("form_id") // to use form_id for mp4 file name
	if formID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Form ID is required"})
		return
	}

	header, err := c.FormFile("video")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to read the video file"})
		return
	}

	// Open the file for uploading to S3
	videoFile, err := header.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to open video file"})
		return
	}
	defer videoFile.Close()

	// Generate unique filename and upload
	fileName := formID + "/selfie/" + header.Filename
	s3URL, err := uploadFileToS3(videoFile, fileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload video to S3"})
		return
	}

	// Success response with S3 URL
	c.JSON(http.StatusAccepted, gin.H{
		"message":  "Video uploaded successfully",
		"file_url": s3URL,
	})
}

func uploadFileToS3(file multipart.File, fileName string) (string, error) {
	// Create an uploader with S3 client
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(fileName),
		Body:        file,
		ContentType: aws.String("video/mp4"), // Adjust as needed
	})
	if err != nil {
		return "", err
	}
	// Construct the S3 file URL
	return fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucketName, fileName), nil
}
