package main

import (
	"gopkg.in/yaml.v3"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"c2c-actions-mlh-workflow-parser/lint"
	"c2c-actions-mlh-workflow-parser/sink"
	"c2c-actions-mlh-workflow-parser/workflow"
)

func TestParse(t *testing.T) {
	root := "../yaml/clean/"
	filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() ||
			!(filepath.Ext(path) == ".yml" ||
				filepath.Ext(path) == ".yaml") {
			return nil
		}
		input, err := os.Open(path)
		if err != nil {
			return err
		}
		defer input.Close()

		sink := &sink.ProblemSink{Filename: path, Output: os.Stdout}
		node := new(workflow.WorkflowNode)
		
		if err := yaml.NewDecoder(input).Decode(&node); err != nil {
			return err
		}

		if err := lint.LintWorkflowRoot(sink, node); err != nil {
			return err
		}
		if len(sink.Problems) > 0 {
			t.Errorf("error(s) found in clean file %s:\n%s", path, strings.Join(sink.Problems, "\n"))
		}

		return nil
	})

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
		input, err := os.Open(path)
		if err != nil {
			return err
		}
		defer input.Close()

		expected, err := os.Open(path + ".exp")

		sink := &sink.ProblemSink{Filename: path, Output: os.Stdout}
		node := new(workflow.WorkflowNode)
		
		if err := yaml.NewDecoder(input).Decode(&node); err != nil {
			return err
		}

		if err := lint.LintWorkflowRoot(sink, node); err != nil {
			return err
		}
		if len(sink.Problems) > 0 {
			t.Errorf("error(s) found in clean file %s:\n%s", path, strings.Join(sink.Problems, "\n"))
		}

		return nil
	})
}
