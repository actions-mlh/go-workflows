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

package gen_mock

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type WorkflowNode struct {
	Raw   *yaml.Node
	Value WorkflowValue
}


// when creating unmarshalYaml function, we check if required exists
func (node *WorkflowNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	// required := []string{"on"}


	// only do this if required exists
	if err := value.Decode(&node.Value); err != nil {
		return err
	}

	
	// fmt.Printf("%+v\n", *&node.Value)	

	return nil
	
}

type WorkflowValue struct {
	Name *WorkflowNameNode `yaml:"name"`
	On   *WorkflowOnNode   `yaml:"on"`
	// Env  WorkflowEnvNode  `yaml:"env"`
	Defaults *WorkflowDefaultsNode `yaml:"defaults"`
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
	Raw   *yaml.Node
	OneOf WorkflowOnOneOf
}

// compare ref if visited dont create append, just reuse name
type WorkflowOnOneOf struct {
	ScalarNode   *OnEventConstants
	SequenceNode *[]OnEventConstants
	MappingNode  *WorkFlowOnValue
}

func (node *WorkflowOnNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value

	switch node.Raw.Kind {
	case yaml.ScalarNode:
		return value.Decode(&node.OneOf.ScalarNode)
	case yaml.SequenceNode:
		return value.Decode(&node.OneOf.SequenceNode)
	case yaml.MappingNode:
		return value.Decode(&node.OneOf.MappingNode)
	default:
		return fmt.Errorf("%d:%d	error	Expected one of: string, array, map type", node.Raw.Line, node.Raw.Column)
	}
}

type OnEventConstants string

const (
	CheckRun   OnEventConstants = "check_run"
	CheckSuite OnEventConstants = "check_suite"
	Create     OnEventConstants = "create"
	Delete     OnEventConstants = "delete"
	Deployment OnEventConstants = "deployment"
)

type WorkFlowOnValue struct {
	CheckRun   *OnCheckRunNode   `yaml:"check_run,omitempty"`
	CheckSuite *OnCheckSuiteNode `yaml:"check_suite"`
	Create     *OnCreateNode     `yaml:"create,omitempty"`
	// Delete OnDeleteNode `yaml:"delete"` -> same as Create
}

type OnCheckRunNode struct {
	Raw   *yaml.Node
	OneOf OnCheckRunOneOf
}

type OnCheckRunOneOf struct {
	// scalarNode *CheckEventObjectNull --> omitempty gives output: Checkrun:
	MappingNode *OnCheckRunValue
}

func (node *OnCheckRunNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value

	switch node.Raw.Kind {
	case yaml.MappingNode:
		return value.Decode(&node.OneOf.MappingNode)
	default:
		return fmt.Errorf("%d:%d	error	Expected one of: string, array, map type", node.Raw.Line, node.Raw.Column)
	}
}

type OnCheckRunValue struct {
	Types CheckRunTypesNode `yaml:"types"`
}

type CheckRunTypesNode struct {
	Raw   *yaml.Node
	Value *[]CheckRunTypesConstants
}

func (node *CheckRunTypesNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type CheckRunTypesConstants string

const (
	CheckRunTypes_Created           CheckRunTypesConstants = "create"
	CheckRunTypes_Rerequested       CheckRunTypesConstants = "rerequested"
	CheckRunTypes_Completed         CheckRunTypesConstants = "completed"
	CheckRunTypes_RerequestedAction CheckRunTypesConstants = "rerequested_action"
)

type OnCheckSuiteNode struct {
	Raw   *yaml.Node
	OneOf OnCheckSuiteOneOf
}

type OnCheckSuiteOneOf struct {
	MappingNode *OnCheckSuiteValue
}

type OnCheckSuiteValue struct {
	Types CheckSuiteTypesNode `yaml:"types"`
}

func (node *OnCheckSuiteNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value

	switch node.Raw.Kind {
	case yaml.MappingNode:
		return value.Decode(&node.OneOf.MappingNode)
	default:
		return fmt.Errorf("%d:%d	error	Expected one of: string, array, map type", node.Raw.Line, node.Raw.Column)
	}
}

type CheckSuiteTypesNode struct {
	Raw   *yaml.Node
	Value *[]CheckSuiteTypesConstants
}

func (node *CheckSuiteTypesNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type CheckSuiteTypesConstants string

const (
	CheckSuiteTypes_Completed   CheckSuiteTypesConstants = "completed"
	CheckSuiteTypes_Requested   CheckSuiteTypesConstants = "requested"
	CheckSuiteTypes_Rerequested CheckSuiteTypesConstants = "rerequested"
)

type OnCreateNode struct {
	Raw   *yaml.Node
	Value *CreateEventObjectOneOf
}

type CreateEventObjectOneOf struct {
	MappingNode *CreateEventObjectValue //if properties or items key do not exist within Create parent field remove this?
}

type CreateEventObjectValue struct {
}

type WorkflowDefaultsNode struct {
	Raw   *yaml.Node
	Value *WorkflowDefaultsValue
}

func (node *WorkflowDefaultsNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type WorkflowDefaultsValue struct {
	Run DefaultsRunNode `yaml:"run"`
}

type DefaultsRunNode struct {
	Raw   *yaml.Node
	Value DefaultsRunValue
}

func (node *DefaultsRunNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type DefaultsRunValue struct {
	Shell            RunShellNode            `yaml:"shell"`
	WorkingDirectory RunWorkingDirectoryNode `yaml:"working_directory"`
}

type RunShellNode struct {
	Raw   *yaml.Node
	Value RunShellAnyOfConstants
}

func (node *RunShellNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value

	switch node.Raw.Kind {
	case yaml.ScalarNode:
		return value.Decode(&node.Value)
	default:
		return fmt.Errorf("%d:%d	error	Expected one of: string type", node.Raw.Line, node.Raw.Column)
	}
}

type RunShellAnyOfConstants string

const (
	RunShell_Bash       RunShellAnyOfConstants = "bash"
	RunShell_Pwsh       RunShellAnyOfConstants = "pwsh"
	RunShell_Python     RunShellAnyOfConstants = "python"
	RunShell_Sh         RunShellAnyOfConstants = "sh"
	RunShell_Cmd        RunShellAnyOfConstants = "cmd"
	RunShell_Powershell RunShellAnyOfConstants = "powershell"
)

type RunWorkingDirectoryNode struct {
	Raw *yaml.Node
	Value string
}

func (node *RunWorkingDirectoryNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}