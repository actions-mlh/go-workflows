package gen_mock

import (
	"fmt"

	"gopkg.in/yaml.v3"

	"c2c-actions-mlh-workflow-parser/parse/lint"
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
	Value *WorkflowValue // OneOf or Value or Scalar(string, int, bool, etc)
}

// when creating unmarshalYaml function, we check if required exists
func (node *WorkflowNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value

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
			// if "type" or oneOf -> array of "type", Does not exist
			// might be OneOf or Value -> so we just use Raw == nil check
		case "on":
			event.On = new(WorkflowOnNode)
			err := valueEntry.Decode(event.On)
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
		default:
			return fmt.Errorf("%d:%d  error  Expected: name, on, defaults, concurrency", node.Raw.Line, node.Raw.Column)
		}
	}
	node.Value = event

	// ----------------------------------------------------------------------------------
	// generated after our oneOf and properties check
	// check if child.Struct.Required field exists
	err := func(value WorkflowValue) error {
		if value.On == nil {
			return fmt.Errorf("%d:%d   error   Required keys: \"on\" \"jobs\"", node.Raw.Line, node.Raw.Column)
		}
		return nil
	}(*node.Value)

	if err != nil {
		return err
	}

	return nil
	// ----------------------------------------------------------------------------------
}

func (node WorkflowNode) Lint(sink *lint.ProblemSink) error {
	if len(node.Raw.Content)%2 != 0 {
		// Uneven set of key value pairs (this shouldn't happen)
		sink.Record(node.Raw, "%d:%d  error  Expected even number of key value pairs", node.Raw.Line, node.Raw.Column)
	}
	if node.Value.Name == nil {
		sink.Record(node.Raw, "%d:%d  error  missing name", node.Raw.Line, node.Raw.Column)
	}
	if node.Value.On == nil {
		sink.Record(node.Raw, "%d:%d  error  missing on", node.Raw.Line, node.Raw.Column)
	}
	if node.Value.Defaults == nil {
		sink.Record(node.Raw, "%d:%d  error  missing defaults", node.Raw.Line, node.Raw.Column)
	}
	if node.Value.Concurrency == nil {
		sink.Record(node.Raw, "%d:%d  error  missing concurrency", node.Raw.Line, node.Raw.Column)
	}
	if node.Value.Jobs == nil {
		sink.Record(node.Raw, "%d:%d  error  missing jobs", node.Raw.Line, node.Raw.Column)
	}
	return nil
}

type WorkflowValue struct { // created at parent level -> code buffer

	// appended to buffer once at child level
	Name *WorkflowNameNode `yaml:"name"`
	On   *WorkflowOnNode   `yaml:"on"`
	// Env  WorkflowEnvNode  `yaml:"env"`
	Defaults    *WorkflowDefaultsNode    `yaml:"defaults"`
	Concurrency *WorkflowConcurrencyNode `yaml:"concurrency"`
	Jobs        *WorkflowJobsNode        `yaml:"jobs"`
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
	OneOf WorkflowOnOneOfKind // Parent + Child + OneOf + (Type or Kind)
}

// read through oneOf array, create all "type"s
// there will be a "type" -> if not on same level, then check $ref
type WorkflowOnOneOfKind struct {
	ScalarNode   *OnEventConstants // decide whether to reuse Constants ?
	SequenceNode *[]OnEventConstants
	MappingNode  *WorkFlowOnValue
}

func (node *WorkflowOnNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value

	switch node.Raw.Kind {
	case yaml.ScalarNode:
		// handles -> "configuration" -> "string", "number", "boolean"
		// output.go -> create array called scalars := []{"string", "number", "boolean"}
		// !!str, !!float, !!int
		if node.Raw.Tag != "!!str" {
			return fmt.Errorf("%d:%d  error  Expected one of: string type", node.Raw.Line, node.Raw.Column)
		}
		return value.Decode(&node.OneOf.ScalarNode)
	case yaml.SequenceNode:
		// will handle if key exists in linter for now, unless we decide to change
		return value.Decode(&node.OneOf.SequenceNode)
	case yaml.MappingNode:
		value := node.Raw

		if len(value.Content)%2 != 0 {
			// Uneven set of key value pairs (this shouldn't happen)
			return fmt.Errorf("%d:%d  error  Expected even number of key value pairs", node.Raw.Line, node.Raw.Column)
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
				// might be OneOf or Value -> so we just use Raw == nil check
				// if event.CheckRun.Raw == nil {
				// 	return fmt.Errorf("%d:%d  error  Unexpected one of: null type", node.Raw.Line, node.Raw.Column)
				// }
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
			default:
				return fmt.Errorf("%d:%d  error  Expected one of: check_run, check_suite, create", node.Raw.Line, node.Raw.Column)
			}
		}
		node.OneOf.MappingNode = event
		return nil
	default:
		return fmt.Errorf("%d:%d  error  Expected one of: string, array, map type", node.Raw.Line, node.Raw.Column)
	}

	// --- no required array field
}

type OnEventConstants string

const (
	CheckRun   OnEventConstants = "check_run"
	CheckSuite OnEventConstants = "check_suite"
	Create     OnEventConstants = "create"
	Delete     OnEventConstants = "delete"
	Deployment OnEventConstants = "deployment"
)

type WorkFlowOnValue struct { // created at parent level -> code buffer

	// appended to buffer once at child level
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

// func (node *OnCheckRunNode) lint(sink *ProblemSink) {

// }
func (node *OnCheckRunNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value

	switch node.Raw.Kind {
	case yaml.MappingNode:
		value := node.Raw
		if len(value.Content)%2 != 0 {
			// Uneven set of key value pairs (this shouldn't happen)
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
				// if "type" of the child does not contain "type": "null",
				// order of check -> child "type" -> $ref ("type")
				if event.Types.Raw == nil {
					return fmt.Errorf("%d:%d  error  Unexpected one of: null type", node.Raw.Line, node.Raw.Column)
				}
			default:
				return fmt.Errorf("%d:%d  error  Expected Keys: \"types\"", node.Raw.Line, node.Raw.Column)
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
			// Uneven set of key value pairs (this shouldn't happen)
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

type WorkflowConcurrencyNode struct {
	Raw   *yaml.Node
	OneOf WorkflowConcurrencyOneOf
}

type WorkflowConcurrencyOneOf struct {
	ScalarNode  *string
	MappingNode *WorkflowConcurrencyValue
}

func (node *WorkflowConcurrencyNode) UnmarshalYAML(value *yaml.Node) error {
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

	err := func(value WorkflowConcurrencyValue) error {
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

type WorkflowConcurrencyValue struct {
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

type WorkflowJobsNode struct {
	Raw   *yaml.Node
	Value []*WorkflowJobsPatternProperties
}

type WorkflowJobsPatternProperties struct {
	ID                string `yaml:"-"`
	PatternProperties *JobsPatternPropertiesNode
}

func (node *WorkflowJobsNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value

	if len(value.Content)%2 != 0 {
		// Uneven set of key value pairs (this shouldn't happen)
		return fmt.Errorf("%d:%d  error  Expected even number of key value pairs", node.Raw.Line, node.Raw.Column)
	}
	for i := 0; i < len(value.Content); i += 2 {
		keyEntry := value.Content[i]

		job := &WorkflowJobsPatternProperties{
			ID: keyEntry.Value,
			PatternProperties: &JobsPatternPropertiesNode{
				ID: keyEntry.Value,
			},
		}
		if err := value.Decode(&job.PatternProperties); err != nil {
			return err // change stderr message
		}

		node.Value = append(node.Value, job)

	}

	return nil
}

type JobsPatternPropertiesNode struct {
	Raw   *yaml.Node
	Value JobsPatternPropertiesOneOfType //Since all oneOf "type"(s) are the same its Types not Kind
	ID    string
}

type JobsPatternPropertiesOneOfType struct {
	ReusableWorkflowCallJob *ReusableWorkflowCallJobValue
	// NormalJob               *NormalJobValue
}

func (node *JobsPatternPropertiesNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	// output.go get both required arrays and loop through contents to get values
	// fmt.Printf("%+v\n", node.ID)

	contentArr := func(content []*yaml.Node) []string {
		var contentArr []string

		for i := 0; i < len(value.Content); i += 2 {
			if node.ID == value.Content[i].Value {
				OuterContent := value.Content[i+1]
				for i := 0; i < len(OuterContent.Content); i += 2 {
					keyEntry := OuterContent.Content[i]
					contentArr = append(contentArr, keyEntry.Value)
				}
				break
			}
		}
		return contentArr
	}(value.Content)

	var jobsPatternPropertiesString string
	for _, content := range contentArr { // content -> string
		mapOfJobs := map[string][]string{
			"reusableWorkflowCallJob": []string{ //could be more than 1 required value, which is why we should keep this
				"uses",
			},
			"normalJob": []string{ //could be more than 1 required value, which is why we should keep this
				"runs-on",
			},
		}

		for keyName, require := range mapOfJobs {
			for _, requiredString := range require {
				if content == requiredString {
					jobsPatternPropertiesString = keyName
				}
				break
			}
			if jobsPatternPropertiesString != "" {
				break
			}
		}
	}

	// var event interface{}
	// switch jobsPatternPropertiesString {
	// case "reusableWorkflowCallJob":
	// 	event = (ReusableWorkflowCallJobValue)(*new(ReusableWorkflowCallJobValue))
	// case "normalJob":

	// default:

	// }

	// event = ReusableWorkflowCallJobValue(*new(ReusableWorkflowCallJobValue)) // type conversion, one type is an alias of another; ex. type Name string
	// fmt.Printf("%+v\n", event.(ReusableWorkflowCallJobValue).Name) //-> an actual type

	if len(value.Content)%2 != 0 {
		// Uneven set of key value pairs (this shouldn't happen)
		return fmt.Errorf("%d:%d  error  Expected even number of key value pairs", node.Raw.Line, node.Raw.Column)
	}
	for i := 0; i < len(value.Content); i += 2 {
		keyEntry := value.Content[i]
		// valueEntry := value.Content[i+1]
		eventKey := keyEntry.Value

		switch eventKey {
		case "name":
			// event.Name = new(ReusableWorkflowCallJobNameNode)
			// err := valueEntry.Decode(event.Name)
			// if err != nil {
			// 	return err
			// }
		case "needs":

		case "uses":

			// case "jobs":
			// 	event.Jobs = new(WorkflowJobsNode) --> new(Parent's Parent + Parent's Child + Node)1
			// 	err := valueEntry.Decode(event.Jobs)
			// 	if err != nil {
			// 		return err
			// 	}
			// 	if event.Jobs.Raw == nil {
			// 		return fmt.Errorf("%d:%d  error  Unexpected one of: null type", node.Raw.Line, node.Raw.Column)
			// 	}
			// default:
			// 	return fmt.Errorf("%d:%d  error  Expected: ", node.Raw.Line, node.Raw.Column)
		}
	}
	// node.Value = event

	return nil
}

type ReusableWorkflowCallJobValue struct {
	Name  ReusableWorkflowCallJobNameNode  `yaml:"name"`
	Needs ReusableWorkflowCallJobNeedsNode `yaml:"needs"`
	Uses  ReusableWorkflowCallJobUsesNode  `yaml:"uses"`
}

type ReusableWorkflowCallJobNameNode struct {
	Raw *yaml.Node
	
}

type ReusableWorkflowCallJobNeedsNode struct {
	Raw *yaml.Node
}

type ReusableWorkflowCallJobUsesNode struct {
	Raw *yaml.Node
}






// type JobsNormalJobValue struct {
// 	Name        *NormalJobNameNode        `yaml:"name"`
// 	Needs       *NormalJobNeedsNode       `yaml:"needs"`
// 	Permissions *NormalJobPermissionsNode `yaml:"permissions"`
// 	If          *NormalJobIfNode          `yaml:"if"`
// 	Uses        *NormalJobUsesNode        `yaml:"uses"`
// 	With        *NormalJobWithNode        `yaml:"with"`
// 	Secrets     *NormalJobSecretsNode     `yaml:"secrets"`
// }

// type NormalJobNameNode struct {
// 	Raw *yaml.Node

// }
// 	Raw *yaml.Node
// 	OneOf *WorkflowJobsOneOf
// }

// two types of oneOf check -> 
// 1) if "type" exists on same key level as "oneOf" -> oneOf refers to the Node
// 2) else: it refers to the Kind
// type WorkflowJobsOneOf struct {
// 	NormalJob *JobsNormalJobValue
// 	ReusableWorkflowCallJob *JobsReusableWorkflowCallJobValue
// }

// func (node *WorkflowJobsNode) UnmarshalYAML(value *yaml.Node) error {
// 	node.Raw = value

	
// }
