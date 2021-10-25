package parser

import (
	"gopkg.in/yaml.v3"
)

// If Parse is able to read the YAML file,
// Parse returns a map[string]interface{} containing the github action info.
// If it's a valid github action, err will be nil.
// If not, err will contain a description of what went wrong.
// If Parse is unable to read the YAML file, it will return nil, err.
func Parse(data []byte) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	err := yaml.Unmarshal(data, obj)
	if err != nil {
		return nil, err
	}

	return obj, err
}
