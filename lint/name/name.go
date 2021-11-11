package name

import (
	"c2c-actions-mlh-workflow-parser/lint/workflow"
	"c2c-actions-mlh-workflow-parser/lint/sink"
)

func Lint(sink *sink.ProblemSink, target *workflow.WorkflowNameNode) error {
	if target.Raw != nil && target.Raw.Tag != "!!str" {
		sink.Record(target.Raw, "unexpected scalar type: %s, expected scalar types: %s", target.Raw.Tag, "!!str")
	}
	return nil
}
