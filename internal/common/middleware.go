package common

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// logs the request body for POST, PUT, and PATCH requests
func LogRequestBodyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Only log for specific methods if needed
		if c.Request.Method != "POST" && c.Request.Method != "PUT" && c.Request.Method != "PATCH" {
			c.Next()
			return
		}

		// Read the body
		bodyBytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Unable to read request body"})
			return
		}

		// Restore the io.ReadCloser to its original state
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		// Log the body (here we use Gin's logger)
		log.Printf("Request Body: %s", string(bodyBytes))

		// Continue to the next handler
		c.Next()
	}
}
