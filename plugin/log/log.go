package log

import (
	"io"
	"os"

	gl "log"

	"github.com/reecerussell/advanced-logging/plugin"
)

func init() {
	plugin.Register("log", &Log{
		internal: gl.New(os.Stderr, "hello", gl.LstdFlags),
	})
}

type Log struct {
	internal *gl.Logger
}

func (l *Log) Output(opts map[string]interface{}) io.Writer {
	prefix, ok := opts["prefix"]
	if ok {
		l.internal.SetPrefix(prefix.(string))
	}

	return l.internal.Writer()
}
