package root

import (
	"gopkg.in/yaml.v3"

	"c2c-actions-mlh-workflow-parser/lint/workflow"
	"c2c-actions-mlh-workflow-parser/lint/sink"
)

func Lint(sink *sink.ProblemSink, target *workflow.WorkflowNode) error {
	workflowKeyNodes := []*yaml.Node{}
	workflowValueNodes := []*yaml.Node{}

	for i := 0; i < len(target.Raw.Content); i += 2 {
		workflowKeyNodes = append(workflowKeyNodes, target.Raw.Content[i])
		workflowValueNodes = append(workflowValueNodes, target.Raw.Content[i+1])
	}

	if err := checkNullPointer(sink, workflowKeyNodes, workflowValueNodes); err != nil {
		return err
	}
	
	if err := checkDuplicateKeys(sink, workflowKeyNodes); err != nil {
		return err
	}
	
	requiredKeys := []string{"on", "jobs"}
	if err := checkRequiredKeys(sink, workflowKeyNodes, requiredKeys, target.Raw); err != nil {
		return err
	}

	expectedKeys := []string{"name", "on", "env", "defaults", "concurrency", "jobs", "permissions"}
	if err := checkUnexpectedKeys(sink, workflowKeyNodes, expectedKeys); err != nil {
		return err
	}

	nameType := []string{"!!str"}
	if target.Value.Name != nil {
		err := checkUnexpectedScalarTypes(sink, target.Value.Name.Raw, nameType)
		if err != nil {
			return err
		}
	}

	jobsType := []string{"!!map"}
	if target.Value.Jobs != nil  && target.Value.Jobs.Raw != nil {
		err := checkUnexpectedScalarTypes(sink, target.Value.Jobs.Raw, jobsType)
		if err != nil {
			return err
		}
	}
	
	return nil
}
