package adapter

import (
	"net/http"

	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/schema/modules/common"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server/errors"
	"github.com/gin-gonic/gin"
)

const DefaultTimeZone = "UTC"

func handle[R any, P any](
	request server.Request[R, P],
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

		header, err := common.DecodeStruct[server.RequestHeader](context.Request.Header)
		if err != nil {
			abortValidationError(err)

			return
		}

		timeZone := DefaultTimeZone
		if len(header.TimeZone) > 0 {
			timeZone = header.TimeZone[0]
		}

		var params P
		paramsSchema := common.ModifySchema(params)
		if validateBody {
			if err := context.ShouldBind(paramsSchema); err != nil {
				abortValidationError(err)

				return
			}
		}

		if err := context.ShouldBindUri(paramsSchema); err != nil {
			abortValidationError(err)

			return
		}

		if err := context.ShouldBindQuery(paramsSchema); err != nil {
			abortValidationError(err)

			return
		}

		params = common.ParseSchema[P](paramsSchema)
		if err := options.Meta.Validator.Make(params); err != nil {
			abortValidationError(err)

			return
		}

		result, err := request(&server.RequestOptions[P]{
			Params:   params,
			Ctx:      context.Request.Context(),
			Session:  *options.Meta.Session,
			TimeZone: timeZone,
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
