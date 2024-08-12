package logger

import (
	"go.uber.org/zap"
)

// Logger holds the application's logger instance.
type Logger struct {
	*zap.Logger
}

// Init initializes and returns a new Logger instance with zap.
func Init(development bool) (*Logger, error) {
	var logger *zap.Logger
	var err error

	if development {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}

	if err != nil {
		return nil, err
	}

	return &Logger{
		Logger: logger,
	}, nil
}

// Sync flushes any buffered log entries.
func (l *Logger) Sync() {
	l.Logger.Sync()
}
