package meta

import "github.com/dmytro-kucherenko/smartner-utils-package/pkg/types"

type key string

const (
	keySession  key = "session"
	keyUserID   key = "userId"
	keyTimeZone key = "timeZone"
)

type Options struct {
	TimeZone string `json:"timeZone" mapstructure:"timeZone" validate:"omitempty,timezone"`
}

type OptionsParams struct {
	TimeZone string `json:"timeZone" mapstructure:"timeZone" validate:"omitempty,timezone"`
}

type optionsMetadata struct {
	TimeZone []string `json:"timeZone" mapstructure:"timeZone" validate:"required,min=1,max=1,dive,timezone"`
}

type Session struct {
	UserID types.ID `json:"userId" mapstructure:"userId" validate:"required,uuid4"`
}

type sessionMetadata struct {
	UserID []string `json:"userId" mapstructure:"userId" validate:"required,min=1,max=1,dive,uuid4"`
}
