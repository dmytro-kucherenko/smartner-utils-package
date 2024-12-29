package types

type Actions interface {
	Info(messages ...any)
	Warn(messages ...any)
	Error(messages ...any)
	Fatal(messages ...any)
	Debug(messages ...any)
}

type Logger interface {
	Actions
	CreateEntry(fields map[string]any) Logger
}
