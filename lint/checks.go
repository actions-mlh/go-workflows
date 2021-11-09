package lint

import (
	"strings"
	"gopkg.in/yaml.v3"
	"regexp"
)

func checkRequiredKeys(raw *yaml.Node, sink *problemSink, workflowKeyNodes []*yaml.Node, requiredKeys []string) error {
	keys := []string{}
	for _, node := range workflowKeyNodes {
		keys = append(keys, node.Value)
	}
		
	for _, key := range requiredKeys {
		if !contains(keys, key) {
			sink.record(raw, "missing required key: %s", key)
		}
	}
	return nil
}

func checkDuplicateKeys(sink *problemSink, nodeKeys []*yaml.Node ) error {
	nonDuplicateKeys := make(map[string]int) 

	for _, nodeKey := range nodeKeys {
		if _, contains := nonDuplicateKeys[nodeKey.Value]; !contains {
			nonDuplicateKeys[nodeKey.Value] = 1
		} else {
			nonDuplicateKeys[nodeKey.Value]++
			sink.record(nodeKey, "Duplicate Keys: %s", nodeKey.Value)
		}
	}
	return nil
}

func checkNullPointer(sink *problemSink, nodeKeys []*yaml.Node, nodeValues []*yaml.Node) error {
	for i := 0; i < len(nodeKeys); i++ {
		nodeKey := nodeKeys[i]
		nodeValue := nodeValues[i]

		if nodeValue.Tag == "!!null" {
			sink.record(nodeKey, "\"%s\" key should not have an empty value", nodeKey.Value)
		}
	}

	return nil
}


func checkUnexpectedKeys(sink *problemSink, expectedKeys []string, nodeKeys []*yaml.Node) error {
	for _, nodeKey := range nodeKeys {
		contains := false

		for _, expectedKey := range expectedKeys {
			if nodeKey.Value == expectedKey {
				contains = true
			}
		}

		if !contains {
			sink.record(nodeKey, "unexpected key \"%s\". expected one of keys %s", nodeKey.Value, strings.Join(expectedKeys, ", "))
		}
	}


	return nil
}

func checkUnexpectedScalarTypes(sink *problemSink, raw *yaml.Node, scalarTypes []string) error {
	contains := false
	for _, scalarType := range scalarTypes {
		if scalarType == raw.Tag {
			contains = true
		}
	}

	if !contains {
		sink.record(raw, "unexpected scalar type: %s, expected scalar types: %s", raw.Tag, strings.Join(scalarTypes, ","))
	}

	return nil
}

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

func contains(slice []string, item string) bool {
	for _, val := range slice {
		if val == item {
			return true
		}
	}
	return false
}
