package logger

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/reecerussell/advanced-logging/config"
	"github.com/reecerussell/advanced-logging/plugin"
)

// These flags determine the mode in which a string is outputted in.
// For example, the output format could differ depending on the mode.
const (
	ModePrint = 1 << iota // indicates output is a print
	ModePanic             // indicates output is a panic
	ModeFatal             // indicates output is a fatal
)

// Logger is a high-level service for handling logging.
type Logger struct {
	printPrefix string
	panicPrefix string
	fatalPrefix string
	writers     []io.Writer
}

// NewLogger returns a new Logger, using the cofiguration c.
// An error will be returned if a plugin has not been registered.
func NewLogger(c *config.Configuration) (*Logger, error) {
	l := &Logger{
		printPrefix: c.PrintPrefix,
		panicPrefix: c.PanicPrefix,
		fatalPrefix: c.FatalPrefix,
	}

	for _, opts := range c.Plugins {
		// Skip plugin if disabled.
		if !opts.Enabled {
			continue
		}

		p, err := plugin.Get(opts.Name)
		if err != nil {
			return nil, err
		}

		w := p.Output(opts.Config)
		l.writers = append(l.writers, w)
	}

	return l, nil
}

// Output is used internally to write a string to each of
// the logger's internal io.Writers. An error will be returned
// if one of the io.Writers fails.
func (l *Logger) Output(s string, mode int) error {
	buf := l.formatHeader(mode)
	buf = append(buf, s...)

	if (len(s) == 0) || s[len(s)-1] != '\n' {
		buf = append(buf, '\n')
	}

	for _, w := range l.writers {
		_, err := w.Write(buf)
		if err != nil {
			return err
		}
	}

	return nil
}

// formatHeader returns a header/prefix for the specified logging mode.
func (l *Logger) formatHeader(mode int) []byte {
	var (
		pre, head string
	)

	switch mode {
	case ModePrint:
		pre = l.printPrefix
		break
	case ModeFatal:
		pre = l.fatalPrefix
		break
	case ModePanic:
		pre = l.panicPrefix
		break
	}

	now := time.Now()

	head = fmt.Sprintf("[%d-%s-%d]", now.Day(), now.Month(), now.Year())
	head += fmt.Sprintf("[%d:%d:%d.%d]", now.Hour(), now.Minute(), now.Second(), now.Nanosecond()/1000)

	return []byte(fmt.Sprintf("[%s]%s ", pre, head))
}

// Log methods

// Print calls l.Output to print to the logger. Aguments
// are handled in the manner of fmt.Print.
func (l *Logger) Print(v ...interface{}) {
	l.Output(fmt.Sprint(v...), ModePrint)
}

// Printf calls l.Output to print to the logger. Arguments
// are handled in the form of fmt.Printf.
func (l *Logger) Printf(format string, v ...interface{}) {
	l.Output(fmt.Sprintf(format, v...), ModePrint)
}

// Println calls l.Output to print to the logger. Arguments
// are handled in the manner of fmt.Println.
func (l *Logger) Println(v ...interface{}) {
	l.Output(fmt.Sprintln(v...), ModePrint)
}

// Fatal methods

// Fatal is equivalent to l.Print() followed by
// a call to os.Exit(1);
func (l *Logger) Fatal(v ...interface{}) {
	l.Output(fmt.Sprint(v...), ModeFatal)
	os.Exit(1)
}

// Fatalf is equivalent to l.Printf() followed by
// a call to os.Exit(1);
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.Output(fmt.Sprintf(format, v...), ModeFatal)
	os.Exit(1)
}

// Fatalln is equivalent to l.Println() followed by
// a call to os.Exit(1);
func (l *Logger) Fatalln(v ...interface{}) {
	l.Output(fmt.Sprintln(v...), ModeFatal)
	os.Exit(1)
}

// Panic methods

// Panic is equivalent to l.Print() follows by call to panic().
func (l *Logger) Panic(v ...interface{}) {
	s := fmt.Sprint(v...)
	l.Output(s, ModePanic)
	panic(s)
}

// Panicf is equivalent to l.Printf() follows by call to panic().
func (l *Logger) Panicf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	l.Output(s, ModePanic)
	panic(s)
}

// Panicln is equivalent to l.Println() follows by call to panic().
func (l *Logger) Panicln(v ...interface{}) {
	s := fmt.Sprintln(v...)
	l.Output(s, ModePanic)
	panic(s)
}
