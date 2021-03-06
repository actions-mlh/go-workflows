package lint

import (
	"encoding/json"
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/actions-mlh/go-workflows/lint/sink"
	"github.com/actions-mlh/go-workflows/lint/workflow"

	"github.com/actions-mlh/go-workflows/lint/jobs"
	"github.com/actions-mlh/go-workflows/lint/name"
	"github.com/actions-mlh/go-workflows/lint/on"
	"github.com/actions-mlh/go-workflows/lint/root"
)

func Lint(filename string, input []byte) ([]sink.ProblemsValue, error) {
	sink := sink.ProblemSink{Filename: filename}
	node := new(workflow.WorkflowNode)
	if len(input) == 0 {
		// empty file
		return sink.Problems, nil
	}
	err := yaml.Unmarshal(input, &node)
	if err != nil {
		return sink.Problems, err
	}

	err = root.Lint(&sink, node)
	if err != nil {
		return sink.Problems, err
	}
	err = name.Lint(&sink, node.Value.Name)
	if err != nil {
		return sink.Problems, err
	}
	err = jobs.Lint(&sink, node.Value.Jobs)
	if err != nil {
		return sink.Problems, err
	}
	err = on.Lint(&sink, node.Value.On)
	return sink.Problems, nil
}

func Spew(input []byte) error {
	fmt.Println("~~ORIGINAL FILE:~~")
	fmt.Println(string(input))
	fmt.Println("~~PROCESSED INFO:~~")
	node := new(workflow.WorkflowNode)
	err := yaml.Unmarshal(input, &node)
	if err != nil {
		return err
	}
	bytes, err := json.Marshal(node)
	if err != nil {
		return err
	}
	var output map[string]interface{}
	err = json.Unmarshal(bytes, &output)
	if err != nil {
		return err
	}
	return printVal(output, 0)
}

func printVal(node map[string]interface{}, lvl int) error {
	for key, val := range node {
		if val == nil {
			continue
		}
		if key != "Raw" {
			spc := strings.Repeat(" ", lvl * 2)
			if _, ok := val.(map[string]interface{}); ok {
				fmt.Printf("%s%s:\n", spc, key)
				printVal(val.(map[string]interface{}), lvl + 1)
			} else if arr, ok := val.([]interface{}); ok {
				for key2, val2 := range arr {
					if _, ok := val2.(map[string]interface{}); ok {
						fmt.Printf("%s%s:\n", spc, key)
						printVal(val2.(map[string]interface{}), lvl + 1)
					} else {
						fmt.Printf("%s%v: %v\n", spc, key2, val2)
					}
				}
			} else {
				fmt.Printf("%s%v: %v\n", spc, key, val)
			}
		}
	}
	return nil
}
