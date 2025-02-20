package types

import (
	"github.com/google/uuid"
)

type ID = uuid.UUID

type IDBind string

func (bind *IDBind) Parse() (ID, error) {
	return uuid.Parse(string(*bind))
}
