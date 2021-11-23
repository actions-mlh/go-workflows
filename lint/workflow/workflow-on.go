package workflow

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

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
		return value.Decode(&node.OneOf.ScalarNode)
	case yaml.SequenceNode:
		return value.Decode(&node.OneOf.SequenceNode)
	case yaml.MappingNode:
		value := node.Raw
		if len(value.Content)%2 != 0 {
			return fmt.Errorf("%d:%d\terror\tCould not process on: value.Contents has odd length, should be paired", node.Raw.Line, node.Raw.Column)
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
			case "delete":
				event.Delete = new(OnDeleteNode)
				err := valueEntry.Decode(event.Delete)
				if err != nil {
					return err
				}
			case "deployment":
				event.Deployment = new(OnDeploymentNode)
				err := valueEntry.Decode(event.Deployment)
				if err != nil {
					return err
				}
			case "deployment_status":
				event.DeploymentStatus = new(OnDeploymentStatusNode)
				err := valueEntry.Decode(event.DeploymentStatus)
				if err != nil {
					return err
				}
			case "discussion":
				event.Discussion = new(OnDiscussionNode)
				err := valueEntry.Decode(event.Discussion)
				if err != nil {
					return err
				}
			case "discussion_comment":
				event.DiscussionComment = new(OnDiscussionCommentNode)
				err := valueEntry.Decode(event.DiscussionComment)
				if err != nil {
					return err
				}
			case "fork":
				event.Fork = new(OnForkNode)
				err := valueEntry.Decode(event.Fork)
				if err != nil {
					return err
				}
			case "gollum":
				event.Gollum = new(OnGollumNode)
				err := valueEntry.Decode(event.Gollum)
				if err != nil {
					return err
				}
			case "issue_comment":
				event.IssueComment = new(OnIssueCommentNode)
				err := valueEntry.Decode(event.IssueComment)
				if err != nil {
					return err
				}
			case "issues":
				event.Issues = new(OnIssuesNode)
				err := valueEntry.Decode(event.Issues)
				if err != nil {
					return err
				}
			case "label":
				event.Label = new(OnLabelNode)
				err := valueEntry.Decode(event.Label)
				if err != nil {
					return err
				}
			case "milestone":
				event.Milestone = new(OnMilestoneNode)
				err := valueEntry.Decode(event.Milestone)
				if err != nil {
					return err
				}
			case "page_build":
				event.PageBuild = new(OnPageBuildNode)
				err := valueEntry.Decode(event.PageBuild)
				if err != nil {
					return err
				}
			case "project":
				event.Project = new(OnProjectNode)
				err := valueEntry.Decode(event.Project)
				if err != nil {
					return err
				}
			case "project_card":
				event.ProjectCard = new(OnProjectCardNode)
				err := valueEntry.Decode(event.ProjectCard)
				if err != nil {
					return err
				}
			case "project_column":
				event.ProjectColumn = new(OnProjectColumnNode)
				err := valueEntry.Decode(event.ProjectColumn)
				if err != nil {
					return err
				}
			case "public":
				event.Public = new(OnPublicNode)
				err := valueEntry.Decode(event.Public)
				if err != nil {
					return err
				}
			case "pull_request":
				event.PullRequest = new(OnPullRequestNode)
				err := valueEntry.Decode(event.PullRequest)
				if err != nil {
					return err
				}
			case "pull_request_review":
				event.PullRequestReview = new(OnPullRequestReviewNode)
				err := valueEntry.Decode(event.PullRequestReview)
				if err != nil {
					return err
				}
			case "pull_request_review_comment":
				event.PullRequestReviewComment = new(OnPullRequestReviewCommentNode)
				err := valueEntry.Decode(event.PullRequestReviewComment)
				if err != nil {
					return err
				}
			case "pull_request_target":
				event.PullRequestTarget = new(OnPullRequestTargetNode)
				err := valueEntry.Decode(event.PullRequestTarget)
				if err != nil {
					return err
				}
			case "push":
				event.Push = new(OnPushNode)
				err := valueEntry.Decode(event.Push)
				if err != nil {
					return err
				}
			case "registry_package":
				event.RegistryPackage = new(OnRegistryPackageNode)
				err := valueEntry.Decode(event.RegistryPackage)
				if err != nil {
					return err
				}
			case "release":
				event.Release = new(OnReleaseNode)
				err := valueEntry.Decode(event.Release)
				if err != nil {
					return err
				}
			case "status":
				event.Status = new(OnStatusNode)
				err := valueEntry.Decode(event.Status)
				if err != nil {
					return err
				}
			case "watch":
				event.Watch = new(OnWatchNode)
				err := valueEntry.Decode(event.Watch)
				if err != nil {
					return err
				}
			case "workflow_run":
				event.WorkflowRun = new(OnWorkflowRunNode)
				err := valueEntry.Decode(event.WorkflowRun)
				if err != nil {
					return err
				}
			case "schedule":
				event.Schedule = new(OnScheduleNode)
				err := valueEntry.Decode(event.Schedule)
				if err != nil {
					return err
				}
			case "workflow_dispatch":
				event.WorkflowDispatch = new(OnWorkflowDispatchNode)
				err := valueEntry.Decode(event.WorkflowDispatch)
				if err != nil {
					return err
				}
			case "repository_dispatch":
				event.RepositoryDispatch = new(OnRepositoryDispatchNode)
				err := valueEntry.Decode(event.RepositoryDispatch)
				if err != nil {
					return err
				}
			case "workflow_call":
				event.WorkflowCall = new(OnWorkflowCallNode)
				err := valueEntry.Decode(event.WorkflowCall)
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
	CheckRun                 *OnCheckRunNode                 `yaml:"check_run"`
	CheckSuite               *OnCheckSuiteNode               `yaml:"check_suite"`
	Create                   *OnCreateNode                   `yaml:"create"`
	Delete                   *OnDeleteNode                   `yaml:"delete"`
	Deployment               *OnDeploymentNode               `yaml:"deployment"`
	DeploymentStatus         *OnDeploymentStatusNode         `yaml:"deployment_status"`
	Discussion               *OnDiscussionNode               `yaml:"discussion"`
	DiscussionComment        *OnDiscussionCommentNode        `yaml:"discussion_comment"`
	Fork                     *OnForkNode                     `yaml:"fork"`
	Gollum                   *OnGollumNode                   `yaml:"gollum"`
	IssueComment             *OnIssueCommentNode             `yaml:"issue_comment"`
	Issues                   *OnIssuesNode                   `yaml:"issues"`
	Label                    *OnLabelNode                    `yaml:"label"`
	Milestone                *OnMilestoneNode                `yaml:"milestone"`
	PageBuild                *OnPageBuildNode                `yaml:"page_build"`
	Project                  *OnProjectNode                  `yaml:"project"`
	ProjectCard              *OnProjectCardNode              `yaml:"project_card"`
	ProjectColumn            *OnProjectColumnNode            `yaml:"project_column"`
	Public                   *OnPublicNode                   `yaml:"public"`
	PullRequest              *OnPullRequestNode              `yaml:"pull_request"`
	PullRequestReview        *OnPullRequestReviewNode        `yaml:"pull_request_review"`
	PullRequestReviewComment *OnPullRequestReviewCommentNode `yaml:"pull_request_review_comment"`
	PullRequestTarget        *OnPullRequestTargetNode        `yaml:"pull_request_target"`
	Push                     *OnPushNode                     `yaml:"push"`
	RegistryPackage          *OnRegistryPackageNode          `yaml:"registry_package"`
	Release                  *OnReleaseNode                  `yaml:"release"`
	Status                   *OnStatusNode                   `yaml:"status"`
	Watch                    *OnWatchNode                    `yaml:"watch"`
	WorkflowRun              *OnWorkflowRunNode              `yaml:"workflow_run"`
	Schedule                 *OnScheduleNode                 `yaml:"schedule"`
	WorkflowDispatch         *OnWorkflowDispatchNode         `yaml:"workflow_dispatch"`
	RepositoryDispatch       *OnRepositoryDispatchNode       `yaml:"repository_dispatch"`
	WorkflowCall             *OnWorkflowCallNode             `yaml:"workflow_call"`
}

type OnCheckRunNode struct {
	Raw   *yaml.Node
	Value *DefinitionsTypeValue
}

func (node *OnCheckRunNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	if len(value.Content)%2 != 0 {
		return fmt.Errorf("%d:%d\terror\tCould not process onCheckRun: value.Contents has odd length, should be paired", node.Raw.Line, node.Raw.Column)
	}
	event := new(DefinitionsTypeValue)
	for i := 0; i < len(value.Content); i += 2 {
		keyEntry := value.Content[i]
		valueEntry := value.Content[i+1]

		eventKey := keyEntry.Value
		switch eventKey {
		case "types":
			event.Types = new(DefinitionsTypeNode)
			err := valueEntry.Decode(event.Types)
			if err != nil {
				return err
			}
		}
	}
	node.Value = event
	return nil
}

type OnCheckSuiteNode struct {
	Raw   *yaml.Node
	Value *DefinitionsTypeValue
}

func (node *OnCheckSuiteNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	if len(value.Content)%2 != 0 {
		return fmt.Errorf("%d:%d\terror\tCould not process onCheckSuite: value.Contents has odd length, should be paired", node.Raw.Line, node.Raw.Column)
	}
	event := new(DefinitionsTypeValue)
	for i := 0; i < len(value.Content); i += 2 {
		keyEntry := value.Content[i]
		valueEntry := value.Content[i+1]

		eventKey := keyEntry.Value
		switch eventKey {
		case "types":
			event.Types = new(DefinitionsTypeNode)
			err := valueEntry.Decode(event.Types)
			if err != nil {
				return err
			}
		}
	}
	node.Value = event
	return nil
}

type OnDiscussionNode struct {
	Raw   *yaml.Node
	Value *DefinitionsTypeValue
}

func (node *OnDiscussionNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	if len(value.Content)%2 != 0 {
		return fmt.Errorf("%d:%d\terror\tCould not process onDiscussion: value.Contents has odd length, should be paired", node.Raw.Line, node.Raw.Column)
	}
	event := new(DefinitionsTypeValue)
	for i := 0; i < len(value.Content); i += 2 {
		keyEntry := value.Content[i]
		valueEntry := value.Content[i+1]

		eventKey := keyEntry.Value
		switch eventKey {
		case "types":
			event.Types = new(DefinitionsTypeNode)
			err := valueEntry.Decode(event.Types)
			if err != nil {
				return err
			}
		}
	}
	node.Value = event
	return nil
}

type OnDiscussionCommentNode struct {
	Raw   *yaml.Node
	Value *DefinitionsTypeValue
}

func (node *OnDiscussionCommentNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	if len(value.Content)%2 != 0 {
		return fmt.Errorf("%d:%d\terror\tCould not process onDiscussionComment: value.Contents has odd length, should be paired", node.Raw.Line, node.Raw.Column)
	}
	event := new(DefinitionsTypeValue)
	for i := 0; i < len(value.Content); i += 2 {
		keyEntry := value.Content[i]
		valueEntry := value.Content[i+1]

		eventKey := keyEntry.Value
		switch eventKey {
		case "types":
			event.Types = new(DefinitionsTypeNode)
			err := valueEntry.Decode(event.Types)
			if err != nil {
				return err
			}
		}
	}
	node.Value = event
	return nil
}

type DefinitionsTypeValue struct {
	Types *DefinitionsTypeNode `yaml:"types"`
}

type DefinitionsTypeNode struct {
	Raw   *yaml.Node
	Value *[]DefinitionsTypesConstants
}

func (node *DefinitionsTypeNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type DefinitionsTypesConstants string

const (
	CheckRunTypes_Created         DefinitionsTypesConstants = "create"
	CheckRunTypes_Rerequested     DefinitionsTypesConstants = "rerequested"
	CheckRunTypes_Completed       DefinitionsTypesConstants = "completed"
	CheckRunTypes_RequestedAction DefinitionsTypesConstants = "requested_action"
)

var CheckRunTypes_Constants = []DefinitionsTypesConstants{
	CheckRunTypes_Created,
	CheckRunTypes_Rerequested,
	CheckRunTypes_Completed,
	CheckRunTypes_RequestedAction,
}

const (
	CheckSuiteTypes_Completed   DefinitionsTypesConstants = "completed"
	CheckSuiteTypes_Requested   DefinitionsTypesConstants = "requested"
	CheckSuiteTypes_Rerequested DefinitionsTypesConstants = "rerequested"
)

var CheckSuiteTypes_Constants = []DefinitionsTypesConstants{
	CheckSuiteTypes_Completed,
	CheckSuiteTypes_Requested,
	CheckSuiteTypes_Rerequested,
}

const (
	DiscussionTypes_Opened          DefinitionsTypesConstants = "opened"
	DiscussionTypes_Edited          DefinitionsTypesConstants = "edited"
	DiscussionTypes_Deleted         DefinitionsTypesConstants = "deleted"
	DiscussionTypes_Transferred     DefinitionsTypesConstants = "transferred"
	DiscussionTypes_Pinned          DefinitionsTypesConstants = "pinned"
	DiscussionTypes_Unpinned        DefinitionsTypesConstants = "unpinned"
	DiscussionTypes_Labeled         DefinitionsTypesConstants = "labeled"
	DiscussionTypes_Unlabeled       DefinitionsTypesConstants = "unlabeled"
	DiscussionTypes_Locked          DefinitionsTypesConstants = "locked"
	DiscussionTypes_Unlocked        DefinitionsTypesConstants = "unlocked"
	DiscussionTypes_CategoryChanged DefinitionsTypesConstants = "category_changed"
	DiscussionTypes_Answered        DefinitionsTypesConstants = "answered"
	DiscussionTypes_Unanswered      DefinitionsTypesConstants = "unanswered"
)

var DiscussionTypes_Constants = []DefinitionsTypesConstants{
	DiscussionTypes_Opened,
	DiscussionTypes_Edited,
	DiscussionTypes_Deleted,
	DiscussionTypes_Transferred,
	DiscussionTypes_Pinned,
	DiscussionTypes_Unpinned,
	DiscussionTypes_Labeled,
	DiscussionTypes_Unlabeled,
	DiscussionTypes_Locked,
	DiscussionTypes_Unlocked,
	DiscussionTypes_CategoryChanged,
	DiscussionTypes_Answered,
	DiscussionTypes_Unanswered,
}

const (
	DiscussionCommentTypes_Created DefinitionsTypesConstants = "created"
	DiscussionCommentTypes_Edited  DefinitionsTypesConstants = "edited"
	DiscussionCommentTypes_Deleted DefinitionsTypesConstants = "deleted"
)

var DiscussionCommentTypes_Constants = []DefinitionsTypesConstants{
	DiscussionCommentTypes_Created,
	DiscussionCommentTypes_Edited,
	DiscussionCommentTypes_Deleted,
}

type OnCreateNode struct {
	Raw   *yaml.Node
	Value *interface{}
}

func (node *OnCreateNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type OnDeleteNode struct {
	Raw   *yaml.Node
	Value *interface{}
}

func (node *OnDeleteNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type OnDeploymentNode struct {
	Raw   *yaml.Node
	Value *interface{}
}

func (node *OnDeploymentNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type OnDeploymentStatusNode struct {
	Raw   *yaml.Node
	Value *interface{}
}

func (node *OnDeploymentStatusNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type OnForkNode struct {
	Raw   *yaml.Node
	Value *interface{}
}

func (node *OnForkNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type OnGollumNode struct {
	Raw   *yaml.Node
	Value *interface{}
}

func (node *OnGollumNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type OnIssueCommentNode struct {
	Raw   *yaml.Node
	Value *interface{}
}

func (node *OnIssueCommentNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type OnIssuesNode struct {
	Raw   *yaml.Node
	Value *interface{}
}

func (node *OnIssuesNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type OnLabelNode struct {
	Raw   *yaml.Node
	Value *interface{}
}

func (node *OnLabelNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type OnMilestoneNode struct {
	Raw   *yaml.Node
	Value *interface{}
}

func (node *OnMilestoneNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type OnPageBuildNode struct {
	Raw   *yaml.Node
	Value *interface{}
}

func (node *OnPageBuildNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type OnProjectNode struct {
	Raw   *yaml.Node
	Value *interface{}
}

func (node *OnProjectNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type OnProjectCardNode struct {
	Raw   *yaml.Node
	Value *interface{}
}

func (node *OnProjectCardNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type OnProjectColumnNode struct {
	Raw   *yaml.Node
	Value *interface{}
}

func (node *OnProjectColumnNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type OnPublicNode struct {
	Raw   *yaml.Node
	Value *interface{}
}

func (node *OnPublicNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type OnPullRequestNode struct {
	Raw   *yaml.Node
	Value *interface{}
}

func (node *OnPullRequestNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type OnPullRequestReviewNode struct {
	Raw   *yaml.Node
	Value *interface{}
}

func (node *OnPullRequestReviewNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type OnPullRequestReviewCommentNode struct {
	Raw   *yaml.Node
	Value *interface{}
}

func (node *OnPullRequestReviewCommentNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type OnPullRequestTargetNode struct {
	Raw   *yaml.Node
	Value *interface{}
}

func (node *OnPullRequestTargetNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type OnPushNode struct {
	Raw   *yaml.Node
	Value *interface{}
}

func (node *OnPushNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type OnRegistryPackageNode struct {
	Raw   *yaml.Node
	Value *interface{}
}

func (node *OnRegistryPackageNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type OnReleaseNode struct {
	Raw   *yaml.Node
	Value *interface{}
}

func (node *OnReleaseNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type OnStatusNode struct {
	Raw   *yaml.Node
	Value *interface{}
}

func (node *OnStatusNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type OnWatchNode struct {
	Raw   *yaml.Node
	Value *interface{}
}

func (node *OnWatchNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type OnWorkflowRunNode struct {
	Raw   *yaml.Node
	Value *interface{}
}

func (node *OnWorkflowRunNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type OnScheduleNode struct {
	Raw   *yaml.Node
	Value *interface{}
}

func (node *OnScheduleNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type OnWorkflowDispatchNode struct {
	Raw   *yaml.Node
	Value *interface{}
}

func (node *OnWorkflowDispatchNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type OnRepositoryDispatchNode struct {
	Raw   *yaml.Node
	Value *interface{}
}

func (node *OnRepositoryDispatchNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type OnWorkflowCallNode struct {
	Raw   *yaml.Node
	Value *interface{}
}

func (node *OnWorkflowCallNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}
