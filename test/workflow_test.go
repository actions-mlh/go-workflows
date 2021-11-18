package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	"c2c-actions-mlh-workflow-parser/lint"
)

/*
   Run the Spew() function on each file in yaml/clean.
   We expect to see each line in the original file appear in the spewed contents.
   If not, throw an error.
*/
func TestWorkflow(t *testing.T) {
	root := "../yaml/clean/"
	t.Log("WORKFLOW TESTS:")
	filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() ||
			!(filepath.Ext(path) == ".yml" ||
				filepath.Ext(path) == ".yaml") {
			return nil
		}
		input, err := os.ReadFile(path)
		if err != nil {
			t.Errorf("error reading file %s: %s", path, err)
		}
		lint.Spew(input)
		return nil
	})
}