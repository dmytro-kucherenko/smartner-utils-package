package interceptors

import (
	"net/http"

	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/log"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server/dtos"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server/errors"
	"github.com/gin-gonic/gin"
)

func Error() gin.HandlerFunc {
	logger := log.New("error middleware")

	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error("panic recovery:", err)

				context.JSON(http.StatusInternalServerError, &dtos.ErrorItem{
					Status:  http.StatusInternalServerError,
					Message: "internal server error",
				})
			}
		}()

		context.Next()

		if err := context.Errors.Last(); err != nil {
			if httpErr, ok := any(err.Err).(*errors.HttpError); ok {
				context.JSON(httpErr.Status(), &dtos.ErrorItem{
					Status:  httpErr.Status(),
					Message: httpErr.Error(),
					Details: httpErr.Details(),
				})
			} else {
				logger.Error(err.Err.Error())
				context.JSON(http.StatusInternalServerError, &dtos.ErrorItem{
					Status:  http.StatusInternalServerError,
					Message: "internal server error",
				})
			}

			return
		}

		if !context.Writer.Written() {
			context.JSON(http.StatusNotImplemented, &dtos.ErrorItem{
				Status:  http.StatusNotImplemented,
				Message: "method did not have response.",
			})
		}
	}
}
