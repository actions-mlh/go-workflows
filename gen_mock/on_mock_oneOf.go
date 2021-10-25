// steps:
// 1) get package name and imports

// 2) type %s struct, parent + child + Node
// 2a) Raw *yaml.Node
// 2b) Value Parent + Child + Value

// 3) func (node *%s) UnmarshalYAML(value *yaml.Node) error {, Parent + Child + Node
// 	      node.Raw = value
// 	      return value.Decode(&node.Value)
//    }

// 4) type %s struct, parent + child + Value struct
// 4a) Child.fieldName Parent + Child + Node `yaml:"lowercased(Child.fieldname)"`

package mock_gen_schema

import (
	"gopkg.in/yaml.v3"
	"fmt"
)

type WorkflowNode struct {
	Raw   *yaml.Node
	Value WorkflowValue
}

func (node *WorkflowNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type WorkflowValue struct {
	Name WorkflowNameNode `yaml:"name"`
	On WorkflowOnNode `yaml:"on"`
}

type WorkflowNameNode struct {
	Raw   *yaml.Node
	Value string
}

func (node *WorkflowNameNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type WorkflowOnNode struct {
	Raw  *yaml.Node
	OneOf WorkflowOnOneOf
}

// compare ref if visited dont create append, just reuse name 
type WorkflowOnOneOf struct {
	scalarNode *OnEventConstants
	sequenceNode *[]OnEventConstants
	mappingNode *[]WorkFlowOnValue
}

func (node *WorkflowOnNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value

	if err := WorkflowOnYAMLAssertType(node.Raw); err != nil {
		return err
	}

	switch node.Raw.Kind {
		case yaml.ScalarNode:
			return value.Decode(&node.OneOf.scalarNode)
		case yaml.SequenceNode:
			return value.Decode(&node.OneOf.sequenceNode)
		case yaml.MappingNode:
			return value.Decode(&node.OneOf.mappingNode)
		default:
			return fmt.Errorf("%d:%d	error	Expected one of: string, array, map type", node.Raw.Line, node.Raw.Column)
	}
}

func WorkflowOnYAMLAssertType(rawNode *yaml.Node) error {
	kinds := []yaml.Kind{yaml.ScalarNode, yaml.SequenceNode, yaml.MappingNode}
	containsFlag := false
	for _, kind := range kinds {
		if rawNode.Kind == kind {
			containsFlag = true
		}
	}
	if !containsFlag {
		return fmt.Errorf("%d:%d	error	Expected one of: string, array, map type", rawNode.Line, rawNode.Column)
	}

	return nil
}

type OnEventConstants string 

const (
	CheckRun OnEventConstants = "check_run"
	CheckSuite OnEventConstants = "check_suite"
	Create OnEventConstants = "create"
	Delete OnEventConstants = "delete"
	Deployment OnEventConstants = "deployment"
)

type WorkFlowOnValue struct {
	CheckRun OnCheckRunNode `yaml:"check_run,omitempty"`
	// CheckSuite OnCheckSuiteNode `yaml:"check_suite"` -> same as Checkrun
	Create OnCreateNode `yaml:"create,omitempty"`
	// Delete OnDeleteNode `yaml:"delete"` -> same as Create
}

type OnCheckRunNode struct {
	Raw *yaml.Node
	OneOf OnCheckRunOneOf 
}

type OnCheckRunOneOf struct {
	// scalarNode *CheckEventObjectNull --> omitempty gives output: Checkrun: 
	mappingNode *[]OnCheckRunValue
}

func (node *OnCheckRunNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value

	if err := OnCheckRunYAMLAssertType(node.Raw); err != nil {
		return err
	}

	switch node.Raw.Kind {
		case yaml.MappingNode:
			return value.Decode(&node.OneOf.mappingNode)
		default:
			return fmt.Errorf("%d:%d	error	Expected one of: string, array, map type", node.Raw.Line, node.Raw.Column)
	}
}

func OnCheckRunYAMLAssertType(rawNode *yaml.Node) error {
	kinds := []yaml.Kind{yaml.MappingNode}
	containsFlag := false
	for _, kind := range kinds {
		if rawNode.Kind == kind {
			containsFlag = true
		}
	}
	if !containsFlag {
		return fmt.Errorf("%d:%d	error	Expected one of: string, array, map type", rawNode.Line, rawNode.Column)
	}

	return nil
}

type OnCheckRunValue struct {
	Types CheckRunTypesNode `yaml:"types"`
}

type CheckRunTypesNode struct {
	Raw *yaml.Node
	Value *[]CheckRunTypesConstants
}

type CheckRunTypesConstants string

const (
	Created CheckRunTypesConstants = "create"
	Rerequested CheckRunTypesConstants = "rerequested"
	Completed CheckRunTypesConstants = "completed"
	RerequestedAction CheckRunTypesConstants = "rerequested_action"
)

type OnCreateNode struct {
	Raw *yaml.Node
	Value CreateEventObjectOneOf 
}

type CreateEventObjectOneOf struct {
	mappingNode *[]CreateEventObjectValue //if properties or items key do not exist within Create parent field remove this?
}

type CreateEventObjectValue struct {

}













// type WorkflowOnOneOf struct {
// 	WorkflowOnString *WorkflowOnStringNode 
// 	WorkflowOnArray *[]WorkflowOnArrayNode   
// }

// type WorkflowOnStringNode struct {
// 	Raw *yaml.Node
// 	Value string
// }

// func (node *WorkflowOnStringNode) UnmarshalYAML(value *yaml.Node) error {
// 	node.Raw = value
// 	return value.Decode(&node.Value)
// }

// type WorkflowOnArrayNode struct {
// 	Raw *yaml.Node
// 	Value string
// }

// func (node *WorkflowOnArrayNode) UnmarshalYAML(value *yaml.Node) error {
// 	node.Raw = value
// 	return value.Decode(&node.Value)
// }

