package adapter

import (
	"net/http"
	"reflect"

	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/schema/modules/common"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server/errors"
	"github.com/gin-gonic/gin"
)

func handle[R any, B any, P any, Q any](
	request server.Request[R, B, P, Q],
	options *RequestConfig,
	validateBody bool,
) []gin.HandlerFunc {
	handlers := make([]gin.HandlerFunc, 0, len(options.Middlewares)+2)
	for _, middleware := range options.Middlewares {
		handlers = append(handlers, gin.HandlerFunc(middleware))
	}

	handlers = append(handlers, func(context *gin.Context) {
		if options.ProvideSession && options.Meta.Session == nil {
			context.Error(errors.NewHttpError(http.StatusUnauthorized, "User session was not found."))
			context.Abort()

			return
		}

		if options.Meta.Session == nil {
			options.Meta.Session = &server.Session{}
		}

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
			paramsSchema := common.ModifySchema(params)
			if err := context.ShouldBindUri(paramsSchema); err != nil {
				abortValidationError(err)

				return
			}

			params = common.ParseSchema[P](paramsSchema)
		}

		var query Q
		if queryType := reflect.TypeOf(query); queryType != nil {
			querySchema := common.ModifySchema(query)
			if err := context.ShouldBindQuery(querySchema); err != nil {
				abortValidationError(err)

				return
			}

			query = common.ParseSchema[Q](querySchema)
		}

		result, err := request(&server.RequestOptions[B, P, Q]{
			Body:    body,
			Params:  params,
			Query:   query,
			Ctx:     context.Request.Context(),
			Session: *options.Meta.Session,
		})

		if err != nil {
			context.Error(err)
			context.Abort()

			return
		}

		context.JSON(options.Status, result)
	})

	return handlers
}
