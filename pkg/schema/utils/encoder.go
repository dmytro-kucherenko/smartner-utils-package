package utils

import (
	"encoding/json"

	"github.com/mitchellh/mapstructure"
)

func DecodeStruct[T any](data any, schema T) (err error) {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result:           schema,
		WeaklyTypedInput: true,
	})

	if err != nil {
		return
	}

	err = decoder.Decode(data)
	if err != nil {
		return
	}

	return
}

func EncodeStruct(data any) (result map[string]any, err error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return
	}

	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return
	}

	return result, nil
}
