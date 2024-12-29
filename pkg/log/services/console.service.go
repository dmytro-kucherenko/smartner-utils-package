package services

import (
	"fmt"

	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/log/types"
	"github.com/sirupsen/logrus"
)

type console struct {
	logger *logrus.Logger
	entry  *logrus.Entry
	name   string
}

func newEntry(name string, entry *logrus.Entry) types.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{FullTimestamp: true, DisableSorting: false})
	logger.SetOutput(logrus.StandardLogger().Out)

	return &console{logger: logger, name: name, entry: entry}
}

func NewConsole(name string) types.Logger {
	return newEntry(name, nil)
}

func (service *console) concat(messages []any) []any {
	nameFormatted := fmt.Sprintf("[%v] ", service.name)
	info := []interface{}{nameFormatted}

	return append(info, messages...)
}

func (service *console) actions() types.Actions {
	if service.entry != nil {
		return service.entry
	}

	return service.logger
}

func (service *console) Info(messages ...any) {
	service.actions().Info(service.concat(messages)...)
}

func (service *console) Warn(messages ...any) {
	service.actions().Warn(service.concat(messages)...)
}

func (service *console) Error(messages ...any) {
	service.actions().Error(service.concat(messages)...)
}

func (service *console) Fatal(messages ...any) {
	service.actions().Fatal(service.concat(messages)...)
}

func (service *console) Debug(messages ...any) {
	service.actions().Debug(service.concat(messages)...)
}

func (service *console) CreateEntry(fields map[string]any) types.Logger {
	entry := service.logger.WithFields(fields)

	return newEntry(service.name, entry)
}
