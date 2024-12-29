package middlewares

import (
	"fmt"
	"time"

	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/log"
	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	logger := log.New("Route")

	return func(context *gin.Context) {
		start := time.Now()
		context.Next()
		end := time.Since(start)

		latency := float64(end.Microseconds()) / 1000.0
		status := context.Writer.Status()
		path := context.Request.URL.Path
		method := context.Request.Method

		fields := map[string]any{}
		if len(context.Errors) == 1 {
			fields["Error"] = context.Errors.Last().Error()
		} else if len(context.Errors) >= 2 {
			for index, err := range context.Errors {
				fields[fmt.Sprint("Error#", index)] = err.Error()
			}
		}

		entry := logger.CreateEntry(fields)
		msg := fmt.Sprintf("[%v-%v] %v %vms\n", method, status, path, latency)

		switch {
		case status >= 500:
			entry.Error(msg)
		case status >= 400:
			entry.Warn(msg)
		default:
			entry.Info(msg)
		}
	}
}
