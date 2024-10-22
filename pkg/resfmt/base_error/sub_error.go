package base_error

import (
	"strings"
)

type SubError struct {
	Message string `json:"message" binding:"required" `
	Label   string `json:"-"`
}

func NewSubError(label string, msg ...string) SubError {
	newError := SubError{
		Message: "Unknown Error",
		Label:   label,
	}
	if len(msg) > 0 {
		newError.Message = strings.TrimSpace(strings.Join(msg, ", "))
	}
	return newError
}

func (err SubError) Error() string {
	return err.Message
}
