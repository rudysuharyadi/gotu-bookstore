package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

// LogInstance is our shared log instance
var LogInstance *Log

// Log is our main object
type Log struct {
	logger *logrus.Logger
}

// Constructor
func NewLog() Log {
	LogInstance = &Log{
		logger: &logrus.Logger{
			Out:   os.Stdout,
			Hooks: make(logrus.LevelHooks),
			Formatter: &logrus.TextFormatter{
				FullTimestamp:          true,
				DisableLevelTruncation: true,
			},
			Level: logrus.InfoLevel,
		},
	}
	LogInstance.Info("logger instance initialized")
	return *LogInstance
}

func (l Log) UpdateLevel(logLevel string) error {
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return err
	}

	l.logger.Level = level
	l.Info("logger instance updated")
	return nil
}
