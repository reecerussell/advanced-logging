package file

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"sync"
	"time"

	"github.com/reecerussell/advanced-logging/plugin"
)

func init() {
	plugin.Register("file", &File{
		mu: sync.Mutex{},
	})
}

type File struct {
	mu sync.Mutex
}

func (f *File) Output(opts map[string]interface{}) io.Writer {
	f.mu.Lock()
	defer f.mu.Unlock()

	dir := opts["outputDir"].(string)

	// Ensure output directory exists.
	err := createDirectory(dir)
	if err != nil {
		panic(err)
	}

	file, err := writeFile(dir)
	if err != nil {
		panic(err)
	}

	return file
}

func writeFile(dir string) (*os.File, error) {
	now := time.Now()
	filename := fmt.Sprintf("log_%d%d%d.txt", now.Year(), now.Month(), now.Day())
	filename = path.Join(dir, filename)

	return os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
}

func createDirectory(dir string) error {
	dir = strings.TrimSpace(dir)
	if dir == "" {
		return fmt.Errorf("file: output directory is empty")
	}

	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		return err
	}

	return nil
}
