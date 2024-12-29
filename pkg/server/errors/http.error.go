package errors

import "errors"

type HttpError struct {
	error
	status  int
	details []string
}

func NewHttpError(status int, message string, details ...string) *HttpError {
	return &HttpError{errors.New(message), status, details}
}

func (err *HttpError) Status() int {
	return err.status
}

func (err *HttpError) Details() []string {
	return err.details
}
