package workflow

import "gopkg.in/yaml.v3"

type WorkflowNameNode struct {
	Raw   *yaml.Node
	Value *string
}

func (node *WorkflowNameNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}


