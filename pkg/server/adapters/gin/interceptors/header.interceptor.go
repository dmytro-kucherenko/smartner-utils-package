package interceptors

import (
	adapter "github.com/dmytro-kucherenko/smartner-utils-package/pkg/schema/adapters/playground"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/schema/modules/common"
	"github.com/gin-gonic/gin"
)

func Header[T any](onValid func(c *gin.Context, data T), onError func(c *gin.Context, err error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data T
		schema := common.ModifySchema(data)
		err := c.ShouldBindHeader(schema)

		if err != nil {
			onError(c, err)

			return
		}

		data = common.ParseSchema[T](schema)

		err = adapter.ValidateStruct(&data)
		if err != nil {
			onError(c, err)

			return
		}

		onValid(c, data)
		c.Next()
	}
}
