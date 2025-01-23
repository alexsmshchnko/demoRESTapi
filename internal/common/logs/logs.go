package logs

import (
	"go.uber.org/zap"
)

type Logger struct {
	Logger *zap.Logger
}

func NewLogger() *Logger {
	config := zap.NewProductionConfig()
	// config.DisableCaller = true
	config.Level.SetLevel(zap.DebugLevel)
	log, err := config.Build()
	if err != nil {
		panic(err)
	}

	return &Logger{log}
}
