package lint

import (
	"gopkg.in/yaml.v3"
)

func Lint(filename string, input []byte) ([]string, error) {
	sink := problemSink{Filename: filename}
	node := new(WorkflowNode)
	err := yaml.Unmarshal(input, &node)
	if err != nil {
		return nil, err
	}
	err = lintWorkflowRoot(&sink, node)
	return sink.Problems, err
}
