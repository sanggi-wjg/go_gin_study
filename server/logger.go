package server

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var logger = logrus.Logger{
	Out:          os.Stderr,
	Formatter:    &logrus.TextFormatter{},
	Level:        logrus.DebugLevel,
	ExitFunc:     os.Exit,
	ReportCaller: false,
	// Hooks:        hooks,
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Stop timer
		end := time.Now()
		latency := end.Sub(start)

		// clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()

		if raw != "" {
			path = path + "?" + raw
		}

		// Use debug level to keep production logs clean.
		logger.Debugf("http: %s %s (%3d) [%v]", method, path, statusCode, latency)
	}
}
