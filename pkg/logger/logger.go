package logger

import (
	"go.uber.org/zap"
)

// Logger holds the application's logger instances.
type Logger struct {
	InfoLogger  *zap.Logger
	ErrorLogger *zap.Logger
}

// Init initializes and returns a new Logger instance with zap.
func Init() (*Logger, error) {
	// Set up development and production configurations for logging
	infoLogger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}

	errorLogger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	return &Logger{
		InfoLogger:  infoLogger,
		ErrorLogger: errorLogger,
	}, nil
}

// Sync flushes any buffered log entries.
func (l *Logger) Sync() {
	l.InfoLogger.Sync()
	l.ErrorLogger.Sync()
}
