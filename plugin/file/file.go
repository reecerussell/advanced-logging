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

// File is a plugin which writes logs to a physical file.
type File struct {
	mu sync.Mutex
}

// Output is the method used by the logger to instansiate the
// plugin. A io.Writer is returned, which is used to write to the file.
func (f *File) Output(opts map[string]interface{}) io.Writer {
	f.mu.Lock()
	defer f.mu.Unlock()

	// Read output directory from config.
	dir := opts["outputDir"].(string)

	// Ensure output directory exists.
	err := createDirectory(dir)
	if err != nil {
		panic(err)
	}

	// Get the file.
	file, err := writeFile(dir)
	if err != nil {
		panic(err)
	}

	return file
}

// writeFile returns a file which can be written to. This is
// either a newly created file, or an existing file that
// has been opened.
func writeFile(dir string) (*os.File, error) {
	for i := 0; true; i++ {
		now := time.Now()

		filename := fmt.Sprintf("log_%d%d%d.txt", now.Year(), now.Month(), now.Day())
		filename = path.Join(dir, filename)

		if i > 0 {
			filename = fmt.Sprintf("%s_%d", filename, i)
		}

		file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			return nil, err
		}

		stats, err := file.Stat()
		if err != nil {
			return nil, err
		}

		if stats.Size() <= 10^5 {
			return file, nil
		}
	}

	return nil, nil
}

// createDirectory ensures the output directory exists, and
// if not, creates it.
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
