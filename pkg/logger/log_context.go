package logger

type LogContext struct {
	fields map[string]interface{}
}

func NewLogContext() LogContext {
	return LogContext{
		fields: make(map[string]interface{}),
	}
}

func (l LogContext) WithFields(fields map[string]interface{}) LogContext {
	l.fields = fields
	return l
}

func (l LogContext) WithField(key string, value interface{}) LogContext {
	l.fields[key] = value
	return l
}

func (l LogContext) WithError(err error) LogContext {
	l.fields["error"] = err
	return l
}
