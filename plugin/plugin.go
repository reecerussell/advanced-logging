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

type Plugin interface {
	Output(opts map[string]interface{}) io.Writer
}

func Get(name string) (Plugin, error) {
	l.RLock()
	defer l.RUnlock()

	plugin, ok := m[name]
	if !ok {
		return nil, fmt.Errorf("plugin '%s' not found", name)
	}

	return plugin, nil
}

func Register(name string, plugin Plugin) {
	l.Lock()
	defer l.Unlock()

	m[name] = plugin
}
