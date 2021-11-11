package lint

import (
	"gopkg.in/yaml.v3"

	"c2c-actions-mlh-workflow-parser/lint/sink"
	"c2c-actions-mlh-workflow-parser/lint/workflow"

	"c2c-actions-mlh-workflow-parser/lint/jobs"
	"c2c-actions-mlh-workflow-parser/lint/name"
	"c2c-actions-mlh-workflow-parser/lint/root"
)

func Lint(filename string, input []byte) ([]string, error) {
	sink := sink.ProblemSink{Filename: filename}
	node := new(workflow.WorkflowNode)

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
	return sink.Problems, nil
}
