package name

import (
	"github.com/actions-mlh/go-workflows/lint/sink"
	"github.com/actions-mlh/go-workflows/lint/workflow"
)

func Lint(sink *sink.ProblemSink, target *workflow.WorkflowNameNode) error {
	if target != nil && target.Raw != nil && target.Raw.Tag != "!!str" {
		sink.Record(target.Raw, "unexpected scalar type: %s, expected scalar types: %s", target.Raw.Tag, "!!str")
	}
	return nil
}
