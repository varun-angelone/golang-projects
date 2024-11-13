// handlers/form.go
package handlers

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	formIDCounter int = 1
	formDataStore     = make(map[string]map[string]string)
	mu            sync.Mutex
)

// serially generate formID
func generateUniqueFormID() string {
	mu.Lock()
	defer mu.Unlock()
	formID := fmt.Sprintf("form_%d", formIDCounter)
	formIDCounter++
	return formID
}

func saveFormData(formID string, data map[string]string) {
	mu.Lock()
	defer mu.Unlock()
	formDataStore[formID] = data
}

func SubmitFormHandler(c *gin.Context) {
	now := time.Now()
	name := c.PostForm("name")
	email := c.PostForm("email")

	// generate a unique form ID and save form data
	formID := generateUniqueFormID()
	saveFormData(formID, map[string]string{"name": name, "email": email})

	c.JSON(http.StatusOK, gin.H{
		"message": "Form submitted successfully",
		"formID":  formID,
	})
	fmt.Println("Form Submitted Successfully", formID)
	fmt.Println("Time taken=", time.Since(now))
}
