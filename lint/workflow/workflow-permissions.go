package workflow

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

type WorkflowPermissionsNode struct {
	Raw   *yaml.Node
	OneOf WorkflowPermissionsOneOf
}

type WorkflowPermissionsOneOf struct {
	ScalarNode  *string
	MappingNode *DefinitionPermissionsValue
}

func (node *WorkflowPermissionsNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	switch node.Raw.Kind {
	case yaml.ScalarNode:
		// TYPE:string
		return value.Decode(&node.OneOf.ScalarNode)
	case yaml.MappingNode:
		if len(value.Content)%2 != 0 {
			return fmt.Errorf("%d:%d\terror\tCould not process permissions: value.Contents has odd length, should be paired", node.Raw.Line, node.Raw.Column)
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
