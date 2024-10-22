package base_error

func NewInternalError(internalCode InternalCode, msg ...string) BaseError {
	msg = append([]string{GetErrorMessage(internalCode)}, msg...)
	return NewBaseError(InternalError, internalCode, nil, msg...)
}

func NewInternalErrorWithCause(cause error, internalCode InternalCode, msg ...string) BaseError {
	msg = append([]string{GetErrorMessage(internalCode)}, msg...)
	return NewBaseError(InternalError, internalCode, []error{cause}, msg...)
}

func NewInternalErrorWithCauses(causes []error, internalCode InternalCode, msg ...string) BaseError {
	msg = append([]string{GetErrorMessage(internalCode)}, msg...)
	return NewBaseError(InternalError, internalCode, causes, msg...)
}

func NewBadRequestError(internalCode InternalCode, msg ...string) BaseError {
	msg = append([]string{GetErrorMessage(internalCode)}, msg...)
	return NewBaseError(BadRequestError, internalCode, nil, msg...)
}

func NewBadRequestErrorWithCause(cause error, internalCode InternalCode, msg ...string) BaseError {
	msg = append([]string{GetErrorMessage(internalCode)}, msg...)
	return NewBaseError(BadRequestError, internalCode, []error{cause}, msg...)
}

func NewBadRequestErrorWithCauses(causes []error, internalCode InternalCode, msg ...string) BaseError {
	msg = append([]string{GetErrorMessage(internalCode)}, msg...)
	return NewBaseError(BadRequestError, internalCode, causes, msg...)
}

func NewRequestHeaderError(internalCode InternalCode, msg ...string) BaseError {
	msg = append([]string{GetErrorMessage(internalCode)}, msg...)
	return NewBaseError(RequestHeaderError, internalCode, nil, msg...)
}

func NewRequestHeaderErrorWithCause(cause error, internalCode InternalCode, msg ...string) BaseError {
	msg = append([]string{GetErrorMessage(internalCode)}, msg...)
	return NewBaseError(RequestHeaderError, internalCode, []error{cause}, msg...)
}

func NewRequestHeaderErrorWithCauses(causes []error, internalCode InternalCode, msg ...string) BaseError {
	msg = append([]string{GetErrorMessage(internalCode)}, msg...)
	return NewBaseError(RequestHeaderError, internalCode, causes, msg...)
}

func NewConflictError(internalCode InternalCode, msg ...string) BaseError {
	msg = append([]string{GetErrorMessage(internalCode)}, msg...)
	return NewBaseError(ConflictError, internalCode, nil, msg...)
}

func NewConflictErrorWithCause(cause error, internalCode InternalCode, msg ...string) BaseError {
	msg = append([]string{GetErrorMessage(internalCode)}, msg...)
	return NewBaseError(ConflictError, internalCode, []error{cause}, msg...)
}

func NewConflictErrorWithCauses(causes []error, internalCode InternalCode, msg ...string) BaseError {
	msg = append([]string{GetErrorMessage(internalCode)}, msg...)
	return NewBaseError(ConflictError, internalCode, causes, msg...)
}

func NewForbiddenError(internalCode InternalCode, msg ...string) BaseError {
	msg = append([]string{GetErrorMessage(internalCode)}, msg...)
	return NewBaseError(ForbiddenError, internalCode, nil, msg...)
}

func NewForbiddenErrorWithCause(cause error, internalCode InternalCode, msg ...string) BaseError {
	msg = append([]string{GetErrorMessage(internalCode)}, msg...)
	return NewBaseError(ForbiddenError, internalCode, []error{cause}, msg...)
}

func NewForbiddenErrorWithCauses(causes []error, internalCode InternalCode, msg ...string) BaseError {
	msg = append([]string{GetErrorMessage(internalCode)}, msg...)
	return NewBaseError(ForbiddenError, internalCode, causes, msg...)
}

func NewRequestTimeoutError(internalCode InternalCode, msg ...string) BaseError {
	msg = append([]string{GetErrorMessage(internalCode)}, msg...)
	return NewBaseError(RequestTimeoutError, internalCode, nil, msg...)
}

func NewRequestTimeoutErrorWithCause(cause error, internalCode InternalCode, msg ...string) BaseError {
	msg = append([]string{GetErrorMessage(internalCode)}, msg...)
	return NewBaseError(RequestTimeoutError, internalCode, []error{cause}, msg...)
}

func NewRequestTimeoutErrorWithCauses(causes []error, internalCode InternalCode, msg ...string) BaseError {
	msg = append([]string{GetErrorMessage(internalCode)}, msg...)
	return NewBaseError(RequestTimeoutError, internalCode, causes, msg...)
}

func NewNotAcceptableError(internalCode InternalCode, msg ...string) BaseError {
	msg = append([]string{GetErrorMessage(internalCode)}, msg...)
	return NewBaseError(NotAcceptableError, internalCode, nil, msg...)
}

func NewNotAcceptableErrorWithCause(cause error, internalCode InternalCode, msg ...string) BaseError {
	msg = append([]string{GetErrorMessage(internalCode)}, msg...)
	return NewBaseError(NotAcceptableError, internalCode, []error{cause}, msg...)
}

func NewNotAcceptableErrorWithCauses(causes []error, internalCode InternalCode, msg ...string) BaseError {
	msg = append([]string{GetErrorMessage(internalCode)}, msg...)
	return NewBaseError(NotAcceptableError, internalCode, causes, msg...)
}

func NewNotFoundError(internalCode InternalCode, msg ...string) BaseError {
	msg = append([]string{GetErrorMessage(internalCode)}, msg...)
	return NewBaseError(NotFoundError, internalCode, nil, msg...)
}

func NewNotFoundErrorWithCause(cause error, internalCode InternalCode, msg ...string) BaseError {
	msg = append([]string{GetErrorMessage(internalCode)}, msg...)
	return NewBaseError(NotFoundError, internalCode, []error{cause}, msg...)
}

func NewNotFoundErrorWithCauses(causes []error, internalCode InternalCode, msg ...string) BaseError {
	msg = append([]string{GetErrorMessage(internalCode)}, msg...)
	return NewBaseError(NotFoundError, internalCode, causes, msg...)
}

func NewTooManyRequestError(internalCode InternalCode, msg ...string) BaseError {
	msg = append([]string{GetErrorMessage(internalCode)}, msg...)
	return NewBaseError(TooManyRequestError, internalCode, nil, msg...)
}

func NewTooManyRequestErrorWithCause(cause error, internalCode InternalCode, msg ...string) BaseError {
	msg = append([]string{GetErrorMessage(internalCode)}, msg...)
	return NewBaseError(TooManyRequestError, internalCode, []error{cause}, msg...)
}

func NewTooManyRequestErrorWithCauses(causes []error, internalCode InternalCode, msg ...string) BaseError {
	msg = append([]string{GetErrorMessage(internalCode)}, msg...)
	return NewBaseError(TooManyRequestError, internalCode, causes, msg...)
}

func NewUnauthorizedError(internalCode InternalCode, msg ...string) BaseError {
	msg = append([]string{GetErrorMessage(internalCode)}, msg...)
	return NewBaseError(UnauthorizedError, internalCode, nil, msg...)
}

func NewUnauthorizedErrorWithCause(cause error, internalCode InternalCode, msg ...string) BaseError {
	msg = append([]string{GetErrorMessage(internalCode)}, msg...)
	return NewBaseError(UnauthorizedError, internalCode, []error{cause}, msg...)
}

func NewUnauthorizedErrorWithCauses(causes []error, internalCode InternalCode, msg ...string) BaseError {
	msg = append([]string{GetErrorMessage(internalCode)}, msg...)
	return NewBaseError(UnauthorizedError, internalCode, causes, msg...)
}

func NewUnprocessableEntityError(internalCode InternalCode, msg ...string) BaseError {
	msg = append([]string{GetErrorMessage(internalCode)}, msg...)
	return NewBaseError(UnprocessableEntityError, internalCode, nil, msg...)
}

func NewUnprocessableEntityErrorWithCause(cause error, internalCode InternalCode, msg ...string) BaseError {
	msg = append([]string{GetErrorMessage(internalCode)}, msg...)
	return NewBaseError(UnprocessableEntityError, internalCode, []error{cause}, msg...)
}

func NewUnprocessableEntityErrorWithCauses(causes []error, internalCode InternalCode, msg ...string) BaseError {
	msg = append([]string{GetErrorMessage(internalCode)}, msg...)
	return NewBaseError(UnprocessableEntityError, internalCode, causes, msg...)
}
