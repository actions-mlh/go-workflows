package on

import (
	"c2c-actions-mlh-workflow-parser/lint/sink"
	"c2c-actions-mlh-workflow-parser/lint/workflow"
)

func Lint(sink *sink.ProblemSink, target *workflow.WorkflowOnNode) error {
	if target != nil && target.Raw != nil {
		// sink.Record(target.Raw, "on working correctly")
	}
	return nil
}
