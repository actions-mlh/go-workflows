package lint

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"strings"
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
		return fmt.Errorf("%d:%d  error  Expected even number of key value pairs", node.Raw.Line, node.Raw.Column)
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
			event.Permissions = new(DefinitionPermissionsNode)
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
	Name        *WorkflowNameNode          `yaml:"name"`
	On          *WorkflowOnNode            `yaml:"on"`
	Env         *WorkflowEnvNode           `yaml:"env"`
	Defaults    *WorkflowDefaultsNode      `yaml:"defaults"`
	Concurrency *WorkflowConcurrencyNode   `yaml:"concurrency"`
	Jobs        *WorkflowJobsNode          `yaml:"jobs"`
	Permissions *DefinitionPermissionsNode `yaml:"permissions"`
}

// --------------------------------------------On----------------------------------------------------

type WorkflowNameNode struct {
	Raw   *yaml.Node
	Value string
}

func (node *WorkflowNameNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	// TYPE:string
	return value.Decode(&node.Value)
}

// --------------------------------------------Name----------------------------------------------------

// --------------------------------------------On----------------------------------------------------

type WorkflowOnNode struct {
	Raw   *yaml.Node
	OneOf WorkflowOnOneOf
}

type WorkflowOnOneOf struct {
	ScalarNode   *OnEventConstants
	SequenceNode *[]OnEventConstants
	MappingNode  *WorkFlowOnValue
}

func (node *WorkflowOnNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value

	switch node.Raw.Kind {
	case yaml.ScalarNode:
		// TYPE:string
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

type WorkflowDefaultsNode struct {
	Raw   *yaml.Node
	Value *WorkflowDefaultsValue
}

func (node *WorkflowDefaultsNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value

	if len(value.Content)%2 != 0 {
		return fmt.Errorf("%d:%d  error  expected even number of key value pairs", node.Raw.Line, node.Raw.Column)
	}
	event := new(WorkflowDefaultsValue)
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

type WorkflowDefaultsValue struct {
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
	Shell            *RunShellNode            `yaml:"shell"`
	WorkingDirectory *RunWorkingDirectoryNode `yaml:"working-directory"`
}

type RunShellNode struct {
	Raw   *yaml.Node
	Value *RunShellConstants
}

func (node *RunShellNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value

	switch node.Raw.Kind {
	case yaml.ScalarNode:
		// TYPE:string
		return value.Decode(&node.Value)
	default:
		return fmt.Errorf("%d:%d  error  Expected any of: string type", node.Raw.Line, node.Raw.Column)
	}
}

type RunShellConstants string

const (
	RunShell_Bash       RunShellConstants = "bash"
	RunShell_Pwsh       RunShellConstants = "pwsh"
	RunShell_Python     RunShellConstants = "python"
	RunShell_Sh         RunShellConstants = "sh"
	RunShell_Cmd        RunShellConstants = "cmd"
	RunShell_Powershell RunShellConstants = "powershell"
)

var RunShell_Constants = []RunShellConstants{
	RunShell_Bash,
	RunShell_Pwsh,
	RunShell_Python,
	RunShell_Sh,
	RunShell_Cmd,
	RunShell_Powershell,
}

type RunWorkingDirectoryNode struct {
	Raw   *yaml.Node
	Value string
}

func (node *RunWorkingDirectoryNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	// TYPE:string
	return node.Raw.Decode(&node.Value)
}

// --------------------------------------------Defaults----------------------------------------------------

// --------------------------------------------Concurrency----------------------------------------------------

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
			return fmt.Errorf("%d:%d  error  expected even number of key value pairs", node.Raw.Line, node.Raw.Column)
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

// --------------------------------------------Concurrency----------------------------------------------------

// --------------------------------------------JOBS----------------------------------------------------

type WorkflowJobsNode struct {
	Raw   *yaml.Node
	Value []*WorkflowJobsPatternProperties
}

func (node *WorkflowJobsNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	if len(value.Content)%2 != 0 {
		return fmt.Errorf("%d:%d  error  Expected even number of key value pairs", node.Raw.Line, node.Raw.Column)
	}
	for i := 0; i < len(value.Content); i += 2 {
		keyEntry := value.Content[i]
		valueEntry := value.Content[i+1]
		job := &WorkflowJobsPatternProperties{
			ID: keyEntry.Value,
		}
		if err := valueEntry.Decode(&job.PatternProperties); err != nil {
			return err
		}
		node.Value = append(node.Value, job)
	}

	return nil
}

type WorkflowJobsPatternProperties struct {
	ID                string `yaml:"-"`
	PatternProperties *JobsPatternPropertiesNode
}

type JobsPatternPropertiesNode struct {
	Raw   *yaml.Node
	Value *JobsPatternPropertiesValue
}

func (node *JobsPatternPropertiesNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	if len(node.Raw.Content)%2 != 0 {
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
			event.Needs = new(JobNeedsNode)
			err := valueEntry.Decode(event.Needs)
			if err != nil {
				return err
			}
		case "permissions":
			event.Permissions = new(JobsPermissionsEventNode)
			err := valueEntry.Decode(event.Permissions)
			if err != nil {
				return err
			}
		case "if":
			event.If = new(JobIfNode)
			err := valueEntry.Decode(event.If)
			if err != nil {
				return err
			}
		case "uses":
			event.Uses = new(JobUsesNode)
			err := valueEntry.Decode(event.Uses)
			if err != nil {
				return err
			}
		case "with":
			event.With = new(JobWithNode)
			err := valueEntry.Decode(event.With)
			if err != nil {
				return err
			}
		case "secrets":
			event.Secrets = new(JobSecretsNode)
			err := valueEntry.Decode(event.Secrets)
			if err != nil {
				return err
			}
		case "runs-on":
			event.RunsOn = new(JobRunsOnNode)
			err := valueEntry.Decode(event.RunsOn)
			if err != nil {
				return err
			}
		case "environment":
			event.Environment = new(JobEnvironmentNode)
			err := valueEntry.Decode(event.Environment)
			if err != nil {
				return err
			}
		case "outputs":
			event.Outputs = new(JobOutputsNode)
			err := valueEntry.Decode(event.Outputs)
			if err != nil {
				return err
			}
		case "env":
			event.Env = new(JobEnvNode)
			err := valueEntry.Decode(event.Env)
			if err != nil {
				return err
			}
		case "defaults":
			event.Defaults = new(JobDefaultsNode)
			err := valueEntry.Decode(event.Defaults)
			if err != nil {
				return err
			}
		case "steps":
			event.Steps = new(JobStepsNode)
			err := valueEntry.Decode(event.Steps)
			if err != nil {
				return err
			}
		case "timeout-minutes":
			event.TimeoutMinutes = new(JobTimeoutMinutesNode)
			err := valueEntry.Decode(event.TimeoutMinutes)
			if err != nil {
				return err
			}
		case "continue-on-error":
			event.ContinueOnError = new(JobContinueOnErrorNode)
			err := valueEntry.Decode(event.ContinueOnError)
			if err != nil {
				return err
			}
		case "container":
			event.Container = new(JobContainerNode)
			err := valueEntry.Decode(event.Container)
			if err != nil {
				return err
			}
		case "services":
			event.Services = new(JobServicesNode)
			err := valueEntry.Decode(event.Services)
			if err != nil {
				return err
			}
		case "concurrency":
			event.Concurrency = new(JobConcurrencyNode)
			err := valueEntry.Decode(event.Concurrency)
			if err != nil {
				return err
			}
		}
	}
	node.Value = event
	return nil
}

type JobsPermissionsEventNode struct {
	Raw   *yaml.Node
	Value *DefinitionPermissionsValue
}

func (node *JobsPermissionsEventNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value

	if len(value.Content)%2 != 0 {
		return fmt.Errorf("%d:%d  error  expected even number of key value pairs", node.Raw.Line, node.Raw.Column)
	}
	event := new(DefinitionPermissionsValue)
	for i := 0; i < len(value.Content); i += 2 {
		keyEntry := value.Content[i]
		valueEntry := value.Content[i+1]
		eventKey := keyEntry.Value
		switch eventKey {
		case "actions":
			event.Actions = new(PermissionsActionsNode)
			err := valueEntry.Decode(event.Actions)
			if err != nil {
				return err
			}
		case "checks":
			event.Checks = new(PermissionsChecksNode)
			err := valueEntry.Decode(event.Checks)
			if err != nil {
				return err
			}
		case "contents":
			event.Contents = new(PermissionsContentsNode)
			err := valueEntry.Decode(event.Contents)
			if err != nil {
				return err
			}
		case "deployments":
			event.Deployments = new(PermissionsDeploymentsNode)
			err := valueEntry.Decode(event.Deployments)
			if err != nil {
				return err
			}
		case "issues":
			event.Issues = new(PermissionsIssuesNode)
			err := valueEntry.Decode(event.Issues)
			if err != nil {
				return err
			}
		case "packages":
			event.Packages = new(PermissionsPackagesNode)
			err := valueEntry.Decode(event.Packages)
			if err != nil {
				return err
			}
		case "pull-requests":
			event.PullRequests = new(PermissionsPullRequestsNode)
			err := valueEntry.Decode(event.PullRequests)
			if err != nil {
				return err
			}
		case "repository-projects":
			event.RepositoryProjects = new(PermissionsRepositoryProjectsNode)
			err := valueEntry.Decode(event.RepositoryProjects)
			if err != nil {
				return err
			}
		case "security-events":
			event.SecurityEvents = new(PermissionsSecurityEventsNode)
			err := valueEntry.Decode(event.SecurityEvents)
			if err != nil {
				return err
			}
		case "statuses":
			event.Statuses = new(PermissionsStatusesNode)
			err := valueEntry.Decode(event.Statuses)
			if err != nil {
				return err
			}
		case "id-token":
			event.IdToken = new(PermissionsIdTokenNode)
			err := valueEntry.Decode(event.IdToken)
			if err != nil {
				return err
			}
		}
	}
	node.Value = event
	return nil
}

type JobsPatternPropertiesValue struct {
	Name           *JobNameNode              `yaml:"name"`
	Needs          *JobNeedsNode             `yaml:"needs"`
	Permissions    *JobsPermissionsEventNode `yaml:"permissions"`
	If             *JobIfNode                `yaml:"if"`
	Uses           *JobUsesNode              `yaml:"uses"`
	With           *JobWithNode              `yaml:"with"`
	Secrets        *JobSecretsNode           `yaml:"secrets"`
	RunsOn         *JobRunsOnNode            `yaml:"runs-on"`
	Environment    *JobEnvironmentNode       `yaml:"environment"`
	Outputs        *JobOutputsNode           `yaml:"outputs"`
	Env            *JobEnvNode               `yaml:"env"`
	Defaults       *JobDefaultsNode          `yaml:"defaults"`
	Steps          *JobStepsNode             `yaml:"steps"`
	TimeoutMinutes *JobTimeoutMinutesNode    `yaml:"timeout-minutes"`
	// Strategy *JobStrategyNode `yaml:"strategy"`
	ContinueOnError *JobContinueOnErrorNode `yaml:"continue-on-error"`
	Container       *JobContainerNode       `yaml:"container"`
	Services        *JobServicesNode        `yaml:"services"`
	Concurrency     *JobConcurrencyNode     `yaml:"concurrency"`
}

type JobNameNode struct {
	Raw   *yaml.Node
	Value string
}

func (node *JobNameNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	scalarTypes := []string{"!!str"}
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

type JobNeedsNode struct {
	Raw   *yaml.Node
	OneOf JobNeedsOneOf
}

type JobNeedsOneOf struct {
	ScalarNode   string
	SequenceNode *[]string
}

func (node *JobNeedsNode) UnmarshalYAML(value *yaml.Node) error {
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
			return fmt.Errorf("%d:%d  error  %s %s", node.Raw.Line, node.Raw.Column, "expected one of scalar types:", strings.Join(scalarTypes, ", "))
		}
		return value.Decode(&node.OneOf.ScalarNode)
	case yaml.SequenceNode:
		return value.Decode(&node.OneOf.SequenceNode)
	}
	return nil
}

type JobIfNode struct {
	Raw   *yaml.Node
	Value *string
}

func (node *JobIfNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	scalarTypes := []string{"!!str"}
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

type JobUsesNode struct {
	Raw   *yaml.Node
	Value *string
}

func (node *JobUsesNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	scalarTypes := []string{"!!str"}
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

type JobWithNode struct {
	Raw   *yaml.Node
	Value []*JobWithValue
}

func (node *JobWithNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	if len(value.Content)%2 != 0 {
		return fmt.Errorf("%d:%d  error  Expected even number of key value pairs", node.Raw.Line, node.Raw.Column)
	}
	for i := 0; i < len(value.Content); i += 2 {
		keyEntry := value.Content[i]
		valueEntry := value.Content[i+1]
		with := &JobWithValue{
			ID: keyEntry.Value,
		}
		if err := valueEntry.Decode(&with.Properties); err != nil {
			return err
		}
		node.Value = append(node.Value, with)
	}
	return nil
}

type JobWithValue struct {
	ID         string `yaml:"-"`
	Properties *WithPropertiesNode
}

type WithPropertiesNode struct {
	Raw   *yaml.Node
	Value *string
}

func (node *WithPropertiesNode) UnmarshalYAML(value *yaml.Node) error {
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

	return node.Raw.Decode(&node.Value)
}

type JobSecretsNode struct {
	Raw   *yaml.Node
	Value []*JobSecretsValue
}

func (node *JobSecretsNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	if len(value.Content)%2 != 0 {
		return fmt.Errorf("%d:%d  error  Expected even number of key value pairs", node.Raw.Line, node.Raw.Column)
	}
	for i := 0; i < len(value.Content); i += 2 {
		keyEntry := value.Content[i]
		valueEntry := value.Content[i+1]
		secrets := &JobSecretsValue{
			ID: keyEntry.Value,
		}
		if err := valueEntry.Decode(&secrets.Properties); err != nil {
			return err
		}
		node.Value = append(node.Value, secrets)
	}
	return nil
}

type JobSecretsValue struct {
	ID         string `yaml:"-"`
	Properties *SecretsPropertiesNode
}

type SecretsPropertiesNode struct {
	Raw   *yaml.Node
	Value *string
}

func (node *SecretsPropertiesNode) UnmarshalYAML(value *yaml.Node) error {
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

	return node.Raw.Decode(&node.Value)
}

type JobRunsOnNode struct {
	Raw   *yaml.Node
	OneOf JobRunsOnOneOf
}

func (node *JobRunsOnNode) UnmarshalYAML(value *yaml.Node) error {
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
	}

	return nil
}

type JobRunsOnOneOf struct {
	ScalarNode   *JobRunsOnConstants
	SequenceNode []*string
}

type JobRunsOnConstants string

const (
	JobRunsOn_Macos1015     JobRunsOnConstants = "macos-10.15"
	JobRunsOn_Macos11       JobRunsOnConstants = "macos-11"
	JobRunsOn_MacosLatest   JobRunsOnConstants = "macos-latest"
	JobRunsOn_SelfHosted    JobRunsOnConstants = "self-hosted"
	JobRunsOn_Ubuntu1804    JobRunsOnConstants = "ubuntu-18.04"
	JobRunsOn_Ubuntu2004    JobRunsOnConstants = "ubuntu-20.04"
	JobRunsOn_UbuntuLatest  JobRunsOnConstants = "ubuntu-latest"
	JobRunsOn_Windows2016   JobRunsOnConstants = "windows-2016"
	JobRunsOn_Windows2019   JobRunsOnConstants = "windows-2019"
	JobRunsOn_Windows2022   JobRunsOnConstants = "windows-2022"
	JobRunsOn_WindowsLatest JobRunsOnConstants = "windows-latest"
)

var JobRunsOn_Constants = []JobRunsOnConstants{
	JobRunsOn_Macos1015,
	JobRunsOn_Macos11,
	JobRunsOn_MacosLatest,
	JobRunsOn_SelfHosted,
	JobRunsOn_Ubuntu1804,
	JobRunsOn_Ubuntu2004,
	JobRunsOn_UbuntuLatest,
	JobRunsOn_Windows2016,
	JobRunsOn_Windows2019,
	JobRunsOn_Windows2022,
	JobRunsOn_WindowsLatest,
}

type JobRunsOnMaps map[string][]string

var (
	JobRunsOnMaps_SelfHosted JobRunsOnMaps = map[string][]string{
		"self-hosted": {
			"self-hosted",
		},
	}

	JobRunsOnMaps_Machine JobRunsOnMaps = map[string][]string{
		"machine": {
			"linux",
			"macos",
			"windows",
		},
	}

	JobRunsOnMaps_Architecture JobRunsOnMaps = map[string][]string{
		"architecture": {
			"ARM32",
			"x64",
			"x86",
		},
	}
)

type JobEnvironmentNode struct {
	Raw   *yaml.Node
	OneOf JobEnvironmentOneOf
}

type JobEnvironmentOneOf struct {
	ScalarNode  *string
	MappingNode *JobEnvironmentValue
}

func (node *JobEnvironmentNode) UnmarshalYAML(value *yaml.Node) error {
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
	case yaml.MappingNode:
		if len(node.Raw.Content)%2 != 0 {
			return fmt.Errorf("%d:%d  error  expected even number of key value pairs", node.Raw.Line, node.Raw.Column)
		}
		event := new(JobEnvironmentValue)
		for i := 0; i < len(value.Content); i += 2 {
			keyEntry := value.Content[i]
			valueEntry := value.Content[i+1]
			eventKey := keyEntry.Value
			switch eventKey {
			case "name":
				event.Name = new(EnvironmentNameNode)
				err := valueEntry.Decode(event.Name)
				if err != nil {
					return err
				}
			case "url":
				event.Url = new(EnvironmentUrlNode)
				err := valueEntry.Decode(event.Url)
				if err != nil {
					return err
				}
			}
		}
		node.OneOf.MappingNode = event
	}
	return nil
}

type JobEnvironmentValue struct {
	Name *EnvironmentNameNode `yaml:"node"`
	Url  *EnvironmentUrlNode  `yaml:"url"`
}

type EnvironmentNameNode struct {
	Raw   *yaml.Node
	Value *string
}

func (node *EnvironmentNameNode) UnmarshalYAML(value *yaml.Node) error {
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

type EnvironmentUrlNode struct {
	Raw   *yaml.Node
	Value *string
}

func (node *EnvironmentUrlNode) UnmarshalYAML(value *yaml.Node) error {
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

type JobOutputsNode struct {
	Raw   *yaml.Node
	Value []*JobOutputsValue
}

func (node *JobOutputsNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	if len(value.Content)%2 != 0 {
		return fmt.Errorf("%d:%d  error  Expected even number of key value pairs", node.Raw.Line, node.Raw.Column)
	}
	for i := 0; i < len(value.Content); i += 2 {
		keyEntry := value.Content[i]
		valueEntry := value.Content[i+1]
		outputs := &JobOutputsValue{
			ID: keyEntry.Value,
		}
		if err := valueEntry.Decode(&outputs.Properties); err != nil {
			return err
		}
		node.Value = append(node.Value, outputs)
	}
	return nil
}

type JobOutputsValue struct {
	ID         string `yaml:"-"`
	Properties *OutputsPropertiesNode
}

type OutputsPropertiesNode struct {
	Raw   *yaml.Node
	Value *string
}

func (node *OutputsPropertiesNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	scalarTypes := []string{"!!str"}
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

type JobEnvNode struct {
	Raw   *yaml.Node
	Value []*JobEnvValue
}

func (node *JobEnvNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	if len(value.Content)%2 != 0 {
		return fmt.Errorf("%d:%d  error  Expected even number of key value pairs", node.Raw.Line, node.Raw.Column)
	}
	for i := 0; i < len(value.Content); i += 2 {
		keyEntry := value.Content[i]
		valueEntry := value.Content[i+1]
		env := &JobEnvValue{
			ID: keyEntry.Value,
		}
		if err := valueEntry.Decode(&env.Properties); err != nil {
			return err
		}
		node.Value = append(node.Value, env)
	}
	return nil
}

type JobEnvValue struct {
	ID         string `yaml:"-"`
	Properties *JobEnvPropertiesNode
}

type JobEnvPropertiesNode struct {
	Raw   *yaml.Node
	Value *string
}

func (node *JobEnvPropertiesNode) UnmarshalYAML(value *yaml.Node) error {
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

type JobDefaultsNode struct {
	Raw   *yaml.Node
	Value *JobDefaultsValue
}

func (node *JobDefaultsNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value

	if len(value.Content)%2 != 0 {
		return fmt.Errorf("%d:%d  error  expected even number of key value pairs", node.Raw.Line, node.Raw.Column)
	}
	event := new(JobDefaultsValue)
	for i := 0; i < len(value.Content); i += 2 {
		keyEntry := value.Content[i]
		valueEntry := value.Content[i+1]
		eventKey := keyEntry.Value
		switch eventKey {
		case "run":
			event.Run = new(JobDefaultsRunNode)
			err := valueEntry.Decode(event.Run)
			if err != nil {
				return err
			}
		}
	}
	node.Value = event
	return value.Decode(&node.Value)
}

type JobDefaultsValue struct {
	Run *JobDefaultsRunNode `yaml:"run"`
}

type JobDefaultsRunNode struct {
	Raw   *yaml.Node
	Value JobDefaultsRunValue
}

func (node *JobDefaultsRunNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type JobDefaultsRunValue struct {
	Shell            *JobRunShellNode            `yaml:"shell"`
	WorkingDirectory *JobRunWorkingDirectoryNode `yaml:"working-directory"`
}

type JobRunShellNode struct {
	Raw   *yaml.Node
	Value RunShellConstants
}

func (node *JobRunShellNode) UnmarshalYAML(value *yaml.Node) error {
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
			return fmt.Errorf("%d:%d  error  %s %s", node.Raw.Line, node.Raw.Column, "expected one of scalar types:", strings.Join(scalarTypes, ", "))
		}
		return value.Decode(&node.Value)
	default:
		return fmt.Errorf("%d:%d  error  Expected any of: string type", node.Raw.Line, node.Raw.Column)
	}
}

type JobRunWorkingDirectoryNode struct {
	Raw   *yaml.Node
	Value *string
}

func (node *JobRunWorkingDirectoryNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	scalarTypes := []string{"!!str"}
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

type JobStepsNode struct {
	Raw   *yaml.Node
	Value []*JobStepsValue
}

func (node *JobStepsNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	if len(value.Content)%2 != 0 {
		return fmt.Errorf("%d:%d  error  expected even number of key value pairs", node.Raw.Line, node.Raw.Column)
	}
	for _, stepsContent := range node.Raw.Content {
		if len(stepsContent.Content)%2 != 0 {
			return fmt.Errorf("%d:%d  error  expected even number of key value pairs", node.Raw.Line, node.Raw.Column)
		}
		event := new(JobStepsValue)
		for i := 0; i < len(stepsContent.Content); i += 2 {
			keyEntry := stepsContent.Content[i]
			valueEntry := stepsContent.Content[i+1]
			eventKey := keyEntry.Value
			switch eventKey {
			case "id":
				event.Id = new(StepsIdNode)
				err := valueEntry.Decode(event.Id)
				if err != nil {
					return err
				}
			case "if":
				event.If = new(StepsIfNode)
				err := valueEntry.Decode(event.If)
				if err != nil {
					return err
				}
			case "name":
				event.Name = new(StepsNameNode)
				err := valueEntry.Decode(event.Name)
				if err != nil {
					return err
				}
			case "uses":
				event.Uses = new(StepsUsesNode)
				err := valueEntry.Decode(event.Uses)
				if err != nil {
					return err
				}
			case "run":
				event.Run = new(StepsRunNode)
				err := valueEntry.Decode(event.Run)
				if err != nil {
					return err
				}
			case "working-directory":
				event.WorkingDirectory = new(StepsWorkingDirectoryNode)
				err := valueEntry.Decode(event.WorkingDirectory)
				if err != nil {
					return err
				}
			case "shell":
				event.Shell = new(StepsShellNode)
				err := valueEntry.Decode(event.Shell)
				if err != nil {
					return err
				}
			case "with":
				event.With = new(StepsWithNode)
				err := valueEntry.Decode(event.With)
				if err != nil {
					return err
				}
			case "env":
				event.Env = new(StepsEnvNode)
				err := valueEntry.Decode(event.Env)
				if err != nil {
					return err
				}
			case "continue-on-error":
				event.ContinueOnError = new(StepsContinueOnErrorNode)
				err := valueEntry.Decode(event.ContinueOnError)
				if err != nil {
					return err
				}
			case "timeout-minutes":
				event.TimeoutMinutes = new(StepsTimeoutMinutesNode)
				err := valueEntry.Decode(event.TimeoutMinutes)
				if err != nil {
					return err
				}
			}
		}
		node.Value = append(node.Value, event)
	}
	return nil
}

type JobStepsValue struct {
	Id               *StepsIdNode               `yaml:"id"`
	If               *StepsIfNode               `yaml:"if"`
	Name             *StepsNameNode             `yaml:"name"`
	Uses             *StepsUsesNode             `yaml:"uses"`
	Run              *StepsRunNode              `yaml:"run"`
	WorkingDirectory *StepsWorkingDirectoryNode `yaml:"working-directory"`
	Shell            *StepsShellNode            `yaml:"shell"`
	With             *StepsWithNode             `yaml:"with"`
	Env              *StepsEnvNode              `yaml"env"`
	ContinueOnError  *StepsContinueOnErrorNode  `yaml:"continue-on-error"`
	TimeoutMinutes   *StepsTimeoutMinutesNode   `yaml:"timeout-minutes"`
}

type StepsIdNode struct {
	Raw   *yaml.Node
	Value *string
}

func (node *StepsIdNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	scalarTypes := []string{"!!str"}
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

type StepsNameNode struct {
	Raw   *yaml.Node
	Value *string
}

func (node *StepsNameNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	scalarTypes := []string{"!!str"}
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

type StepsIfNode struct {
	Raw   *yaml.Node
	Value *string
}

func (node *StepsIfNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	scalarTypes := []string{"!!str"}
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

type StepsUsesNode struct {
	Raw   *yaml.Node
	Value *string
}

func (node *StepsUsesNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	scalarTypes := []string{"!!str"}
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

type StepsRunNode struct {
	Raw   *yaml.Node
	Value *string
}

func (node *StepsRunNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	scalarTypes := []string{"!!str"}
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

type StepsWorkingDirectoryNode struct {
	Raw   *yaml.Node
	Value *string
}

func (node *StepsWorkingDirectoryNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	scalarTypes := []string{"!!str"}
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

type StepsShellNode struct {
	Raw   *yaml.Node
	Value *RunShellConstants
}

func (node *StepsShellNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	scalarTypes := []string{"!!str"}
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

type StepsWithNode struct {
	Raw   *yaml.Node
	Value []*StepsWithValue
}

func (node *StepsWithNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value

	if len(value.Content)%2 != 0 {
		return fmt.Errorf("%d:%d  error  Expected even number of key value pairs", node.Raw.Line, node.Raw.Column)
	}
	for i := 0; i < len(value.Content); i += 2 {
		keyEntry := value.Content[i]
		valueEntry := value.Content[i+1]
		with := &StepsWithValue{
			ID: keyEntry.Value,
		}
		if err := valueEntry.Decode(&with.Properties); err != nil {
			return err
		}
		node.Value = append(node.Value, with)
	}
	return nil
}

type StepsWithValue struct {
	ID         string `yaml:"-"`
	Properties *StepsWithPropertiesNode
}

type StepsWithPropertiesNode struct {
	Raw   *yaml.Node
	Value *string
}

func (node *StepsWithPropertiesNode) UnmarshalYAML(value *yaml.Node) error {
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

type StepsEnvNode struct {
	Raw   *yaml.Node
	Value *string
}

func (node *StepsEnvNode) UnmarshalYAML(value *yaml.Node) error {
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

type StepsContinueOnErrorNode struct {
	Raw   *yaml.Node
	Value *string
}

func (node *StepsContinueOnErrorNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	scalarTypes := []string{"!!str", "!!bool"}
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

type StepsTimeoutMinutesNode struct {
	Raw   *yaml.Node
	Value *string
}

func (node *StepsTimeoutMinutesNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	scalarTypes := []string{"!!int"}
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

type JobTimeoutMinutesNode struct {
	Raw   *yaml.Node
	Value *string
}

func (node *JobTimeoutMinutesNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	scalarTypes := []string{"!!int"}
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

type JobContinueOnErrorNode struct {
	Raw   *yaml.Node
	Value *string
}

func (node *JobContinueOnErrorNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	scalarTypes := []string{"!!bool", "!!str"}
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

type JobContainerNode struct {
	Raw   *yaml.Node
	OneOf JobContainerOneOf
}

type JobContainerOneOf struct {
	ScalarNode  *string
	MappingNode *JobContainerValue
}

func (node *JobContainerNode) UnmarshalYAML(value *yaml.Node) error {
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
	case yaml.MappingNode:
		if len(value.Content)%2 != 0 {
			return fmt.Errorf("%d:%d  error  expected even number of key value pairs", node.Raw.Line, node.Raw.Column)
		}
		event := new(JobContainerValue)
		for i := 0; i < len(value.Content); i += 2 {
			keyEntry := value.Content[i]
			valueEntry := value.Content[i+1]
			eventKey := keyEntry.Value
			switch eventKey {
			case "image":
				event.Image = new(ContainerImageNode)
				err := valueEntry.Decode(event.Image)
				if err != nil {
					return err
				}
			case "credentials":
				event.Credentials = new(ContainerCredentialsNode)
				err := valueEntry.Decode(event.Credentials)
				if err != nil {
					return err
				}
			case "env":
				event.Env = new(ContainerEnvNode)
				err := valueEntry.Decode(event.Env)
				if err != nil {
					return err
				}
			case "ports":
				event.Ports = new(ContainerPortsNode)
				err := valueEntry.Decode(event.Ports)
				if err != nil {
					return err
				}
			case "volumes":
				event.Volumes = new(ContainerVolumesNode)
				err := valueEntry.Decode(event.Volumes)
				if err != nil {
					return err
				}
			case "options":
				event.Options = new(ContainerOptionsNode)
				err := valueEntry.Decode(event.Options)
				if err != nil {
					return err
				}
			}
		}
		node.OneOf.MappingNode = event
	}
	return nil
}

type JobContainerValue struct {
	Image       *ContainerImageNode       `yaml:"image"`
	Credentials *ContainerCredentialsNode `yaml:"credentials"`
	Env         *ContainerEnvNode         `yaml:"env"`
	Ports       *ContainerPortsNode       `yaml:"ports"`
	Volumes     *ContainerVolumesNode     `yaml:"volumes"`
	Options     *ContainerOptionsNode     `yaml:"options"`
}

type ContainerImageNode struct {
	Raw   *yaml.Node
	Value *string
}

func (node *ContainerImageNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	scalarTypes := []string{"!!str"}
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

type ContainerCredentialsNode struct {
	Raw   *yaml.Node
	Value *ContainerCredentialsValue
}

func (node *ContainerCredentialsNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value

	if len(value.Content)%2 != 0 {
		return fmt.Errorf("%d:%d  error  expected even number of key value pairs", node.Raw.Line, node.Raw.Column)
	}
	event := new(ContainerCredentialsValue)
	for i := 0; i < len(value.Content); i += 2 {
		keyEntry := value.Content[i]
		valueEntry := value.Content[i+1]
		eventKey := keyEntry.Value
		switch eventKey {
		case "username":
			event.Username = new(CredentialsUsernameNode)
			err := valueEntry.Decode(event.Username)
			if err != nil {
				return err
			}
		case "password":
			event.Password = new(CredentialsPasswordNode)
			err := valueEntry.Decode(event.Password)
			if err != nil {
				return err
			}
		}
	}
	node.Value = event

	return nil
}

type ContainerCredentialsValue struct {
	Username *CredentialsUsernameNode `yaml:"username"`
	Password *CredentialsPasswordNode `yaml:"password"`
}

type CredentialsUsernameNode struct {
	Raw   *yaml.Node
	Value *string
}

func (node *CredentialsUsernameNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	scalarTypes := []string{"!!str"}
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

type CredentialsPasswordNode struct {
	Raw   *yaml.Node
	Value *string
}

func (node *CredentialsPasswordNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	scalarTypes := []string{"!!str"}
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

type ContainerEnvNode struct {
	Raw   *yaml.Node
	Value []*ContainerEnvValue
}

func (node *ContainerEnvNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value

	if len(value.Content)%2 != 0 {
		return fmt.Errorf("%d:%d  error  Expected even number of key value pairs", node.Raw.Line, node.Raw.Column)
	}
	for i := 0; i < len(value.Content); i += 2 {
		keyEntry := value.Content[i]
		valueEntry := value.Content[i+1]
		env := &ContainerEnvValue{
			ID: keyEntry.Value,
		}
		if err := valueEntry.Decode(&env.Properties); err != nil {
			return err
		}
		node.Value = append(node.Value, env)
	}
	return nil
}

type ContainerEnvValue struct {
	ID         string `yaml:"-"`
	Properties *EnvPropertiesNode
}

type ContainerEnvPropertiesNode struct {
	Raw   *yaml.Node
	Value *string
}

func (node *ContainerEnvPropertiesNode) UnmarshalYAML(value *yaml.Node) error {
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

type ContainerPortsNode struct {
	Raw   *yaml.Node
	Value []*string
}

func (node *ContainerPortsNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	for _, content := range node.Raw.Content {
		scalarTypes := []string{"!!str", "!!int"}
		contains := false
		for _, scalarType := range scalarTypes {
			if content.Tag == scalarType {
				contains = true
			}
		}
		if !contains {
			return fmt.Errorf("%d:%d  error  %s %s", node.Raw.Line, node.Raw.Column, "expected one of scalar types:", strings.Join(scalarTypes, ", "))
		}
	}

	return value.Decode(&node.Value)
}

type ContainerVolumesNode struct {
	Raw   *yaml.Node
	Value []*string
}

func (node *ContainerVolumesNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	for _, content := range node.Raw.Content {
		scalarTypes := []string{"!!str"}
		contains := false
		for _, scalarType := range scalarTypes {
			if content.Tag == scalarType {
				contains = true
			}
		}
		if !contains {
			return fmt.Errorf("%d:%d  error  %s %s", node.Raw.Line, node.Raw.Column, "expected one of scalar types:", strings.Join(scalarTypes, ", "))
		}
	}

	return value.Decode(&node.Value)
}

type ContainerOptionsNode struct {
	Raw   *yaml.Node
	Value *string
}

func (node *ContainerOptionsNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	scalarTypes := []string{"!!str"}
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

type JobServicesNode struct {
	Raw   *yaml.Node
	Value []*JobServicesValue
}

func (node *JobServicesNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	if len(value.Content)%2 != 0 {
		return fmt.Errorf("%d:%d  error  Expected even number of key value pairs", node.Raw.Line, node.Raw.Column)
	}
	for i := 0; i < len(value.Content); i += 2 {
		keyEntry := value.Content[i]
		valueEntry := value.Content[i+1]
		services := &JobServicesValue{
			ID: keyEntry.Value,
		}
		if err := valueEntry.Decode(&services.Properties); err != nil {
			return err
		}
		node.Value = append(node.Value, services)
	}
	return nil
}

type JobServicesValue struct {
	ID         string `yaml:"-"`
	Properties *JobServicesPropertiesNode
}

type JobServicesPropertiesNode struct {
	Raw   *yaml.Node
	Value *JobContainerValue
}

func (node *JobServicesPropertiesNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	if len(value.Content)%2 != 0 {
		return fmt.Errorf("%d:%d  error  expected even number of key value pairs", node.Raw.Line, node.Raw.Column)
	}
	event := new(JobContainerValue)
	for i := 0; i < len(value.Content); i += 2 {
		keyEntry := value.Content[i]
		valueEntry := value.Content[i+1]
		eventKey := keyEntry.Value
		switch eventKey {
		case "image":
			event.Image = new(ContainerImageNode)
			err := valueEntry.Decode(event.Image)
			if err != nil {
				return err
			}
		case "credentials":
			event.Credentials = new(ContainerCredentialsNode)
			err := valueEntry.Decode(event.Credentials)
			if err != nil {
				return err
			}
		case "env":
			event.Env = new(ContainerEnvNode)
			err := valueEntry.Decode(event.Env)
			if err != nil {
				return err
			}
		case "ports":
			event.Ports = new(ContainerPortsNode)
			err := valueEntry.Decode(event.Ports)
			if err != nil {
				return err
			}
		case "volumes":
			event.Volumes = new(ContainerVolumesNode)
			err := valueEntry.Decode(event.Volumes)
			if err != nil {
				return err
			}
		case "options":
			event.Options = new(ContainerOptionsNode)
			err := valueEntry.Decode(event.Options)
			if err != nil {
				return err
			}
		}
	}
	node.Value = event
	return nil
}

type JobConcurrencyNode struct {
	Raw   *yaml.Node
	OneOf JobConcurrencyOneOf
}

type JobConcurrencyOneOf struct {
	ScalarNode  string
	MappingNode *WorkflowConcurrencyValue
}

func (node *JobConcurrencyNode) UnmarshalYAML(value *yaml.Node) error {
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
			return fmt.Errorf("%d:%d  error  %s %s", node.Raw.Line, node.Raw.Column, "expected one of scalar types:", strings.Join(scalarTypes, ", "))
		}
		return value.Decode(&node.OneOf.ScalarNode)
	case yaml.MappingNode:
		if len(value.Content)%2 != 0 {
			return fmt.Errorf("%d:%d  error  expected even number of key value pairs", node.Raw.Line, node.Raw.Column)
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

// --------------------------------------------JOBS----------------------------------------------------

// --------------------------------------------ENV----------------------------------------------------
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

// --------------------------------------------ENV----------------------------------------------------

// --------------------------------------------Permissions----------------------------------------------------

type DefinitionPermissionsNode struct {
	Raw   *yaml.Node
	OneOf DefinitionPermissionsOneOf
}

type DefinitionPermissionsOneOf struct {
	ScalarNode  *string
	MappingNode *DefinitionPermissionsValue
}

func (node *DefinitionPermissionsNode) UnmarshalYAML(value *yaml.Node) error {
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
			return fmt.Errorf("%d:%d  error  %s %s", node.Raw.Line, node.Raw.Column, "expected one of scalar types:", strings.Join(scalarTypes, ", "))
		}
	
		return value.Decode(&node.OneOf.ScalarNode)
	case yaml.MappingNode:
		if len(value.Content)%2 != 0 {
			return fmt.Errorf("%d:%d  error  expected even number of key value pairs", node.Raw.Line, node.Raw.Column)
		}
		event := new(DefinitionPermissionsValue)
		for i := 0; i < len(value.Content); i += 2 {
			keyEntry := value.Content[i]
			valueEntry := value.Content[i+1]
			eventKey := keyEntry.Value
			switch eventKey {
			case "actions":
				event.Actions = new(PermissionsActionsNode)
				err := valueEntry.Decode(event.Actions)
				if err != nil {
					return err
				}
			case "checks":
				event.Checks = new(PermissionsChecksNode)
				err := valueEntry.Decode(event.Checks)
				if err != nil {
					return err
				}
			case "contents":
				event.Contents = new(PermissionsContentsNode)
				err := valueEntry.Decode(event.Contents)
				if err != nil {
					return err
				}
			case "deployments":
				event.Deployments = new(PermissionsDeploymentsNode)
				err := valueEntry.Decode(event.Deployments)
				if err != nil {
					return err
				}
			case "issues":
				event.Issues = new(PermissionsIssuesNode)
				err := valueEntry.Decode(event.Issues)
				if err != nil {
					return err
				}
			case "packages":
				event.Packages = new(PermissionsPackagesNode)
				err := valueEntry.Decode(event.Packages)
				if err != nil {
					return err
				}
			case "pull-requests":
				event.PullRequests = new(PermissionsPullRequestsNode)
				err := valueEntry.Decode(event.PullRequests)
				if err != nil {
					return err
				}
			case "repository-projects":
				event.RepositoryProjects = new(PermissionsRepositoryProjectsNode)
				err := valueEntry.Decode(event.RepositoryProjects)
				if err != nil {
					return err
				}
			case "security-events":
				event.SecurityEvents = new(PermissionsSecurityEventsNode)
				err := valueEntry.Decode(event.SecurityEvents)
				if err != nil {
					return err
				}
			case "statuses":
				event.Statuses = new(PermissionsStatusesNode)
				err := valueEntry.Decode(event.Statuses)
				if err != nil {
					return err
				}
			case "id-token":
				event.IdToken = new(PermissionsIdTokenNode)
				err := valueEntry.Decode(event.IdToken)
				if err != nil {
					return err
				}
			}
		}
		node.OneOf.MappingNode = event
		return nil
	}

	return nil
}

type DefinitionPermissionsValue struct {
	Actions            *PermissionsActionsNode            `yaml:"actions"`
	Checks             *PermissionsChecksNode             `yaml:"checks"`
	Contents           *PermissionsContentsNode           `yaml:"contents"`
	Deployments        *PermissionsDeploymentsNode        `yaml:"deployments"`
	Issues             *PermissionsIssuesNode             `yaml:"issues"`
	Packages           *PermissionsPackagesNode           `yaml:"packages"`
	PullRequests       *PermissionsPullRequestsNode       `yaml:"pull-requests"`
	RepositoryProjects *PermissionsRepositoryProjectsNode `yaml:"repository-projects"`
	SecurityEvents     *PermissionsSecurityEventsNode     `yaml:"security-events"`
	Statuses           *PermissionsStatusesNode           `yaml:"statuses"`
	IdToken            *PermissionsIdTokenNode            `yaml:"id-token"`
}

type PermissionsActionsNode struct {
	Raw   *yaml.Node
	Value *JobPermissionsConstants
}

func (node *PermissionsActionsNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type PermissionsChecksNode struct {
	Raw   *yaml.Node
	Value *JobPermissionsConstants
}

func (node *PermissionsChecksNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type PermissionsContentsNode struct {
	Raw   *yaml.Node
	Value *JobPermissionsConstants
}

func (node *PermissionsContentsNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type PermissionsDeploymentsNode struct {
	Raw   *yaml.Node
	Value *JobPermissionsConstants
}

func (node *PermissionsDeploymentsNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type PermissionsIssuesNode struct {
	Raw   *yaml.Node
	Value *JobPermissionsConstants
}

func (node *PermissionsIssuesNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type PermissionsPackagesNode struct {
	Raw   *yaml.Node
	Value *JobPermissionsConstants
}

func (node *PermissionsPackagesNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type PermissionsPullRequestsNode struct {
	Raw   *yaml.Node
	Value *JobPermissionsConstants
}

func (node *PermissionsPullRequestsNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type PermissionsRepositoryProjectsNode struct {
	Raw   *yaml.Node
	Value *JobPermissionsConstants
}

func (node *PermissionsRepositoryProjectsNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type PermissionsSecurityEventsNode struct {
	Raw   *yaml.Node
	Value *JobPermissionsConstants
}

func (node *PermissionsSecurityEventsNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type PermissionsStatusesNode struct {
	Raw   *yaml.Node
	Value *JobPermissionsConstants
}

func (node *PermissionsStatusesNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type PermissionsIdTokenNode struct {
	Raw   *yaml.Node
	Value *JobPermissionsConstants
}

func (node *PermissionsIdTokenNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type JobPermissionsConstants string

const (
	JobPermissions_Read  JobPermissionsConstants = "read"
	JobPermissions_Write JobPermissionsConstants = "write"
	JobPermissions_None  JobPermissionsConstants = "none"
)

var JobPermissions_Constants = []JobPermissionsConstants{
	JobPermissions_Read,
	JobPermissions_Write,
	JobPermissions_None,
}

// --------------------------------------------Permissions----------------------------------------------------
