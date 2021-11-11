package workflow

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

type WorkflowConcurrencyNode struct {
	Raw   *yaml.Node
	OneOf WorkflowConcurrencyOneOf
}

type WorkflowConcurrencyOneOf struct {
	ScalarNode  string
	MappingNode *WorkflowConcurrencyValue
}

func (node *WorkflowConcurrencyNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value

	switch node.Raw.Kind {
	case yaml.ScalarNode:
		// TYPE:string
		return value.Decode(&node.OneOf.ScalarNode)
	case yaml.MappingNode:
		if len(value.Content)%2 != 0 {
			return fmt.Errorf("%d:%d\terror\tCould not process concurrency: value.Contents has odd length, should be paired", node.Raw.Line, node.Raw.Column)
		}
		event := new(WorkflowConcurrencyValue)
		for i := 0; i < len(value.Content); i += 2 {
			keyEntry := value.Content[i]
			valueEntry := value.Content[i+1]
			eventKey := keyEntry.Value
			switch eventKey {
			case "group":
				event.Group = new(ConcurrencyGroupNode)
				err := valueEntry.Decode(event.Group)
				if err != nil {
					return err
				}
			case "cancel-in-progress":
				event.CancelInProgress = new(ConcurrencyCancelInProgressNode)
				err := valueEntry.Decode(event.CancelInProgress)
				if err != nil {
					return err
				}
			}
		}
		node.OneOf.MappingNode = event
	}
	return nil
}

type WorkflowConcurrencyValue struct {
	Group            *ConcurrencyGroupNode            `yaml:"group"`
	CancelInProgress *ConcurrencyCancelInProgressNode `yaml:"cancel-in-progress"`
}

type ConcurrencyGroupNode struct {
	Raw   *yaml.Node
	Value string
}

func (node *ConcurrencyGroupNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	// TYPE:string
	return value.Decode(&node.Value)
}

type ConcurrencyCancelInProgressNode struct {
	Raw   *yaml.Node
	Value string
}

func (node *ConcurrencyCancelInProgressNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	// TYPE:string
	value.Decode(&node.Value)
	return nil
}
