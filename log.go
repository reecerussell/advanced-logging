package log

import (
	"fmt"

	"github.com/reecerussell/advanced-logging/config"
	"github.com/reecerussell/advanced-logging/logger"

	// Needed to register plugins
	_ "github.com/reecerussell/advanced-logging/plugin/file"
	_ "github.com/reecerussell/advanced-logging/plugin/log"
)

var (
	// DefaultLogger is the default output used to
	// write log data. This is initialised using
	// configuration found in a config file.
	DefaultLogger *logger.Logger
)

func init() {
	// Read config.
	c, err := config.ReadFromFile()
	if err != nil {
		panic(err)
	}

	// Init default logger.
	l, err := logger.NewLogger(c)
	if err != nil {
		panic(err)
	}

	DefaultLogger = l
}

// returns the default logger.
func mustUseDefaultLogger() *logger.Logger {
	if DefaultLogger == nil {
		panic(fmt.Errorf("cannot use default logger is it has not be assigned"))
	}

	return DefaultLogger
}

// Print methods

// Print calls Print() on the default logger. Arguments
// are handled in the manner of fmt.Print.
func Print(v ...interface{}) {
	mustUseDefaultLogger().Print(v...)
}

// Printf calls Printf() on the default logger. Arguments
// are handled in the manner of fmt.Printf.
func Printf(format string, v ...interface{}) {
	mustUseDefaultLogger().Printf(format, v...)
}

// Println calls Println() on the default logger. Arguments
// are handled in the manner of fmt.Println.
func Println(v ...interface{}) {
	mustUseDefaultLogger().Println(v...)
}

// Panic methods

// Panic is equivalent to Print() followed by a call to panic().
func Panic(v ...interface{}) {
	mustUseDefaultLogger().Panic(v...)
}

// Panicf is equivalent to Printf() followed by a call to panic().
func Panicf(format string, v ...interface{}) {
	mustUseDefaultLogger().Panicf(format, v...)
}

// Panicln is equivalent to Println() followed by a call to panic().
func Panicln(v ...interface{}) {
	mustUseDefaultLogger().Panicln(v...)
}

// Fatal methods

// Fatal is equivalent to Print() followed by a call to os.Exit(1).
func Fatal(v ...interface{}) {
	mustUseDefaultLogger().Fatal(v...)
}

// Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
func Fatalf(format string, v ...interface{}) {
	mustUseDefaultLogger().Fatalf(format, v...)
}

// Fatalln is equivalent to Println() followed by a call to os.Exit(1).
func Fatalln(v ...interface{}) {
	mustUseDefaultLogger().Fatalln(v...)
}
