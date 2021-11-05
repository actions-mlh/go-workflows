package lint

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

// field checks ->
// oneOf, allOf, anyOf
// $ref -> nested checks for its reference (child node) -> reasonsing: "on" -> "check_run"
// oneOf, allOf, anyOf
// Properties(objects), Items(array of constants)
// Properties(objects), Items(array of constants)
// "type" -> scalar (string, bool, float, etc ...)

type workflowNode struct {
	Raw   *yaml.Node
	Value *workflowValue
}

func (node *workflowNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	if len(value.Content)%2 != 0 {
		return fmt.Errorf("%d:%d  error  Expected even number of key value pairs", node.Raw.Line, node.Raw.Column)
	}

	event := new(workflowValue)
	for i := 0; i < len(value.Content); i += 2 {
		keyEntry := value.Content[i]
		valueEntry := value.Content[i+1]
		eventKey := keyEntry.Value
		switch eventKey {
		case "name":
			event.Name = new(workflowNameNode)
			err := valueEntry.Decode(event.Name)
			if err != nil {
				return err
			}
		case "on":
			event.On = new(workflowOnNode)
			err := valueEntry.Decode(event.On)
			if err != nil {
				return err
			}
		case "env":
			event.Env = new(workflowEnvNode)
			err := valueEntry.Decode(event.Env)
			if err != nil {
				return err
			}
		case "defaults":
			event.Defaults = new(workflowDefaultsNode)
			err := valueEntry.Decode(event.Defaults)
			if err != nil {
				return err
			}
		case "concurrency":
			event.Concurrency = new(workflowConcurrencyNode)
			err := valueEntry.Decode(event.Concurrency)
			if err != nil {
				return err
			}
		case "jobs":
			event.Jobs = new(workflowJobsNode)
			err := valueEntry.Decode(event.Jobs)
			if err != nil {
				return err
			}
		}
	}
	node.Value = event
	return nil
}

type workflowValue struct {
	Name        *workflowNameNode        `yaml:"name"`
	On          *workflowOnNode          `yaml:"on"`
	Env         *workflowEnvNode         `yaml:"env"`
	Defaults    *workflowDefaultsNode    `yaml:"defaults"`
	Concurrency *workflowConcurrencyNode `yaml:"concurrency"`
	Jobs        *workflowJobsNode        `yaml:"jobs"`
}

// --------------------------------------------On----------------------------------------------------

type workflowNameNode struct {
	Raw   *yaml.Node
	Value string
}

func (node *workflowNameNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value

	scalarTypes := []string{"!!str"}
	contains := false
	for _, scalarType := range scalarTypes {
		if node.Raw.Tag == scalarType {
			contains = true
		}
	}
	if !contains {
		return fmt.Errorf("%d:%d  error  %s %s", node.Raw.Line, node.Raw.Column, "expected one of scalar types:", strings.Join(scalarTypes, ","))
	}

	return value.Decode(&node.Value)
}

// --------------------------------------------Name----------------------------------------------------

// --------------------------------------------On----------------------------------------------------

type workflowOnNode struct {
	Raw   *yaml.Node
	OneOf workflowOnOneOf
}

type workflowOnOneOf struct {
	ScalarNode   *OnEventConstants
	SequenceNode *[]OnEventConstants
	MappingNode  *WorkFlowOnValue
}

func (node *workflowOnNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value

	switch node.Raw.Kind {
	case yaml.ScalarNode:
		scalarTypes := []string{"!!str"}
		contains := false
		for _, scalarType := range scalarTypes {
			if node.Raw.Tag == scalarType {
				contains = true
			}
		}
		if !contains {
			return fmt.Errorf("%d:%d  error  %s %s", node.Raw.Line, node.Raw.Column, "expected one of scalar types:", strings.Join(scalarTypes, ","))
		}
		return value.Decode(&node.OneOf.ScalarNode)
	case yaml.SequenceNode:
		return value.Decode(&node.OneOf.SequenceNode)
	case yaml.MappingNode:
		value := node.Raw
		if len(value.Content)%2 != 0 {
			return fmt.Errorf("%d:%d  error  expected even number of key value pairs", node.Raw.Line, node.Raw.Column)
		}
		event := new(WorkFlowOnValue)
		for i := 0; i < len(value.Content); i += 2 {
			keyEntry := value.Content[i]
			valueEntry := value.Content[i+1]
			eventKey := keyEntry.Value
			switch eventKey {
			case "check_run":
				event.CheckRun = new(OnCheckRunNode)
				err := valueEntry.Decode(event.CheckRun)
				if err != nil {
					return err
				}
			case "check_suite":
				event.CheckSuite = new(OnCheckSuiteNode)
				err := valueEntry.Decode(event.CheckSuite)
				if err != nil {
					return err
				}
			case "create":
				event.Create = new(OnCreateNode)
				err := valueEntry.Decode(event.Create)
				if err != nil {
					return err
				}
			}
		}
		node.OneOf.MappingNode = event
	}
	return nil
}

type OnEventConstants string

const (
	OnEvent_CheckRun   OnEventConstants = "check_run"
	OnEvent_CheckSuite OnEventConstants = "check_suite"
	OnEvent_Create     OnEventConstants = "create"
	OnEvent_Delete     OnEventConstants = "delete"
	OnEvent_Deployment OnEventConstants = "deployment"
)

var OnEvent_Constants = []OnEventConstants{
	OnEvent_CheckRun,
	OnEvent_CheckSuite,
	OnEvent_Create,
	OnEvent_Delete,
	OnEvent_Deployment}

type WorkFlowOnValue struct {
	CheckRun   *OnCheckRunNode   `yaml:"check_run,omitempty"`
	CheckSuite *OnCheckSuiteNode `yaml:"check_suite,omitempty"`
	Create     *OnCreateNode     `yaml:"create,omitempty"`
}

type OnCheckRunNode struct {
	Raw   *yaml.Node
	OneOf OnCheckRunOneOf
}

type OnCheckRunOneOf struct {
	MappingNode *OnCheckRunValue
}

func (node *OnCheckRunNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value

	switch node.Raw.Kind {
	case yaml.MappingNode:
		value := node.Raw
		if len(value.Content)%2 != 0 {
			return fmt.Errorf("%d:%d  error  Expected even number of key value pairs", node.Raw.Line, node.Raw.Column)
		}
		event := new(OnCheckRunValue)
		for i := 0; i < len(value.Content); i += 2 {
			keyEntry := value.Content[i]
			valueEntry := value.Content[i+1]

			eventKey := keyEntry.Value
			switch eventKey {
			case "types":
				event.Types = new(CheckRunTypesNode)
				err := valueEntry.Decode(event.Types)
				if err != nil {
					return err
				}
			}
		}
		node.OneOf.MappingNode = event
		return nil
	default:
		return fmt.Errorf("%d:%d  error  Expected one of: map type", node.Raw.Line, node.Raw.Column)
	}
}

type OnCheckRunValue struct {
	Types *CheckRunTypesNode `yaml:"types,omitempty"`
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
	CheckRunTypes_Created         CheckRunTypesConstants = "create"
	CheckRunTypes_Rerequested     CheckRunTypesConstants = "rerequested"
	CheckRunTypes_Completed       CheckRunTypesConstants = "completed"
	CheckRunTypes_RequestedAction CheckRunTypesConstants = "requested_action"
)

type OnCheckSuiteNode struct {
	Raw   *yaml.Node
	OneOf *OnCheckSuiteOneOf
}

type OnCheckSuiteOneOf struct {
	MappingNode *OnCheckSuiteValue
}

type OnCheckSuiteValue struct {
	Types *CheckSuiteTypesNode `yaml:"types"`
}

func (node *OnCheckSuiteNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	switch node.Raw.Kind {
	case yaml.MappingNode:
		value := node.Raw
		if len(value.Content)%2 != 0 {
			return fmt.Errorf("%d:%d  error  Expected even number of key value pairs", node.Raw.Line, node.Raw.Column)
		}
		event := new(OnCheckSuiteValue)
		for i := 0; i < len(value.Content); i += 2 {
			keyEntry := value.Content[i]
			valueEntry := value.Content[i+1]
			eventKey := keyEntry.Value
			switch eventKey {
			case "types":
				event.Types = new(CheckSuiteTypesNode)
				err := valueEntry.Decode(event.Types)
				if err != nil {
					return err
				}
				// if "type" of the child does not contain "type": "null",
				// order of check -> child "type" -> $ref ("type")
				if event.Types.Value == nil {
					return fmt.Errorf("%d:%d  error  Unexpected one of: null type", node.Raw.Line, node.Raw.Column)
				}
			default:
				return fmt.Errorf("%d:%d  error  Expected one of: types", node.Raw.Line, node.Raw.Column)
			}
		}
		node.OneOf.MappingNode = event
		return nil
	default:
		return fmt.Errorf("%d:%d  error  Expected one of: map type", node.Raw.Line, node.Raw.Column)
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

// --------------------------------------------On----------------------------------------------------

// --------------------------------------------Defaults----------------------------------------------------

type workflowDefaultsNode struct {
	Raw   *yaml.Node
	Value *workflowDefaultsValue
}

func (node *workflowDefaultsNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value

	if len(value.Content)%2 != 0 {
		return fmt.Errorf("%d:%d  error  expected even number of key value pairs", node.Raw.Line, node.Raw.Column)
	}
	event := new(workflowDefaultsValue)
	for i := 0; i < len(value.Content); i += 2 {
		keyEntry := value.Content[i]
		valueEntry := value.Content[i+1]
		eventKey := keyEntry.Value
		switch eventKey {
		case "run":
			event.Run = new(DefaultsRunNode)
			err := valueEntry.Decode(event.Run)
			if err != nil {
				return err
			}
		}
	}
	node.Value = event
	return value.Decode(&node.Value)
}

type workflowDefaultsValue struct {
	Run *DefaultsRunNode `yaml:"run"`
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
	WorkingDirectory RunWorkingDirectoryNode `yaml:"working-directory"`
}

type RunShellNode struct {
	Raw   *yaml.Node
	Value RunShellAnyOfConstants
}

func (node *RunShellNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value

	switch node.Raw.Kind {
	case yaml.ScalarNode:
		if node.Raw.Tag != "!!str" {
			// this return the specific scalar value
			return fmt.Errorf("%d:%d  error  Expected any of: string type", node.Raw.Line, node.Raw.Column)
		}
		return value.Decode(&node.Value)
	default:
		// can change this to: Expect any of: (scalar, sequence, mapping) instead of its scalar value
		return fmt.Errorf("%d:%d  error  Expected any of: string type", node.Raw.Line, node.Raw.Column)
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
	Raw   *yaml.Node
	Value string
}

func (node *RunWorkingDirectoryNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

// --------------------------------------------Defaults----------------------------------------------------

// --------------------------------------------Concurrency----------------------------------------------------

type workflowConcurrencyNode struct {
	Raw   *yaml.Node
	OneOf workflowConcurrencyOneOf
}

type workflowConcurrencyOneOf struct {
	ScalarNode  *string
	MappingNode *workflowConcurrencyValue
}

func (node *workflowConcurrencyNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value

	switch node.Raw.Kind {
	case yaml.ScalarNode:
		return value.Decode(&node.OneOf.ScalarNode)
	case yaml.MappingNode:
		if err := value.Decode(&node.OneOf.MappingNode); err != nil {
			return err
		}
	default:
		return fmt.Errorf("%d:%d  error  Expected one of: string, map type", node.Raw.Line, node.Raw.Column)
	}

	err := func(value workflowConcurrencyValue) error {
		if value.Group == nil {
			return fmt.Errorf("%d:%d  error  Required keys: \"group\"", node.Raw.Line, node.Raw.Column)
		}
		return nil
	}(*node.OneOf.MappingNode) //checks if oneOf exists, if it does add oneOf and its type: "object" , "array", "scalar"
	if err != nil {
		return err
	}
	return nil
}

type workflowConcurrencyValue struct {
	Group            *ConcurrencyGroupNode            `yaml:"group"`
	CancelInProgress *ConcurrencyCancelInProgressNode `yaml:"cancel-in-progress"`
}

type ConcurrencyGroupNode struct {
	Raw   *yaml.Node
	Value string
}

type ConcurrencyCancelInProgressNode struct {
	Raw   *yaml.Node
	Value bool
}

func (node *ConcurrencyGroupNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

func (node *ConcurrencyCancelInProgressNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

// --------------------------------------------Concurrency----------------------------------------------------

// --------------------------------------------JOBS----------------------------------------------------

type workflowJobsNode struct {
	Raw   *yaml.Node
	Value []*workflowJobsPatternProperties
}

func (node *workflowJobsNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	if len(value.Content)%2 != 0 {
		return fmt.Errorf("%d:%d  error  Expected even number of key value pairs", node.Raw.Line, node.Raw.Column)
	}
	for i := 0; i < len(value.Content); i += 2 {
		keyEntry := value.Content[i]
		job := &workflowJobsPatternProperties{
			ID: keyEntry.Value,
		}
		if err := value.Decode(&job.PatternProperties); err != nil {
			return err // change stderr message
		}
		node.Value = append(node.Value, job)
	}

	return nil
}
type workflowJobsPatternProperties struct {
	ID                string `yaml:"-"`
	PatternProperties *JobsPatternPropertiesNode
}

type JobsPatternPropertiesNode struct {
	Raw   *yaml.Node
	Value *JobsPatternPropertiesValue
}

func (node *JobsPatternPropertiesNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value

	if len(value.Content)%2 != 0 {
		return fmt.Errorf("%d:%d  error  expected even number of key value pairs", node.Raw.Line, node.Raw.Column)
	}
	event := new(JobsPatternPropertiesValue)
	for i := 0; i < len(value.Content); i += 2 {
		keyEntry := value.Content[i]
		valueEntry := value.Content[i+1]
		eventKey := keyEntry.Value
		switch eventKey {
		case "name":
			event.Name = new(JobNameNode)
			err := valueEntry.Decode(event.Name)
			if err != nil {
				return err
			}
		case "needs":

		case "uses":

		}
	}
	node.Value = event
	return nil
}

type JobsPatternPropertiesValue struct {
	Name  *JobNameNode  `yaml:"name"`
	Needs *JobNeedsNode `yaml:"needs"`
	Uses  *JobUsesNode  `yaml:"uses"`
}

type JobNameNode struct {
	Raw   *yaml.Node
	Value string
}

func (node JobNameNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type JobNeedsNode struct {
	Raw *yaml.Node
}

type JobUsesNode struct {
	Raw *yaml.Node
}

// --------------------------------------------JOBS----------------------------------------------------


// --------------------------------------------ENV----------------------------------------------------
type workflowEnvNode struct {
	Raw   *yaml.Node
	Value []*workflowEnvValue
}

func (node *workflowEnvNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value

	if len(value.Content)%2 != 0 {
		return fmt.Errorf("%d:%d  error  Expected even number of key value pairs", node.Raw.Line, node.Raw.Column)
	}
	for i := 0; i < len(value.Content); i += 2 {
		keyEntry := value.Content[i]
		env := &workflowEnvValue{
			ID: keyEntry.Value,
			Properties: &EnvPropertiesNode{
				Raw: keyEntry,
			},
		}
		if err := value.Decode(&env.Properties); err != nil {
			return err 
		}
		node.Value = append(node.Value, env)
	}
	return nil
}

type workflowEnvValue struct {
	ID         string `yaml:"-"`
	Properties *EnvPropertiesNode
}

type EnvPropertiesNode struct {
	Raw *yaml.Node
	Value string
}

func (node *EnvPropertiesNode) UnmarshalYAML(value *yaml.Node) error {
	scalarTypes := []string{"!!str", "!!bool", "!!int", "!!float"}
	contains := false
	for _, scalarType := range scalarTypes {
		if node.Raw.Tag == scalarType {
			contains = true
		}
	}
	if !contains {
		return fmt.Errorf("%d:%d  error  %s %s", node.Raw.Line, node.Raw.Column, "expected one of scalar types:", strings.Join(scalarTypes, ", "))
	}

	return node.Raw.Decode(&node.Value)
}

// --------------------------------------------ENV----------------------------------------------------
