package workflow

import "gopkg.in/yaml.v3"

type WorkflowEnvNode struct {
	Raw   *yaml.Node
	Value []*WorkflowEnvValue
}

func (node *WorkflowEnvNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	if len(value.Content)%2 != 0 {
		return fmt.Errorf("%d:%d  error  Expected even number of key value pairs", node.Raw.Line, node.Raw.Column)
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
	scalarTypes := []string{"!!str", "!!bool", "!!int"}
	contains := false
	for _, scalarType := range scalarTypes {
		if node.Raw.Tag == scalarType {
			contains = true
		}
	}
	if !contains {
		return fmt.Errorf("%d:%d  error  %s %s", node.Raw.Line, node.Raw.Column, "expected one of scalar types:", strings.Join(scalarTypes, ", "))
	}

	return value.Decode(&node.Value)
}
