package name

import (
	"c2c-actions-mlh-workflow-parser/lint/sink"
	"c2c-actions-mlh-workflow-parser/lint/workflow"
)

func Lint(sink *sink.ProblemSink, target *workflow.WorkflowNameNode) error {
	if target != nil && target.Raw != nil && target.Raw.Tag != "!!str" {
		sink.Record(target.Raw, "unexpected scalar type: %s, expected scalar types: %s", target.Raw.Tag, "!!str")
	}
	return nil
}
