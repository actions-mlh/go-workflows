package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/actions-mlh/go-workflows/lint"
)

/*
   Scan the directory of clean files and attempt to lint each of them.
   The files come from https://github.com/actions/starter-workflows and should contain no errors.
   If any errors are found, fail the test and report the found errors.

   see yaml/clean/README.md for more information.
*/
func TestClean(t *testing.T) {
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
		// lint call
		problems, err := lint.Lint(path, input)
		if err != nil {
			t.Errorf("error linting file %s:\n%s", path, err)
		}
		if len(problems) > 0 {
			var arrayOfProblems []string
			for _, problem := range problems {
				arrayOfProblems = append(arrayOfProblems, problem.ProblemMsg)
			}
			t.Errorf("error(s) found in clean file %s:\n%s", path, strings.Join(arrayOfProblems, "\n"))
		}
		return nil
	})
}

func TestDirty(t *testing.T) {
	/*
	   Scan the directory of dirty files.
	   Each dirty file has a corresponding .exp file containing the errors expected from linting.
	   We loop through both the expected errors and the ones produced from the input.
	   Any expected error not found will be reported,
	   and any errors not specified in the .exp file will be reported as well.

	   see yaml/dirty/README.md for more information.
	*/
	t.Log("DIRTY TESTS:")
	root := "../yaml/dirty/"
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
		// lint call
		problems, err := lint.Lint(path, input)
		if err != nil {
			t.Errorf("error linting file %s: %s", path, err)
		}
		var arrayOfProblems []string
		for _, problem := range problems {
			arrayOfProblems = append(arrayOfProblems, problem.ProblemMsg)
		}

		expected := strings.Split(string(expectedBytes), "\n")
		for _, expProblem := range expected {
			if expProblem == "" {
				continue
			}
			if !contains(arrayOfProblems, expProblem) {
				t.Errorf("missing EXPECTED problem:\n%s", expProblem)
			}
		}

		for _, problem := range arrayOfProblems {
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
