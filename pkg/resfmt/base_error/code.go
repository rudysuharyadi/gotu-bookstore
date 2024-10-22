package base_error

type Code string

const (
	BadRequestError          Code = "Bad Request Error"
	RequestHeaderError       Code = "Request Header Error"
	ConflictError            Code = "Conflict Error"
	InternalError            Code = "Internal Error"
	ForbiddenError           Code = "Forbidden Error"
	RequestTimeoutError      Code = "Request Timeout Error"
	NotAcceptableError       Code = "Not Acceptable Error"
	NotFoundError            Code = "Not Found Error"
	TooManyRequestError      Code = "Too Many Request Error"
	UnauthorizedError        Code = "Unauthorized Error"
	UnprocessableEntityError Code = "Unprocessable Entity Error"
)
