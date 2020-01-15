package log

import (
	"io"
	"os"

	gl "log"

	"github.com/reecerussell/advanced-logging/plugin"
)

func init() {
	plugin.Register("log", &Log{
		internal: gl.New(os.Stderr, "", gl.LstdFlags),
	})
}

// Log is a plugin for the standard go log package.
type Log struct {
	internal *gl.Logger
}

// Output writes to the Stderr output, with the default log package.
func (l *Log) Output(opts map[string]interface{}) io.Writer {
	prefix, ok := opts["prefix"]
	if ok {
		l.internal.SetPrefix(prefix.(string))
	}

	return l.internal.Writer()
}
