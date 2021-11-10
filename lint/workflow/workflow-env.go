package workflow

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

type WorkflowEnvNode struct {
	Raw   *yaml.Node
	Value []*WorkflowEnvValue
}

func (node *WorkflowEnvNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	if len(value.Content)%2 != 0 {
		return fmt.Errorf("%d:%d\terror\tCould not process env: value.Contents has odd length, should be paired", node.Raw.Line, node.Raw.Column)
	}
	for i := 0; i < len(value.Content); i += 2 {
		keyEntry := value.Content[i]
		valueEntry := value.Content[i+1]
		env := &WorkflowEnvValue{
			ID: keyEntry.Value,
		}
		if err := valueEntry.Decode(&env.Properties); err != nil {
			return err
		}
		node.Value = append(node.Value, env)
	}
	return nil
}

type WorkflowEnvValue struct {
	ID         string `yaml:"-"`
	Properties *EnvPropertiesNode
}

type EnvPropertiesNode struct {
	Raw   *yaml.Node
	Value *string
}

func (node *EnvPropertiesNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	// TYPE:string,bool,int
	return value.Decode(&node.Value)
}
