package api

import (
	"errors"
)

type GPT interface {
	Send(string) (string, error)
}

var ErrNotSupportModel = errors.New("not support model")
