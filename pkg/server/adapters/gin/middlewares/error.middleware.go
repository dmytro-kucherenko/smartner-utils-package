package middlewares

import (
	"net/http"

	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/log"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server/dtos"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server/errors"
	"github.com/gin-gonic/gin"
)

func Error() gin.HandlerFunc {
	logger := log.New("Error Middleware")

	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Fatal("Panic Recovery:", err)

				context.JSON(http.StatusInternalServerError, &dtos.ErrorResponse{
					Status:  http.StatusInternalServerError,
					Message: "Internal Server Error",
				})
			}
		}()

		context.Next()

		if err := context.Errors.Last(); err != nil {
			if httpErr, ok := any(err.Err).(*errors.HttpError); ok {
				context.JSON(httpErr.Status(), &dtos.ErrorResponse{
					Status:  httpErr.Status(),
					Message: httpErr.Error(),
					Details: httpErr.Details(),
				})
			} else {
				context.JSON(http.StatusInternalServerError, &dtos.ErrorResponse{
					Status:  http.StatusInternalServerError,
					Message: "Internal Server Error",
				})
			}

			return
		}

		if !context.Writer.Written() {
			context.JSON(http.StatusNotImplemented, &dtos.ErrorResponse{
				Status:  http.StatusNotImplemented,
				Message: "Method did not have response.",
			})
		}
	}
}
