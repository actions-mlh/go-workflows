package lint

import (
	"gopkg.in/yaml.v3"
	"c2c-actions-mlh-workflow-parser/sink"
	"c2c-actions-mlh-workflow-parser/workflow"
)

func Lint(filename string, input []byte) ([]string, error) {
	sink := sink.ProblemSink{Filename: filename}
	node := new(workflow.WorkflowNode)
	err := yaml.Unmarshal(input, &node)
	if err != nil {
		return nil, err
	}
	err = lintWorkflowRoot(&sink, node)
	return sink.Problems, err
}
