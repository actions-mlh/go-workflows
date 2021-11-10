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
	CheckRun          *OnCheckRunNode          `yaml:"check_run"`
	CheckSuite        *OnCheckSuiteNode        `yaml:"check_suite"`
	Create            *OnCreateNode            `yaml:"create"`
	Delete            *OnDeleteNode            `yaml:"delete"`
	Deployment        *OnDeploymentNode        `yaml:"deployment"`
	DeploymentStatus  *OnDeploymentStatusNode  `yaml:"deployment_status"`
	Discussion        *OnDiscussionNode        `yaml:"discussion"`
	DiscussionComment *OnDiscussionCommentNode `yaml:"discussion_comment"`
	Fork              *OnForkNode              `yaml:"fork"`
	Gollum            *OnGollumNode            `yaml:"gollum"`
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
	Value *string
}

func (node *OnCreateNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type OnDeleteNode struct {
	Raw   *yaml.Node
	Value *string
}

func (node *OnDeleteNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type OnDeploymentNode struct {
	Raw   *yaml.Node
	Value *string
}

func (node *OnDeploymentNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type OnDeploymentStatusNode struct {
	Raw   *yaml.Node
	Value *string
}

func (node *OnDeploymentStatusNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type OnForkNode struct {
	Raw   *yaml.Node
	Value *string
}

func (node *OnForkNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

type OnGollumNode struct {
	Raw   *yaml.Node
	Value *string
}

func (node *OnGollumNode) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}
