package util

import (
	"strings"
	"gopkg.in/yaml.v3"
	"c2c-actions-mlh-workflow-parser/sink"
	// "fmt"
)

func CheckRequiredKeys(raw *yaml.Node, sink *sink.ProblemSink, workflowKeyNodes []*yaml.Node, requiredKeys []string) error {
	keys := []string{}
	for _, node := range workflowKeyNodes {
		keys = append(keys, node.Value)
	}
		
	for _, key := range requiredKeys {
		if !contains(keys, key) {
			sink.Record(raw, "missing required key: %s", key)
		}
	}
	return nil
}

func CheckDuplicateKeys(sink *sink.ProblemSink, nodeKeys []*yaml.Node ) error {
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

func CheckNullPointer(sink *sink.ProblemSink, nodeKeys []*yaml.Node, nodeValues []*yaml.Node) error {
	for i := 0; i < len(nodeKeys); i++ {
		nodeKey := nodeKeys[i]
		nodeValue := nodeValues[i]

		if nodeValue.Tag == "!!null" {
			sink.Record(nodeKey, "\"%s\" key should not have an empty value", nodeKey.Value)
		}
	}

	return nil
}


func CheckUnexpectedKeys(sink *sink.ProblemSink, expectedKeys []string, nodeKeys []*yaml.Node) error {
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

func CheckUnexpectedScalarTypes(sink *sink.ProblemSink, raw *yaml.Node, scalarTypes []string) error {
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
