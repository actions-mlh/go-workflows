package jobs

import (
	"regexp"
	"gopkg.in/yaml.v3"

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

func checkJobNames(sink *sink.ProblemSink, raw *yaml.Node) error {
	for i := 0; i < len(raw.Content); i += 2 {
		valid, err := regexp.MatchString("^[a-zA-Z_][a-zA-Z0-9-_]*$", raw.Content[i].Value)
		if err != nil {
			return err
		}
		if !valid { 
			sink.Record(raw.Content[i], "invalid job name \"%s\"", raw.Content[i].Value)
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

		if jobValue.PatternProperties != nil {
			if jobValue.PatternProperties.Value.Needs != nil {
				if jobValue.PatternProperties.Value.Needs.OneOf.ScalarNode != nil {
					prev := *jobValue.PatternProperties.Value.Needs.OneOf.ScalarNode
					arrayOfjobNeedsRelations = append(arrayOfjobNeedsRelations, []string{next, prev})
				} else if jobValue.PatternProperties.Value.Needs.OneOf.SequenceNode != nil {
					for _, prev := range *jobValue.PatternProperties.Value.Needs.OneOf.SequenceNode {
						arrayOfjobNeedsRelations = append(arrayOfjobNeedsRelations, []string{next, prev})
					}
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

func isCyclic(currentJobName string, needsAdjacencyList map[string][]string, checked map[string]bool, path map[string]bool) bool{
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
		if childReturnValue { break }
	}

	// process itself(parent node) after children, by removing itself from path and checking itself
	path[currentJobName] = false
	checked[currentJobName] = true

	return childReturnValue
}
