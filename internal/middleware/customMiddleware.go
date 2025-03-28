package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func CustomLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Custom Logger")
		t := time.Now()
		// before request
		c.Next()
		// after request
		latency := time.Since(t)
		log.Print(latency)
		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}
