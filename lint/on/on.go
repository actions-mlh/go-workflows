package on

import (
	"github.com/actions-mlh/go-workflows/lint/sink"
	"github.com/actions-mlh/go-workflows/lint/workflow"
)

func Lint(sink *sink.ProblemSink, target *workflow.WorkflowOnNode) error {
	if target != nil && target.Raw != nil {
		// sink.Record(target.Raw, "on working correctly")
	}
	return nil
}
