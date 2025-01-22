package logs

import "go.uber.org/zap"

func NewLogger() *zap.Logger {
	config := zap.NewProductionConfig()
	// config.DisableCaller = true
	config.Level.SetLevel(zap.DebugLevel)
	log, err := config.Build()
	if err != nil {
		panic(err)
	}
	return log
}
