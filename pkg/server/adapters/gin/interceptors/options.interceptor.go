package interceptors

import (
	"net/http"

	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/meta"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server/errors"
	"github.com/gin-gonic/gin"
)

func onValidOptions(c *gin.Context, options meta.Options) {
	if options.TimeZone == "" {
		options.TimeZone = "UTC"
	}

	ctx := meta.SetOptionsContext(c.Request.Context(), options)
	c.Request = c.Request.WithContext(ctx)
}

func onErrorOptions(c *gin.Context, err error) {
	c.Error(errors.NewHttpError(http.StatusBadRequest, "invalid metadata", err.Error()))
	c.Abort()
}

func Options() gin.HandlerFunc {
	return Header(onValidOptions, onErrorOptions)
}
