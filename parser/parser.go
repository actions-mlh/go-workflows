package parser

import (
	"gopkg.in/yaml.v3"
)

// Lint returns true, -1, nil if the lint passes.
// If not, it returns false, lineNo, error.
func Parse(data []byte) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	err := yaml.Unmarshal(data, obj)

	return obj, err
}
