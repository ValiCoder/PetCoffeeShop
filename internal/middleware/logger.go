package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger — middleware for logging updates to console
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Pass control to the next handler
		c.Next()

		// Calculate execution time after the handler has finished
		latency := time.Since(t)
		status := c.Writer.Status()
		path := c.Request.URL.Path
		method := c.Request.Method

		log.Printf("[API-LOG] %s | %d | %s | %v", method, status, path, latency)
	}
}
