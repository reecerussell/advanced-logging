package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/reecerussell/advanced-logging/config"

	"github.com/reecerussell/advanced-logging/logger"

	// Needed to register plugins
	_ "github.com/reecerussell/advanced-logging/plugin/file"
	_ "github.com/reecerussell/advanced-logging/plugin/log"
)

func main() {
	f := "Hello %s!"
	v := "World"
	s := fmt.Sprintf(f, v)

	bytes, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	var c config.Configuration
	_ = json.Unmarshal(bytes, &c)

	logger, err := logger.NewLogger(&c)
	if err != nil {
		panic(err)
	}

	logger.Print(s)
}
