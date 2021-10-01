package parser

import (
	"io/fs"
	"path/filepath"
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	root := "./yaml/"
	
	filepath.Walk(root, func (path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		data, err := os.ReadFile(path)
		if err != nil {
			t.Errorf("failed to read file %v", path)
		}
		
		_, err = Parse(data)
		if err != nil {
			t.Fatalf("error reading test file %v:\n%v", path, err)
		}
		return nil
	})
}
