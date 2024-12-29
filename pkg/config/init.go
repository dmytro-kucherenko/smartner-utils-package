package config

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func Init[T any](path string, getSchema func() T) T {
	godotenv.Load(path)
	validate := validator.New()
	schema := getSchema()

	if err := validate.Struct(schema); err != nil {
		panic(fmt.Sprintln("Enviromental variables error:", err.Error()))
	}

	return schema
}
