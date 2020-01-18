package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const (
	defaultConfigFilePath  = "logger.json"
	configFilePathVariable = "LOGGER_CONFIG"
)

// Configuration stores settings that can be used globally.
type Configuration struct {
	PrintPrefix string    `json:"printPrefix"`
	FatalPrefix string    `json:"fatalPrefix"`
	PanicPrefix string    `json:"panicPrefix"`
	Plugins     []*Plugin `json:"plugins"`
}

// ReadFromFile reads logger configuration from a JSON file.
// An error is returned if the files cannot be read, otherwise nil.
func ReadFromFile() (*Configuration, error) {
	var (
		bytes []byte
		err   error
	)

	// Read data from config file.
	if filepath := os.Getenv(configFilePathVariable); filepath != "" {
		// Use filepath stored in environment variable.
		bytes, err = ioutil.ReadFile(filepath)
	} else {
		// Use default filepath.
		bytes, err = ioutil.ReadFile(defaultConfigFilePath)
	}

	if err != nil {
		return nil, err
	}

	var c Configuration

	// Unmarshal the file data.
	_ = json.Unmarshal(bytes, &c)

	return &c, nil
}

// Plugin stores settings for each plugin.
type Plugin struct {
	Name    string                 `json:"name"`
	Enabled bool                   `json:"enabled"`
	Config  map[string]interface{} `json:"config"`
}
