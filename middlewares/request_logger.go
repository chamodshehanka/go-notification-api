package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

func Logger() gin.HandlerFunc {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true, // Enable colors
		FullTimestamp: true,
	})

	return func(c *gin.Context) {
		start := time.Now()

		// Process request
		c.Next()

		// Log request
		end := time.Now()
		latency := end.Sub(start)
		clientIP := c.ClientIP()
		method := c.Request.Method
		path := c.Request.URL.Path
		statusCode := c.Writer.Status()
		logEntry := fmt.Sprintf("| %3d | %13v | %15s | %-7s %s",
			statusCode,
			latency,
			clientIP,
			method,
			path,
		)
		if statusCode >= 400 {
			logger.Error(logEntry)
		} else {
			logger.Info(logEntry)
		}
	}
}
