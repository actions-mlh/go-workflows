package root

import (
	"strings"
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

func checkNullPointer(sink *sink.ProblemSink, nodeKeys []*yaml.Node, nodeValues []*yaml.Node) error {
	for i := 0; i < len(nodeKeys); i++ {
		nodeKey := nodeKeys[i]
		nodeValue := nodeValues[i]

		if nodeValue.Tag == "!!null" {
			sink.Record(nodeKey, "\"%s\" key should not have an empty value", nodeKey.Value)
		}
	}

	return nil
}

func checkDuplicateKeys(sink *sink.ProblemSink, nodeKeys []*yaml.Node ) error {
	nonDuplicateKeys := make(map[string]int) 

	for _, nodeKey := range nodeKeys {
		if _, contains := nonDuplicateKeys[nodeKey.Value]; !contains {
			nonDuplicateKeys[nodeKey.Value] = 1
		} else {
			nonDuplicateKeys[nodeKey.Value]++
			sink.Record(nodeKey, "Duplicate Keys: %s", nodeKey.Value)
		}
	}
	return nil
}

func checkRequiredKeys(sink *sink.ProblemSink, nodeKeys []*yaml.Node, requiredKeys []string, root *yaml.Node) error {
	keys := []string{}
	for _, node := range nodeKeys {
		keys = append(keys, node.Value)
	}
		
	for _, key := range requiredKeys {
		if !contains(keys, key) {
			sink.Record(root, "missing required key: %s", key)
		}
	}
	return nil
}

func checkUnexpectedKeys(sink *sink.ProblemSink, nodeKeys []*yaml.Node, expectedKeys []string) error {
	for _, nodeKey := range nodeKeys {
		contains := false

		for _, expectedKey := range expectedKeys {
			if nodeKey.Value == expectedKey {
				contains = true
			}
		}

		if !contains {
			sink.Record(nodeKey, "unexpected key \"%s\". expected one of keys %s", nodeKey.Value, strings.Join(expectedKeys, ", "))
		}
	}


	return nil
}

func checkUnexpectedScalarTypes(sink *sink.ProblemSink, raw *yaml.Node, scalarTypes []string) error {
	contains := false
	for _, scalarType := range scalarTypes {
		if scalarType == raw.Tag {
			contains = true
		}
	}

	if !contains {
		sink.Record(raw, "unexpected scalar type: %s, expected scalar types: %s", raw.Tag, strings.Join(scalarTypes, ","))
	}

	return nil
}

func contains(slice []string, item string) bool {
	for _, val := range slice {
		if val == item {
			return true
		}
	}
	return false
}
