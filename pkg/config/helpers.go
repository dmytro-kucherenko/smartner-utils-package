package config

import (
	"os"
	"strconv"
)

func GetEnvString(key string) string {
	return os.Getenv(key)
}

func GetEnvInt(key string) int {
	value, _ := strconv.Atoi(GetEnvString(key))

	return value
}
