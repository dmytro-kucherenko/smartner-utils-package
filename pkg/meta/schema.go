package meta

import "context"

func GetTimeZone(ctx context.Context) (string, bool) {
	value, ok := ctx.Value(keyTimeZone).(string)

	return value, ok
}

func GetSession(ctx context.Context) (Session, bool) {
	value, ok := ctx.Value(keySession).(Session)

	return value, ok
}
