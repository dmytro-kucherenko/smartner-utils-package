package adapter

import (
	"net/http"
	"reflect"

	adapter "github.com/dmytro-kucherenko/smartner-utils-package/pkg/schema/adapters/playground"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/schema/modules/common"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server/errors"
	"github.com/gin-gonic/gin"
)

func handle[R any, P any](
	request server.Request[R, P],
	options *RequestConfig,
	validateBody bool,
) []gin.HandlerFunc {
	handlers := make([]gin.HandlerFunc, 0, len(options.Interceptors)+2)
	for _, middleware := range options.Interceptors {
		handlers = append(handlers, gin.HandlerFunc(middleware))
	}

	handlers = append(handlers, func(c *gin.Context) {
		abortValidationError := func(err error) {
			c.Error(errors.NewHttpError(http.StatusBadRequest, "validation error", err.Error()))
			c.Abort()
		}

		var params P
		if paramsType := reflect.TypeOf(params); paramsType != nil {
			paramsSchema := common.ModifySchema(params)
			if validateBody {
				if err := c.ShouldBind(paramsSchema); err != nil {
					abortValidationError(err)

					return
				}
			}

			if err := c.ShouldBindUri(paramsSchema); err != nil {
				abortValidationError(err)

				return
			}

			if err := c.ShouldBindQuery(paramsSchema); err != nil {
				abortValidationError(err)

				return
			}

			params = common.ParseSchema[P](paramsSchema)
			if err := adapter.ValidateStruct(params); err != nil {
				abortValidationError(err)

				return
			}
		}

		result, err := request(c.Request.Context(), params)

		if err != nil {
			c.Error(err)
			c.Abort()

			return
		}

		c.JSON(options.Status, result)
	})

	return handlers
}
