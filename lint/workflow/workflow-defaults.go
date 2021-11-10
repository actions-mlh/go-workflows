package workflow

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

type WorkflowDefaultsNode struct {
	Raw   *yaml.Node
	Value *WorkflowDefaultsValue
}

func (node *WorkflowDefaultsNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value

	if len(value.Content)%2 != 0 {
		return fmt.Errorf("%d:%d\terror\tCould not process defaults: value.Contents has odd length, should be paired", node.Raw.Line, node.Raw.Column)
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
	// TYPE:string
	return value.Decode(&node.Value)
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
