package jobs

import (
	"c2c-actions-mlh-workflow-parser/lint/workflow"
	"c2c-actions-mlh-workflow-parser/lint/sink"
)

func Lint(sink *sink.ProblemSink, target *workflow.WorkflowJobsNode) error {

	if target != nil && target.Raw != nil {
		if err := checkJobNames(sink, target.Raw); err != nil {
			return err
		}

		if err := checkCyclicDependencies(sink, target); err != nil {
			return err
		}
	}
	
	return nil
}
