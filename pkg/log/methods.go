package log

import (
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/log/services"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/log/types"
)

func New(name string) types.Logger {
	return services.NewConsole(name)
}
