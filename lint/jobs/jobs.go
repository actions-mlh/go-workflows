package jobs

import (
	"regexp"

	"c2c-actions-mlh-workflow-parser/lint/sink"
	"c2c-actions-mlh-workflow-parser/lint/workflow"
)

func Lint(sink *sink.ProblemSink, target *workflow.WorkflowJobsNode) error {

	if target != nil && target.Raw != nil {
		if err := checkJobType(sink, target); err != nil {
			return err
		}
		if err := checkJobNames(sink, target); err != nil {
			return err
		}
		if err := checkCyclicDependencies(sink, target); err != nil {
			return err
		}
	}

	return nil
}

func checkJobType(sink *sink.ProblemSink, target *workflow.WorkflowJobsNode) error {
	if target.Raw.Tag != "!!map" {
		sink.Record(target.Raw, "unexpected scalar type: %s, expected scalar types: %s", target.Raw.Tag, "!!map")
	}
	return nil
}

func checkJobNames(sink *sink.ProblemSink, target *workflow.WorkflowJobsNode) error {
	for i := 0; i < len(target.Raw.Content); i += 2 {
		valid, err := regexp.MatchString("^[a-zA-Z_][a-zA-Z0-9-_]*$", target.Raw.Content[i].Value)
		if err != nil {
			return err
		}
		if !valid {
			sink.Record(target.Raw.Content[i], "invalid job name \"%s\"", target.Raw.Content[i].Value)
		}
	}
	return nil
}

func checkCyclicDependencies(sink *sink.ProblemSink, target *workflow.WorkflowJobsNode) error {
	arrayOfjobNeedsRelations := [][]string{}
	checked := make(map[string]bool)
	path := make(map[string]bool)

	for _, jobValue := range target.Value {
		next := jobValue.ID
		checked[next] = false
		path[next] = false

		if jobValue.PatternProperties != nil &&
			jobValue.PatternProperties.Value.Needs != nil {

			oneOf := jobValue.PatternProperties.Value.Needs.OneOf
			if oneOf.ScalarNode != nil {
				prev := *oneOf.ScalarNode
				arrayOfjobNeedsRelations = append(arrayOfjobNeedsRelations, []string{next, prev})
			} else if oneOf.SequenceNode != nil {
				for _, prev := range *oneOf.SequenceNode {
					arrayOfjobNeedsRelations = append(arrayOfjobNeedsRelations, []string{next, prev})
				}
			}
		}
	}

	needsAdjacencyList := make(map[string][]string)
	for _, relation := range arrayOfjobNeedsRelations {
		needsAdjacencyList[relation[0]] = append(needsAdjacencyList[relation[0]], relation[1])
	}

	for _, jobValue := range target.Value {
		currentJobName := jobValue.ID
		if isCyclic(currentJobName, needsAdjacencyList, checked, path) {
			sink.Record(jobValue.PatternProperties.Value.Needs.Raw, "contains cyclic dependencies")
		}
	}

	return nil
}

func isCyclic(currentJobName string, needsAdjacencyList map[string][]string, checked map[string]bool, path map[string]bool) bool {
	// base cases
	if checked[currentJobName] { // no cycle is formed with currentJobName
		return false
	}
	if path[currentJobName] { // path already has true (marked) node, is cyclic
		return true
	}

	path[currentJobName] = true
	childReturnValue := false
	// scan children using postorder DFS
	for _, child := range needsAdjacencyList[currentJobName] {
		childReturnValue = isCyclic(child, needsAdjacencyList, checked, path)
		if childReturnValue {
			break
		}
	}

	// process itself(parent node) after children, by removing itself from path and checking itself
	path[currentJobName] = false
	checked[currentJobName] = true

	return childReturnValue
}
