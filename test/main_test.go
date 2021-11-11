package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"c2c-actions-mlh-workflow-parser/lint"
)

func TestParse(t *testing.T) {
	root := "../yaml/clean/"

	t.Log("CLEAN TESTS:")
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
		problems, err := lint.Lint(path, input)
		if err != nil {
			t.Errorf("error linting file %s:\n%s", path, err)
		}
		if len(problems) > 0 {
			t.Errorf("error(s) found in clean file %s:\n%s", path, strings.Join(problems, "\n"))
		}
		return nil
	})

	t.Log("DIRTY TESTS:")
	root = "../yaml/dirty/"
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
		expectedBytes, err := os.ReadFile(path + ".exp")
		if err != nil {
			t.Errorf("error reading file %s.exp: %s", path, err)
		}
		problems, err := lint.Lint(path, input)
		if err != nil {
			t.Errorf("error linting file %s: %s", path, err)
		}
		expected := strings.Split(string(expectedBytes), "\n")
		for _, expProblem := range expected {
			if expProblem == "" {
				continue
			}
			if !contains(problems, expProblem) {
				t.Errorf("missing EXPECTED problem:\n%s", expProblem)
			}
		}

		for _, problem := range problems {
			if !contains(expected, problem) {
				t.Errorf("found UNEXPECTED problem:\n%s", problem)
			}
		}
		return nil
	})
}

// yes, i know i'm in O(n^2), but if you have more than 100 errors in one YAML file you should rethink your life
func contains(slice []string, item string) bool {
	for _, val := range slice {
		if val == item {
			return true
		}
	}
	return false
}
