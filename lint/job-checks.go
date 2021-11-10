package lint

import (
	"regexp"
	"gopkg.in/yaml.v3"
)

func checkJobNames(sink *problemSink, raw *yaml.Node) error {
	for i := 0; i < len(raw.Content); i += 2 {
		valid, err := regexp.MatchString("^[a-zA-Z_][a-zA-Z0-9-_]*$", raw.Content[i].Value)
		if err != nil {
			return err
		}
		if !valid {
			sink.record(raw.Content[i], "invalid job name \"%s\"", raw.Content[i].Value)
		}
	}
	return nil
}

func checkCyclicDependencies(sink *problemSink, target *WorkflowJobsNode) error {
	adjList := [][]string{}
	checked := make(map[string]bool)
	path := make(map[string]bool)

	for _, jobValue := range target.Value {
		next := jobValue.ID
		checked[next] = false
		path[next] = false

		if jobValue.PatternProperties.Value.Needs.OneOf.ScalarNode != nil {
			prev := *jobValue.PatternProperties.Value.Needs.OneOf.ScalarNode
			adjList = append(adjList, []string{next, prev})
		} else if jobValue.PatternProperties.Value.Needs.OneOf.SequenceNode != nil {
			for _, prev := range *jobValue.PatternProperties.Value.Needs.OneOf.SequenceNode {
				adjList = append(adjList, []string{next, prev})
			}
		}

	}
	needsMap := make(map[string][]string)
	for _, relation := range adjList {
		needsMap[relation[0]] = append(needsMap[relation[0]], relation[1])
	}

	for _, jobValue := range target.Value {
		currentJobName := jobValue.ID
		if isCyclic(currentJobName, needsMap, checked, path) {
			sink.record(jobValue.PatternProperties.Value.Needs.Raw, "contains cyclic dependencies")
		}
	}

	return nil
}

func isCyclic(currentJobName string, needsMap map[string][]string, checked map[string]bool, path map[string]bool) bool{
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
	for _, child := range needsMap[currentJobName] {
		childReturnValue = isCyclic(child, needsMap, checked, path)
		if childReturnValue { break }
	}

	// process itself(parent node) after children, by removing itself from path and checking itself
	path[currentJobName] = false
	checked[currentJobName] = true

	return childReturnValue
}