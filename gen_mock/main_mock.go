package mock_gen_schema

// import (
// 	"gopkg.in/yaml.v3"
// )

// type WorkflowNode struct {
// 	Raw   *yaml.Node
// 	Value WorkflowValue
// }

// type WorkflowValue struct {
// 	Name NameNode           `yaml:"name"`
// 	Jobs map[string]JobNode `yaml:"jobs"`
// 	// On       OnNode       `yaml:"on"`
// 	// Env      EnvNode      `yaml:"env"`
// 	// Defaults DefaultsNode `yaml:"defaults"`
// }

// func (node *WorkflowNode) UnmarshalYAML(value *yaml.Node) error {
// 	node.Raw = value
// 	return value.Decode(&node.Value)
// }

// type NameNode struct {
// 	Raw   *yaml.Node
// 	Value string
// }

// func (node *NameNode) UnmarshalYAML(value *yaml.Node) error {
// 	node.Raw = value
// 	return value.Decode(&node.Value)
// }

// type JobNode struct {
// 	Raw   *yaml.Node
// 	Value JobValue
// }

// func (node *JobNode) UnmarshalYAML(value *yaml.Node) error {
// 	node.Raw = value
// 	return value.Decode(&node.Value)
// }

// type JobValue struct {
// 	Name  string     `yaml:"name"`
// 	Steps []StepNode `yaml:"steps"`
// 	Uses  string     `yaml:"uses"`
// }

// type StepNode struct {
// 	Raw   *yaml.Node
// 	Value StepValue
// }

// type StepValue struct {
// 	Name string `yaml:"name"`
// 	Run  string `yaml:"run"`
// }

// func (node *StepNode) UnmarshalYAML(value *yaml.Node) error {
// 	node.Raw = value
// 	return value.Decode(&node.Value)
// }