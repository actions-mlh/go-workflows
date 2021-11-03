// Code generated by schema-generate. DO NOT EDIT.
package gen
import (
	"fmt"
	"gopkg.in/yaml.v3"
)

// Root 
type Root struct {
  // Concurrency ensures that only a single job or workflow using the same concurrency group will run at a time. A concurrency group can be any string or expression. The expression can use any context except for the secrets context. 
  // You can also specify concurrency at the workflow level. 
  // When a concurrent job or workflow is queued, if another job or workflow using the same concurrency group in the repository is in progress, the queued job or workflow will be pending. Any previously pending job or workflow in the concurrency group will be canceled. To also cancel any currently running job or workflow in the same concurrency group, specify cancel-in-progress: true.
	Concurrency *interface{} `yaml:"concurrency,omitempty"`
  // A map of default settings that will apply to all jobs in the workflow.
	Defaults *Root_Defaults `yaml:"defaults,omitempty"`
  // A map of environment variables that are available to all jobs and steps in the workflow.
	Env *interface{} `yaml:"env,omitempty"`
  // A workflow run is made up of one or more jobs. Jobs run in parallel by default. To run jobs sequentially, you can define dependencies on other jobs using the jobs.<job_id>.needs keyword.
  // Each job runs in a fresh instance of the virtual environment specified by runs-on.
  // You can run an unlimited number of jobs as long as you are within the workflow usage limits. For more information, see https://help.github.com/en/github/automating-your-workflow-with-github-actions/workflow-syntax-for-github-actions#usage-limits.
	Jobs *map[string]*interface{} `yaml:"jobs"`
  // The name of your workflow. GitHub displays the names of your workflows on your repository's actions page. If you omit this field, GitHub sets the name to the workflow's filename.
	Name *string `yaml:"name,omitempty"`
  // The name of the GitHub event that triggers the workflow. You can provide a single event string, array of events, array of event types, or an event configuration map that schedules a workflow or restricts the execution of a workflow to specific files, tags, or branch changes. For a list of available events, see https://help.github.com/en/github/automating-your-workflow-with-github-actions/events-that-trigger-workflows.
	On *interface{} `yaml:"on"`
	Permissions *interface{} `yaml:"permissions,omitempty"`
	Raw *yaml.Node
}

func (node *Root) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	for i := 0; i < len(value.Content); i++ {
		nodeName := value.Content[i]
		switch nodeName.Value {
			
			case "concurrency":
				i++
				if i >= len(value.Content) {
					return fmt.Errorf("value.Content mismatch")
				}
				nodeValue := value.Content[i]
				node.Concurrency = new(interface{})
				err := nodeValue.Decode(node.Concurrency)
				if err != nil {
					return err
				}
			
			case "defaults":
				i++
				if i >= len(value.Content) {
					return fmt.Errorf("value.Content mismatch")
				}
				nodeValue := value.Content[i]
				node.Defaults = new(Root_Defaults)
				err := nodeValue.Decode(node.Defaults)
				if err != nil {
					return err
				}
			
			case "env":
				i++
				if i >= len(value.Content) {
					return fmt.Errorf("value.Content mismatch")
				}
				nodeValue := value.Content[i]
				node.Env = new(interface{})
				err := nodeValue.Decode(node.Env)
				if err != nil {
					return err
				}
			
			case "jobs":
				i++
				if i >= len(value.Content) {
					return fmt.Errorf("value.Content mismatch")
				}
				nodeValue := value.Content[i]
				node.Jobs = new(map[string]*interface{})
				err := nodeValue.Decode(node.Jobs)
				if err != nil {
					return err
				}
			
			case "name":
				i++
				if i >= len(value.Content) {
					return fmt.Errorf("value.Content mismatch")
				}
				nodeValue := value.Content[i]
				node.Name = new(string)
				err := nodeValue.Decode(node.Name)
				if err != nil {
					return err
				}
			
			case "on":
				i++
				if i >= len(value.Content) {
					return fmt.Errorf("value.Content mismatch")
				}
				nodeValue := value.Content[i]
				node.On = new(interface{})
				err := nodeValue.Decode(node.On)
				if err != nil {
					return err
				}
			
			case "permissions":
				i++
				if i >= len(value.Content) {
					return fmt.Errorf("value.Content mismatch")
				}
				nodeValue := value.Content[i]
				node.Permissions = new(interface{})
				err := nodeValue.Decode(node.Permissions)
				if err != nil {
					return err
				}
			
		}
	}
	return nil
}

// Root_Defaults 
type Root_Defaults struct {
	Run *Root_Defaults_Run `yaml:"run,omitempty"`
	Raw *yaml.Node
}

func (node *Root_Defaults) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	for i := 0; i < len(value.Content); i++ {
		nodeName := value.Content[i]
		switch nodeName.Value {
			
			case "run":
				i++
				if i >= len(value.Content) {
					return fmt.Errorf("value.Content mismatch")
				}
				nodeValue := value.Content[i]
				node.Run = new(Root_Defaults_Run)
				err := nodeValue.Decode(node.Run)
				if err != nil {
					return err
				}
			
		}
	}
	return nil
}

// Root_Defaults_Run 
type Root_Defaults_Run struct {
	Shell *string `yaml:"shell,omitempty"`
	Working_Directory *string `yaml:"working-directory,omitempty"`
	Raw *yaml.Node
}

func (node *Root_Defaults_Run) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	for i := 0; i < len(value.Content); i++ {
		nodeName := value.Content[i]
		switch nodeName.Value {
			
			case "shell":
				i++
				if i >= len(value.Content) {
					return fmt.Errorf("value.Content mismatch")
				}
				nodeValue := value.Content[i]
				node.Shell = new(string)
				err := nodeValue.Decode(node.Shell)
				if err != nil {
					return err
				}
			
			case "working_directory":
				i++
				if i >= len(value.Content) {
					return fmt.Errorf("value.Content mismatch")
				}
				nodeValue := value.Content[i]
				node.Working_Directory = new(string)
				err := nodeValue.Decode(node.Working_Directory)
				if err != nil {
					return err
				}
			
		}
	}
	return nil
}
