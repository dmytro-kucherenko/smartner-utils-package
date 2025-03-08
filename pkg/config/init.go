package config

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func Init[T any](path string, getSchema func() T) (T, error) {
	godotenv.Load(path)
	validate := validator.New()
	schema := getSchema()

	if err := validate.Struct(schema); err != nil {
		return schema, fmt.Errorf("enviromental variables error: %v", err.Error())
	}

	return schema, nil
}
