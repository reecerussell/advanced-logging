package plugin

import (
	"fmt"
	"io"
	"sync"
)

var (
	l = sync.RWMutex{}
	m = map[string]Plugin{}
)

// Plugin is a high-level interface used by the logger to
// output data through a plugin.
type Plugin interface {
	Output(opts map[string]interface{}) io.Writer
}

// Get returns a plugin from memory.
func Get(name string) (Plugin, error) {
	l.RLock()
	defer l.RUnlock()

	plugin, ok := m[name]
	if !ok {
		return nil, fmt.Errorf("plugin '%s' not found", name)
	}

	return plugin, nil
}

// Register adds a plugin and stores it in memory.
func Register(name string, plugin Plugin) {
	l.Lock()
	defer l.Unlock()

	m[name] = plugin
}
