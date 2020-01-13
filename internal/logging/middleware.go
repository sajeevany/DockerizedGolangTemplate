package logging

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

const LoggerKey = "logger"

//SetCtxLogger - Sets a logger as defined by LoggerKey
func SetCtxLogger(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(LoggerKey, logger)
		c.Next()
	}
}

//LogRequest - Logs requests
func LogRequest(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process request
		c.Next()

		// Get finish time
		end := time.Now()

		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		method := c.Request.Method
		statusCode := c.Writer.Status()
		duration := end.Sub(start)
		timeFormatted := end.Format("2006-01-02 15:04:05")

		msg := fmt.Sprintf("[%v] %v (%v) %v %v %v", timeFormatted, method, path, statusCode, duration, raw)
		logger.WithFields(logrus.Fields{
			"endTime":  timeFormatted,
			"method":   method,
			"path":     path,
			"status":   statusCode,
			"duration": duration,
			"raw":      raw,
		}).Info(msg)
	}
}
