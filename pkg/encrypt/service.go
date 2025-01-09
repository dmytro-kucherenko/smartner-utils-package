package encrypt

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"net/http"

	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server/errors"
)

type Service struct {
	secret string
	rounds uint8
}

func NewService(secret string, rounds uint8) *Service {
	return &Service{secret, rounds}
}

func (service *Service) genSalt() (salt string, err error) {
	sequence := make([]byte, service.rounds)
	_, err = rand.Read(sequence)
	if err != nil {
		err = errors.NewHttpError(http.StatusInternalServerError, MessageFailed)

		return
	}

	salt = hex.EncodeToString(sequence)

	return
}

func (service *Service) get(data string, salt string) Value {
	hasher := sha256.New()
	hasher.Write([]byte(data + salt + service.secret))

	hash := hex.EncodeToString(hasher.Sum(nil))

	return Value{Hash: hash, Salt: salt}
}

func (service *Service) Gen(data string) (value Value, err error) {
	salt, err := service.genSalt()
	if err != nil {
		return
	}

	value = service.get(data, salt)

	return
}

func (service *Service) Verify(data string, value Value) bool {
	computed := service.get(data, value.Salt)

	return subtle.ConstantTimeCompare([]byte(computed.Hash), []byte(value.Hash)) == 1
}
