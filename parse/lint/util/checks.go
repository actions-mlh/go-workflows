package util

import (
	"strings"
	"gopkg.in/yaml.v3"
	"c2c-actions-mlh-workflow-parser/parse/sink"
	// "fmt"
)

func CheckRequiredKeys(raw *yaml.Node, sink *sink.ProblemSink, keys []*yaml.Node, requiredKeys map[string]bool) error {
	for _, key := range keys {
		if _, contains := requiredKeys[key.Value]; contains {
			requiredKeys[key.Value] = true
		}
	}

	required := []string{}
	for k, v := range requiredKeys {
		if !v{
			required = append(required, k)
		}
	}

	if len(required) != 0 {
		sink.Record(raw, "Required Keys: %s", strings.Join(required, ","))
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

// func CheckUnexpectedTypes() error {

// }