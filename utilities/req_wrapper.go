package utilities

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func HandleRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		log.Printf("Receive request - [%s %s]", c.Request.Method, c.Request.URL)
		c.Set("service_name", "shopaholic_API_service")
		c.Next()

		latency := time.Since(t)
		status := c.Writer.Status()
		log.Printf("Return response - [%s %s] with status %v, took %v", c.Request.Method, c.Request.URL, status, latency)
	}
}
