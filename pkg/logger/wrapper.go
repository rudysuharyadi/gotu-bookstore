package logger

import "github.com/sirupsen/logrus"

func (l Log) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

func (l Log) ErrorfWithContext(context LogContext, format string, args ...interface{}) {
	l.logger.WithFields(l.convertToLogrusFields(context.fields)).Errorf(format, args...)
}

func (l Log) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l Log) ErrorWithContext(context LogContext, args ...interface{}) {
	l.logger.WithFields(l.convertToLogrusFields(context.fields)).Error(args...)
}

func (l Log) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l Log) InfofWithContext(context LogContext, format string, args ...interface{}) {
	l.logger.WithFields(l.convertToLogrusFields(context.fields)).Infof(format, args...)
}

func (l Log) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l Log) InfoWithContext(context LogContext, args ...interface{}) {
	l.logger.WithFields(l.convertToLogrusFields(context.fields)).Info(args...)
}

func (l Log) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

func (l Log) DebugfWithContext(context LogContext, format string, args ...interface{}) {
	l.logger.WithFields(l.convertToLogrusFields(context.fields)).Debugf(format, args...)
}

func (l Log) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l Log) DebugWithContext(context LogContext, args ...interface{}) {
	l.logger.WithFields(l.convertToLogrusFields(context.fields)).Debug(args...)
}

func (l Log) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

func (l Log) WarnfWithContext(context LogContext, format string, args ...interface{}) {
	l.logger.WithFields(l.convertToLogrusFields(context.fields)).Warnf(format, args...)
}

func (l Log) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l Log) WarnWithContext(context LogContext, args ...interface{}) {
	l.logger.WithFields(l.convertToLogrusFields(context.fields)).Warn(args...)
}

func (l Log) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

func (l Log) FatalfWithContext(context LogContext, format string, args ...interface{}) {
	l.logger.WithFields(l.convertToLogrusFields(context.fields)).Fatalf(format, args...)
}

func (l Log) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l Log) FatalWithContext(context LogContext, args ...interface{}) {
	l.logger.WithFields(l.convertToLogrusFields(context.fields)).Fatal(args...)
}

func (l Log) Panicf(format string, args ...interface{}) {
	l.logger.Panicf(format, args...)
}

func (l Log) PanicfWithContext(context LogContext, format string, args ...interface{}) {
	l.logger.WithFields(l.convertToLogrusFields(context.fields)).Panicf(format, args...)
}

func (l Log) Panic(args ...interface{}) {
	l.logger.Panic(args...)
}

func (l Log) PanicWithContext(context LogContext, args ...interface{}) {
	l.logger.WithFields(l.convertToLogrusFields(context.fields)).Panic(args...)
}

func (l Log) convertToLogrusFields(fields map[string]interface{}) logrus.Fields {
	logrusFields := logrus.Fields{}
	for index, val := range fields {
		logrusFields[index] = val
	}
	return logrusFields
}
