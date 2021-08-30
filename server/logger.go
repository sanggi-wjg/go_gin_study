package server

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func InitLogger(debug bool) {
	Log = &logrus.Logger{
		Out:          os.Stdout,
		Formatter:    &logrus.TextFormatter{},
		Level:        logrus.DebugLevel,
		ExitFunc:     os.Exit,
		ReportCaller: false,
	}

	if debug {
		Log.SetLevel(logrus.DebugLevel)
		Log.SetFormatter(&logrus.TextFormatter{
			DisableColors: false,
			FullTimestamp: true,
		})
	} else {
		//Log.SetOutput()
		Log.SetLevel(logrus.InfoLevel)
		Log.SetFormatter(&logrus.JSONFormatter{})
	}
}

func Debug(v ...interface{}) {
	Log.Debugln(v)

}
func Info(v ...interface{}) {
	Log.Infoln(v)
}

func Warn(v ...interface{}) {
	Log.Warnln(v)
}

func Error(v ...interface{}) {
	Log.Errorln(v)
}

func Fatal(v ...interface{}) {
	Log.Fatalln(v)
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
		Log.Debugf("http: %s %s (%3d) [%v]", method, path, statusCode, latency)
	}
}
