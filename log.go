package log

import (
	"fmt"

	"github.com/reecerussell/advanced-logging/config"
	"github.com/reecerussell/advanced-logging/logger"
)

var (
	configuration *config.Configuration
	defaultLogger *logger.Logger
)

func init() {
	configuration = new(config.Configuration)
	l, err := logger.NewLogger(configuration)
	if err != nil {
		panic(err)
	}

	defaultLogger = l
}

func mustUseDefaultLogger() *logger.Logger {
	if defaultLogger == nil {
		panic(fmt.Errorf("cannot use default logger is it has not be assigned"))
	}

	return defaultLogger
}

func Print(v ...interface{}) {
	mustUseDefaultLogger().Print(v...)
}
