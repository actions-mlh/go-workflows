package test

import (
	"io/fs"
	// "os"
	"path/filepath"
	"testing"
)

func TestParse(t *testing.T) {
	root := "./yaml/"

	filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		// data, err := os.ReadFile(path)
		if err != nil {
			t.Errorf("failed to read file %v", path)
		}

		// _, err = Parse(data)
		if err != nil {
			t.Errorf("error reading test file %v:\n%v", path, err)
		}
		return nil
	})
}
