package interceptors

import (
	"net/http"

	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/meta"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server/errors"
	"github.com/gin-gonic/gin"
)

func onValidAuth(c *gin.Context, session meta.Session) {
	ctx := meta.SetSessionContext(c.Request.Context(), session)
	c.Request = c.Request.WithContext(ctx)
}

func onErrorAuth(c *gin.Context, err error) {
	c.Error(errors.NewHttpError(http.StatusUnauthorized, "invalid session", err.Error()))
	c.Abort()
}

func Auth() gin.HandlerFunc {
	return Header(onValidAuth, onErrorAuth)
}
