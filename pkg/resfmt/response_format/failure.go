package response_format

import (
	"encoding/json"
	"fmt"
	"gotu-bookstore/pkg/resfmt/base_error"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

type Failure struct {
	StatusCode int    `json:"status_code"`
	Status     string `json:"status" binding:"required"`
	base_error.BaseError
}

func NewFailure(err error) *Failure {
	var baseError base_error.BaseError
	ok := errors.As(err, &baseError)
	if ok {
		return generateFailure(baseError)
	}

	var failure Failure
	ok = errors.As(err, &failure)
	if ok {
		return &failure
	}

	return generateFailure(base_error.NewBaseError(base_error.InternalError, base_error.DefaultInternalCode, []error{err}, err.Error()))
}

func generateFailure(baseError base_error.BaseError) *Failure {
	failureStatus := "failure"
	switch baseError.Code {
	case base_error.BadRequestError:
		return &Failure{StatusCode: http.StatusBadRequest, Status: failureStatus, BaseError: baseError}
	case base_error.UnprocessableEntityError:
		return &Failure{StatusCode: http.StatusUnprocessableEntity, Status: failureStatus, BaseError: baseError}
	case base_error.ForbiddenError:
		return &Failure{StatusCode: http.StatusForbidden, Status: failureStatus, BaseError: baseError}
	case base_error.UnauthorizedError:
		return &Failure{StatusCode: http.StatusUnauthorized, Status: failureStatus, BaseError: baseError}
	case base_error.ConflictError:
		return &Failure{StatusCode: http.StatusConflict, Status: failureStatus, BaseError: baseError}
	case base_error.NotAcceptableError:
		return &Failure{StatusCode: http.StatusNotAcceptable, Status: failureStatus, BaseError: baseError}
	case base_error.RequestTimeoutError:
		return &Failure{StatusCode: http.StatusRequestTimeout, Status: failureStatus, BaseError: baseError}
	case base_error.NotFoundError:
		return &Failure{StatusCode: http.StatusNotFound, Status: failureStatus, BaseError: baseError}
	case base_error.TooManyRequestError:
		return &Failure{StatusCode: http.StatusTooManyRequests, Status: failureStatus, BaseError: baseError}
	case base_error.RequestHeaderError:
		return &Failure{StatusCode: http.StatusBadRequest, Status: failureStatus, BaseError: baseError}
	case base_error.InternalError:
		return &Failure{StatusCode: http.StatusInternalServerError, Status: failureStatus, BaseError: baseError}
	default:
		return &Failure{StatusCode: http.StatusInternalServerError, Status: failureStatus, BaseError: baseError}
	}
}

func (f *Failure) MarshalJSON() ([]byte, error) {
	type Alias Failure
	aux := &struct {
		StatusCode int    `json:"status_code"`
		Status     string `json:"status" binding:"required"`
		*base_error.BaseError
		ErrorMessages map[string]interface{} `json:"errors,omitempty"`
	}{
		StatusCode: f.StatusCode,
		Status:     f.Status,
		BaseError:  &f.BaseError,
	}

	if aux.BaseError.Causes != nil && len(aux.BaseError.Causes) > 0 {
		aux.ErrorMessages = make(map[string]interface{})
		for i, v := range aux.BaseError.Causes {
			var subError base_error.SubError
			ok := errors.As(v, &subError)
			if ok {
				aux.ErrorMessages[subError.Label] = subError.Message
			} else {
				key := fmt.Sprintf("%s_%d", strings.ToLower(string(base_error.DefaultInternalCode)), i)
				aux.ErrorMessages[key] = v.Error()
			}
		}
	}

	return json.Marshal(aux)
}
