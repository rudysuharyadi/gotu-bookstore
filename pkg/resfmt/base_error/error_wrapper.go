package base_error

import (
	"github.com/pkg/errors"
)

// New returns a error with supplied message
func New(msg string) error {
	return errors.New(msg)
}

// Wrap returns new error by annotating the passed error with message
func Wrap(err error, msg string) error {
	return errors.Wrap(err, msg)
}

// Wrapf returns a new error by annotating passed error with formated message
func Wrapf(err error, msg string, args ...interface{}) error {
	return errors.Wrapf(err, msg, args...)
}

func Is(err error, targetError error) bool {
	return errors.Is(err, targetError)
}
