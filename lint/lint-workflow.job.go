package lint

import (
	// "gopkg.in/yaml.v3"
	// "fmt"
)

func lintWorkflowJobs(sink *problemSink, target *WorkflowJobsNode) error {

	if target != nil  && target.Raw != nil {
		if err := checkJobNames(sink, target.Raw); err != nil {
			return err
		}

		// if err := checkCyclicDependencies(sink, target); err != nil {
		// 	return err
		// }

	}
	
	return nil
}
