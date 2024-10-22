package base_error

import (
	"fmt"
	"time"
)

// BaseError holds code and description of an error
type BaseError struct {
	Code         Code         `json:"code" binding:"required"`
	Timestamp    string       `json:"timestamp" binding:"required"`
	InternalCode InternalCode `json:"internal_code,omitempty"`
	Causes       []error      `json:"-"`
	SubError
}

func NewBaseError(code Code, internalCode InternalCode, causes []error, msg ...string) BaseError {
	newError := BaseError{
		Code:         code,
		Timestamp:    time.Now().Format(time.RFC3339),
		InternalCode: internalCode,
		Causes:       causes,
		SubError:     NewSubError("Error", msg...),
	}
	return newError
}

// Implement this two interface so that BaseError will identified as error.
func (err BaseError) Error() string {
	return fmt.Sprintf("[%s] %s\n", string(err.InternalCode), err.Message)
}
