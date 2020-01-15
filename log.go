package log

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/reecerussell/advanced-logging/config"
	"github.com/reecerussell/advanced-logging/logger"

	// Needed to register plugins
	_ "github.com/reecerussell/advanced-logging/plugin/file"
	_ "github.com/reecerussell/advanced-logging/plugin/log"
)

const (
	defaultConfigFilePath  = "logger.json"
	configFilePathVariable = "LOGGER_CONFIG"
)

var (
	configuration *config.Configuration
	defaultLogger *logger.Logger
)

func init() {
	bytes, err := ioutil.ReadFile(
		getVariableOrDefault(configFilePathVariable, defaultConfigFilePath))
	if err != nil {
		panic(err)
	}

	var configuration config.Configuration
	_ = json.Unmarshal(bytes, &configuration)

	l, err := logger.NewLogger(&configuration)
	if err != nil {
		panic(err)
	}

	defaultLogger = l
}

// returns the default logger.
func mustUseDefaultLogger() *logger.Logger {
	if defaultLogger == nil {
		panic(fmt.Errorf("cannot use default logger is it has not be assigned"))
	}

	return defaultLogger
}

func Print(v ...interface{}) {
	mustUseDefaultLogger().Print(v...)
}

// returns the value of an environment variable. If the environment variable
// has no data, the defaultValue is returned.
func getVariableOrDefault(name, defaultValue string) string {
	if v := os.Getenv(name); v != "" {
		return v
	}

	return defaultValue
}
