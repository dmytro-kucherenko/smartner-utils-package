package types

import (
	"github.com/google/uuid"
)

type ID = uuid.UUID

type IDBind struct {
	value ID
}

func (bind *IDBind) UnmarshalParam(param string) error {
	UUID, err := uuid.Parse(param)
	if err != nil {
		return err
	}

	bind.value = ID(UUID)

	return nil
}

func (bind IDBind) Value() ID {
	return bind.value
}
