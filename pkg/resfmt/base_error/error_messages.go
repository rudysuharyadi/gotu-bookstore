package base_error

var ErrorMessages map[InternalCode]string

func GetErrorMessage(code InternalCode) string {
	if ErrorMessages == nil {
		return "Error messages not set correctly"
	}

	msg, ok := ErrorMessages[code]
	if !ok {
		return "Unknown error code"
	}
	return msg
}
