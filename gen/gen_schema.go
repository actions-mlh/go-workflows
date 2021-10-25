// Code generated by schema-generate. DO NOT EDIT.
package gen_schema
import (
	"gopkg.in/yaml.v3"
)

// Definitions_Concurrency 
type Definitions_Concurrency struct {

  // To cancel any currently running job or workflow in the same concurrency group, specify cancel-in-progress: true.
  Definitions_Concurrency_CancelInProgress Definitions_Concurrency_CancelInProgressRaw `yaml:"cancel-in-progress,omitempty"`

  // When a concurrent job or workflow is queued, if another job or workflow using the same concurrency group in the repository is in progress, the queued job or workflow will be pending. Any previously pending job or workflow in the concurrency group will be canceled.
  Definitions_Concurrency_Group Definitions_Concurrency_GroupRaw `yaml:"group"`
}

// Definitions_Container 
type Definitions_Container struct {

  // If the image's container registry requires authentication to pull the image, you can use credentials to set a map of the username and password. The credentials are the same values that you would provide to the `docker login` command.
  Definitions_Container_Credentials Definitions_Container_CredentialsRaw `yaml:"credentials,omitempty"`

  // Sets an array of environment variables in the container.
  Definitions_Container_Env Definitions_Container_EnvRaw `yaml:"env,omitempty"`

  // The Docker image to use as the container to run the action. The value can be the Docker Hub image name or a registry name.
  Definitions_Container_Image Definitions_Container_ImageRaw `yaml:"image"`

  // Additional Docker container resource options. For a list of options, see https://docs.docker.com/engine/reference/commandline/create/#options.
  Definitions_Container_Options Definitions_Container_OptionsRaw `yaml:"options,omitempty"`

  // Sets an array of ports to expose on the container.
  Definitions_Container_Ports Definitions_Container_PortsRaw `yaml:"ports,omitempty"`

  // Sets an array of volumes for the container to use. You can use volumes to share data between services or other steps in a job. You can specify named Docker volumes, anonymous Docker volumes, or bind mounts on the host.
  // To specify a volume, you specify the source and destination path: <source>:<destinationPath>
  // The <source> is a volume name or an absolute path on the host machine, and <destinationPath> is an absolute path in the container.
  Definitions_Container_Volumes Definitions_Container_VolumesRaw `yaml:"volumes,omitempty"`
}

// Definitions_Container_Credentials If the image's container registry requires authentication to pull the image, you can use credentials to set a map of the username and password. The credentials are the same values that you would provide to the `docker login` command.
type Definitions_Container_Credentials struct {
  Definitions_Container_Credentials_Password Definitions_Container_Credentials_PasswordRaw `yaml:"password,omitempty"`
  Definitions_Container_Credentials_Username Definitions_Container_Credentials_UsernameRaw `yaml:"username,omitempty"`
}

// Definitions_Defaults 
type Definitions_Defaults struct {
  Definitions_Defaults_Run Definitions_Defaults_RunRaw `yaml:"run,omitempty"`
}

// Definitions_Defaults_Run 
type Definitions_Defaults_Run struct {
  Definitions_Defaults_Run_Shell Definitions_Defaults_Run_ShellRaw `yaml:"shell,omitempty"`
  Definitions_Defaults_Run_WorkingDirectory Definitions_Defaults_Run_WorkingDirectoryRaw `yaml:"working-directory,omitempty"`
}

// Definitions_Env 
type Definitions_Env struct {
  AdditionalProperties AdditionalPropertiesRaw `yaml:"-,omitempty"`
}

// Definitions_Environment The environment that the job references
type Definitions_Environment struct {

  // The name of the environment configured in the repo.
  Definitions_Environment_Name Definitions_Environment_NameRaw `yaml:"name"`

  // A deployment URL
  Definitions_Environment_Url Definitions_Environment_UrlRaw `yaml:"url,omitempty"`
}

// Definitions_PermissionsEvent 
type Definitions_PermissionsEvent struct {
  Definitions_PermissionsEvent_Actions Definitions_PermissionsEvent_ActionsRaw `yaml:"actions,omitempty"`
  Definitions_PermissionsEvent_Checks Definitions_PermissionsEvent_ChecksRaw `yaml:"checks,omitempty"`
  Definitions_PermissionsEvent_Contents Definitions_PermissionsEvent_ContentsRaw `yaml:"contents,omitempty"`
  Definitions_PermissionsEvent_Deployments Definitions_PermissionsEvent_DeploymentsRaw `yaml:"deployments,omitempty"`
  Definitions_PermissionsEvent_Issues Definitions_PermissionsEvent_IssuesRaw `yaml:"issues,omitempty"`
  Definitions_PermissionsEvent_Packages Definitions_PermissionsEvent_PackagesRaw `yaml:"packages,omitempty"`
  Definitions_PermissionsEvent_PullRequests Definitions_PermissionsEvent_PullRequestsRaw `yaml:"pull-requests,omitempty"`
  Definitions_PermissionsEvent_RepositoryProjects Definitions_PermissionsEvent_RepositoryProjectsRaw `yaml:"repository-projects,omitempty"`
  Definitions_PermissionsEvent_SecurityEvents Definitions_PermissionsEvent_SecurityEventsRaw `yaml:"security-events,omitempty"`
  Definitions_PermissionsEvent_Statuses Definitions_PermissionsEvent_StatusesRaw `yaml:"statuses,omitempty"`
}

// Definitions_Ref 
type Definitions_Ref struct {
  Definitions_Ref_Branches Definitions_Ref_BranchesRaw `yaml:"branches,omitempty"`
  Definitions_Ref_BranchesIgnore Definitions_Ref_BranchesIgnoreRaw `yaml:"branches-ignore,omitempty"`
  Definitions_Ref_Paths Definitions_Ref_PathsRaw `yaml:"paths,omitempty"`
  Definitions_Ref_PathsIgnore Definitions_Ref_PathsIgnoreRaw `yaml:"paths-ignore,omitempty"`
  Definitions_Ref_Tags Definitions_Ref_TagsRaw `yaml:"tags,omitempty"`
  Definitions_Ref_TagsIgnore Definitions_Ref_TagsIgnoreRaw `yaml:"tags-ignore,omitempty"`
}

// Properties 
type Properties struct {

  // Concurrency ensures that only a single job or workflow using the same concurrency group will run at a time. A concurrency group can be any string or expression. The expression can use any context except for the secrets context. 
  // You can also specify concurrency at the workflow level. 
  // When a concurrent job or workflow is queued, if another job or workflow using the same concurrency group in the repository is in progress, the queued job or workflow will be pending. Any previously pending job or workflow in the concurrency group will be canceled. To also cancel any currently running job or workflow in the same concurrency group, specify cancel-in-progress: true.
  Properties_Concurrency Properties_ConcurrencyRaw `yaml:"concurrency,omitempty"`

  // A map of default settings that will apply to all jobs in the workflow.
  Properties_Defaults Properties_DefaultsRaw `yaml:"defaults,omitempty"`

  // A map of environment variables that are available to all jobs and steps in the workflow.
  Properties_Env Properties_EnvRaw `yaml:"env,omitempty"`

  // A workflow run is made up of one or more jobs. Jobs run in parallel by default. To run jobs sequentially, you can define dependencies on other jobs using the jobs.<job_id>.needs keyword.
  // Each job runs in a fresh instance of the virtual environment specified by runs-on.
  // You can run an unlimited number of jobs as long as you are within the workflow usage limits. For more information, see https://help.github.com/en/github/automating-your-workflow-with-github-actions/workflow-syntax-for-github-actions#usage-limits.
  Properties_Jobs Properties_JobsRaw `yaml:"jobs"`

  // The name of your workflow. GitHub displays the names of your workflows on your repository's actions page. If you omit this field, GitHub sets the name to the workflow's filename.
  Properties_Name Properties_NameRaw `yaml:"name,omitempty"`

  // The name of the GitHub event that triggers the workflow. You can provide a single event string, array of events, array of event types, or an event configuration map that schedules a workflow or restricts the execution of a workflow to specific files, tags, or branch changes. For a list of available events, see https://help.github.com/en/github/automating-your-workflow-with-github-actions/events-that-trigger-workflows.
  Properties_On Properties_OnRaw `yaml:"on"`
  Properties_Permissions Properties_PermissionsRaw `yaml:"permissions,omitempty"`
}

// Properties_Jobs A workflow run is made up of one or more jobs. Jobs run in parallel by default. To run jobs sequentially, you can define dependencies on other jobs using the jobs.<job_id>.needs keyword.
// Each job runs in a fresh instance of the virtual environment specified by runs-on.
// You can run an unlimited number of jobs as long as you are within the workflow usage limits. For more information, see https://help.github.com/en/github/automating-your-workflow-with-github-actions/workflow-syntax-for-github-actions#usage-limits.
type Properties_Jobs struct {
}

type Definitions_Concurrency_CancelInProgressRaw struct {
	Raw *yaml.Node
	Value bool
}


func (node *Definitions_Concurrency_CancelInProgressRaw) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}


type Definitions_Concurrency_GroupRaw struct {
	Raw *yaml.Node
	Value string
}


func (node *Definitions_Concurrency_GroupRaw) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}


type Definitions_Container_CredentialsRaw struct {
	Raw *yaml.Node
	Value *Definitions_Container_Credentials
}


func (node *Definitions_Container_CredentialsRaw) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}


type Definitions_Container_EnvRaw struct {
	Raw *yaml.Node
	Value *Definitions_Env
}


func (node *Definitions_Container_EnvRaw) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}


type Definitions_Container_ImageRaw struct {
	Raw *yaml.Node
	Value string
}


func (node *Definitions_Container_ImageRaw) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}


type Definitions_Container_OptionsRaw struct {
	Raw *yaml.Node
	Value string
}


func (node *Definitions_Container_OptionsRaw) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}


type Definitions_Container_PortsRaw struct {
	Raw *yaml.Node
	Value []interface{}
}


func (node *Definitions_Container_PortsRaw) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}


type Definitions_Container_VolumesRaw struct {
	Raw *yaml.Node
	Value []string
}


func (node *Definitions_Container_VolumesRaw) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}


type Definitions_Container_Credentials_PasswordRaw struct {
	Raw *yaml.Node
	Value string
}


func (node *Definitions_Container_Credentials_PasswordRaw) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}


type Definitions_Container_Credentials_UsernameRaw struct {
	Raw *yaml.Node
	Value string
}


func (node *Definitions_Container_Credentials_UsernameRaw) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}


type Definitions_Defaults_RunRaw struct {
	Raw *yaml.Node
	Value *Definitions_Defaults_Run
}


func (node *Definitions_Defaults_RunRaw) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}


type Definitions_Defaults_Run_ShellRaw struct {
	Raw *yaml.Node
	Value string
}


func (node *Definitions_Defaults_Run_ShellRaw) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}


type Definitions_Defaults_Run_WorkingDirectoryRaw struct {
	Raw *yaml.Node
	Value string
}


func (node *Definitions_Defaults_Run_WorkingDirectoryRaw) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}


type AdditionalPropertiesRaw struct {
	Raw *yaml.Node
	Value map[string]interface{}
}


func (node *AdditionalPropertiesRaw) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}


type Definitions_Environment_NameRaw struct {
	Raw *yaml.Node
	Value string
}


func (node *Definitions_Environment_NameRaw) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}


type Definitions_Environment_UrlRaw struct {
	Raw *yaml.Node
	Value string
}


func (node *Definitions_Environment_UrlRaw) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}


type Definitions_PermissionsEvent_ActionsRaw struct {
	Raw *yaml.Node
	Value string
}


func (node *Definitions_PermissionsEvent_ActionsRaw) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}


type Definitions_PermissionsEvent_ChecksRaw struct {
	Raw *yaml.Node
	Value string
}


func (node *Definitions_PermissionsEvent_ChecksRaw) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}


type Definitions_PermissionsEvent_ContentsRaw struct {
	Raw *yaml.Node
	Value string
}


func (node *Definitions_PermissionsEvent_ContentsRaw) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}


type Definitions_PermissionsEvent_DeploymentsRaw struct {
	Raw *yaml.Node
	Value string
}


func (node *Definitions_PermissionsEvent_DeploymentsRaw) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}


type Definitions_PermissionsEvent_IssuesRaw struct {
	Raw *yaml.Node
	Value string
}


func (node *Definitions_PermissionsEvent_IssuesRaw) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}


type Definitions_PermissionsEvent_PackagesRaw struct {
	Raw *yaml.Node
	Value string
}


func (node *Definitions_PermissionsEvent_PackagesRaw) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}


type Definitions_PermissionsEvent_PullRequestsRaw struct {
	Raw *yaml.Node
	Value string
}


func (node *Definitions_PermissionsEvent_PullRequestsRaw) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}


type Definitions_PermissionsEvent_RepositoryProjectsRaw struct {
	Raw *yaml.Node
	Value string
}


func (node *Definitions_PermissionsEvent_RepositoryProjectsRaw) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}


type Definitions_PermissionsEvent_SecurityEventsRaw struct {
	Raw *yaml.Node
	Value string
}


func (node *Definitions_PermissionsEvent_SecurityEventsRaw) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}


type Definitions_PermissionsEvent_StatusesRaw struct {
	Raw *yaml.Node
	Value string
}


func (node *Definitions_PermissionsEvent_StatusesRaw) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}


type Definitions_Ref_BranchesRaw struct {
	Raw *yaml.Node
	Value []string
}


func (node *Definitions_Ref_BranchesRaw) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}


type Definitions_Ref_BranchesIgnoreRaw struct {
	Raw *yaml.Node
	Value []string
}


func (node *Definitions_Ref_BranchesIgnoreRaw) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}


type Definitions_Ref_PathsRaw struct {
	Raw *yaml.Node
	Value []string
}


func (node *Definitions_Ref_PathsRaw) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}


type Definitions_Ref_PathsIgnoreRaw struct {
	Raw *yaml.Node
	Value []string
}


func (node *Definitions_Ref_PathsIgnoreRaw) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}


type Definitions_Ref_TagsRaw struct {
	Raw *yaml.Node
	Value []string
}


func (node *Definitions_Ref_TagsRaw) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}


type Definitions_Ref_TagsIgnoreRaw struct {
	Raw *yaml.Node
	Value []string
}


func (node *Definitions_Ref_TagsIgnoreRaw) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}


type Properties_ConcurrencyRaw struct {
	Raw *yaml.Node
	Value interface{}
}


func (node *Properties_ConcurrencyRaw) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}


type Properties_DefaultsRaw struct {
	Raw *yaml.Node
	Value *Definitions_Defaults
}


func (node *Properties_DefaultsRaw) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}


type Properties_EnvRaw struct {
	Raw *yaml.Node
	Value *Definitions_Env
}


func (node *Properties_EnvRaw) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}


type Properties_JobsRaw struct {
	Raw *yaml.Node
	Value *Properties_Jobs
}


func (node *Properties_JobsRaw) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}


type Properties_NameRaw struct {
	Raw *yaml.Node
	Value string
}


func (node *Properties_NameRaw) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}


type Properties_OnRaw struct {
	Raw *yaml.Node
	Value interface{}
}


func (node *Properties_OnRaw) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}


type Properties_PermissionsRaw struct {
	Raw *yaml.Node
	Value interface{}
}


func (node *Properties_PermissionsRaw) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	return value.Decode(&node.Value)
}

