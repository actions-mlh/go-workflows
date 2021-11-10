package workflow

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

// field checks ->
// oneOf, allOf, anyOf
// $ref -> nested checks for its reference (child node) -> reasonsing: "on" -> "check_run"
// oneOf, allOf, anyOf
// Properties(objects), Items(array of constants)
// Properties(objects), Items(array of constants)
// "type" -> scalar (string, bool, float, etc ...)

type WorkflowNode struct {
	Raw   *yaml.Node
	Value *WorkflowValue
}

func (node *WorkflowNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	if len(value.Content)%2 != 0 {
		return fmt.Errorf("%d:%d\terror\tCould not process root: value.Contents has odd length, should be paired", node.Raw.Line, node.Raw.Column)
	}

	event := new(WorkflowValue)
	for i := 0; i < len(value.Content); i += 2 {
		keyEntry := value.Content[i]
		valueEntry := value.Content[i+1]
		eventKey := keyEntry.Value
		switch eventKey {
		case "name":
			event.Name = new(WorkflowNameNode)
			err := valueEntry.Decode(event.Name)
			if err != nil {
				return err
			}
		case "on":
			event.On = new(WorkflowOnNode)
			err := valueEntry.Decode(event.On)
			if err != nil {
				return err
			}
		case "env":
			event.Env = new(WorkflowEnvNode)
			err := valueEntry.Decode(event.Env)
			if err != nil {
				return err
			}
		case "defaults":
			event.Defaults = new(WorkflowDefaultsNode)
			err := valueEntry.Decode(event.Defaults)
			if err != nil {
				return err
			}
		case "concurrency":
			event.Concurrency = new(WorkflowConcurrencyNode)
			err := valueEntry.Decode(event.Concurrency)
			if err != nil {
				return err
			}
		case "jobs":
			event.Jobs = new(WorkflowJobsNode)
			err := valueEntry.Decode(event.Jobs)
			if err != nil {
				return err
			}
		case "permissions":
			event.Permissions = new(WorkflowPermissionsNode)
			err := valueEntry.Decode(event.Permissions)
			if err != nil {
				return err
			}
		}
	}
	node.Value = event
	return nil
}

type WorkflowValue struct {
	Name        *WorkflowNameNode        `yaml:"name"`
	On          *WorkflowOnNode          `yaml:"on"`
	Env         *WorkflowEnvNode         `yaml:"env"`
	Defaults    *WorkflowDefaultsNode    `yaml:"defaults"`
	Concurrency *WorkflowConcurrencyNode `yaml:"concurrency"`
	Jobs        *WorkflowJobsNode        `yaml:"jobs"`
	Permissions *WorkflowPermissionsNode `yaml:"permissions"`
}
