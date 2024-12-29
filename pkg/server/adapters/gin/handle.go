package adapter

import (
	"net/http"
	"reflect"

	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server/errors"
	"github.com/gin-gonic/gin"
)

func handle[R any, B any, P any, Q any](status int, middlewares []gin.HandlerFunc, request server.Request[R, B, P, Q], validateBody bool) []gin.HandlerFunc {
	handlers := make([]gin.HandlerFunc, 0, len(middlewares)+2)
	for _, middleware := range middlewares {
		handlers = append(handlers, gin.HandlerFunc(middleware))
	}

	handlers = append(handlers, func(context *gin.Context) {
		abortValidationError := func(err error) {
			context.Error(errors.NewHttpError(http.StatusBadRequest, "Validation Error", err.Error()))
			context.Abort()
		}

		var body B
		if validateBody {
			if err := context.ShouldBindJSON(&body); err != nil {
				abortValidationError(err)

				return
			}
		}

		var params P
		if paramsType := reflect.TypeOf(params); paramsType != nil {
			paramsSchema := server.TransformDataToSchema(params)
			if err := context.ShouldBindUri(paramsSchema); err != nil {
				abortValidationError(err)

				return
			}

			params = server.TransformSchemaToData[P](paramsSchema)
		}

		var query Q
		if queryType := reflect.TypeOf(query); queryType != nil {
			querySchema := server.TransformDataToSchema(query)
			if err := context.ShouldBindQuery(querySchema); err != nil {
				abortValidationError(err)

				return
			}

			query = server.TransformSchemaToData[Q](querySchema)
		}

		result, err := request(&server.RequestOptions[B, P, Q]{
			Body:   body,
			Params: params,
			Query:  query,
			Ctx:    context.Request.Context(),
		})

		if err != nil {
			context.Error(err)
			context.Abort()

			return
		}

		context.JSON(status, result)
	})

	return handlers
}
