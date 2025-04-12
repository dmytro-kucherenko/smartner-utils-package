package adapter

type MethodConfig struct {
	Public bool
}

type CallerConfig map[string]MethodConfig

type Caller interface {
	Config() CallerConfig
}
