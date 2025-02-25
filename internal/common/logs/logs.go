package logs

import (
	"go.uber.org/zap"
)

type Logger struct {
	Logger *zap.Logger
}

func NewLogger() *Logger {
	config := zap.NewProductionConfig()
	config.DisableCaller = true
	config.Level.SetLevel(zap.DebugLevel)
	log, err := config.Build()
	if err != nil {
		panic(err)
	}

	return &Logger{log}
}

func (l *Logger) Info(msg string) {
	l.Logger.Info(msg)
}

func (l *Logger) Warn(msg string, err error) {
	l.Logger.Warn(msg, zap.Error(err))
}

func (l *Logger) Err(msg string, err error) {
	l.Logger.Error(msg, zap.Error(err))
}
