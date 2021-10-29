package lint

import (
	"c2c-actions-mlh-workflow-parser/gen_mock"
	"fmt"
	"reflect"
	"strings"
	"gopkg.in/yaml.v3"
)

// issues:
// 1) Check for duplicate keys in ?

func checkRequiredKeys(raw *yaml.Node, sink *ProblemSink, keys []interface{}, required map[string]bool) error {
	counter := len(required)

	for _, key := range keys {
		str := fmt.Sprintf("%v", key)
		_, ok := required[str]

		if ok {
			counter--
			required[str] = true
		}
	}

	requiredKeys := []string{}
	if counter != 0 {
		for k, v := range required {
			if !v {
				requiredKeys = append(requiredKeys, strings.ToLower(k))
			}
		}
		sink.Record(raw, "Required Keys: %s", strings.Join(requiredKeys, ","))
	}

	return nil
}

func LintWorkflow(sink *ProblemSink, target *gen_mock.WorkflowNode) error { // (sink *ProblemSink, workflow *gen_mock.WorkflowNode) error {
	workflow := target.Value

	// ---------------------------------------------------------------------------
	reflectedStructKeys := reflect.ValueOf(*workflow).Type()
	reflectedStructValues := reflect.ValueOf(*workflow)
	var keys []interface{}
	for i := 0; i < reflectedStructKeys.NumField(); i++ {
		if !reflect.ValueOf(reflectedStructValues.Field(i).Interface()).IsNil() {
			keys = append(keys, reflectedStructKeys.Field(i).Name)
		}
	}
	// ---------------------------------------------------------------------------

	if err := checkRequiredKeys(target.Raw, sink, keys, map[string]bool{"On": false, "Jobs": false}); err != nil {
		return err
	}

	if err := lintWorkflowName(sink, workflow.Name, target.Raw); err != nil {
		return err
	}

	if err := lintWorkflowJobs(sink, workflow.Jobs, target.Raw); err != nil {
		return err
	}
	return nil
}

func lintWorkflowName(sink *ProblemSink, target *gen_mock.WorkflowNameNode, raw *yaml.Node) error {
	nameNode := target

	if nameNode.Raw == nil {
		sink.Record(raw, "name cannot be null")
	}

	return nil
}

func lintWorkflowJobs(sink *ProblemSink, target *gen_mock.WorkflowJobsNode, raw *yaml.Node) error {
	jobsNode := target

	if jobsNode != nil {
		if jobsNode.Raw == nil {
			sink.Record(raw, "jobs cannot be null")
		}

		// jobIDArray := []string{}
		validateJobID := func(jobValueID string) error {
			return nil
		}
		// fmt.Printf("%+v\n", jobsNode.Raw.Content[2])
		for _, jobValue := range jobsNode.Value {
			validateJobID(jobValue.ID)

			fmt.Printf("%+v\n", *jobValue.PatternProperties.Value.Needs.OneOf.SequenceNode)
		} 
		// create tests for each yaml test that displays errors
	}

	return nil
}



// func (node WorkflowNode) Lint(sink *lint.ProblemSink) error {
// 	if len(node.Raw.Content)%2 != 0 {
// 		// Uneven set of key value pairs (this shouldn't happen)
// 		sink.Record(node.Raw, "%d:%d  error  Expected even number of key value pairs", node.Raw.Line, node.Raw.Column)
// 	}
// 	if node.Value.Name == nil {
// 		sink.Record(node.Raw, "%d:%d  error  missing name", node.Raw.Line, node.Raw.Column)
// 	}
// 	if node.Value.On == nil {
// 		sink.Record(node.Raw, "%d:%d  error  missing on", node.Raw.Line, node.Raw.Column)
// 	} else {
// 		node.Value.On.Lint(sink)
// 	}
// 	if node.Value.Jobs == nil {
// 		sink.Record(node.Raw, "%d:%d  error  missing jobs", node.Raw.Line, node.Raw.Column)
// 	} else {
// 		node.Value.Jobs.Lint(sink)
// 	}
// 	return nil
// }

// func (node *WorkflowOnNode) Lint(sink *lint.ProblemSink) error {
// 	if node.Raw.Kind == yaml.ScalarNode && node.Raw.Tag != "!!str" {
// 		sink.Record(node.Raw, "%d:%d  error  Expected one of: string type", node.Raw.Line, node.Raw.Column)
// 	}
// 	if len(node.Raw.Content)%2 != 0 {
// 		// Uneven set of key value pairs (this shouldn't happen)
// 		sink.Record(node.Raw, "%d:%d  error  Expected even number of key value pairs", node.Raw.Line, node.Raw.Column)
// 	}
// 	return nil
// }

// func (node *OnCheckRunNode) Lint(sink *lint.ProblemSink) error {
// 	return nil
// }

// func (node *WorkflowJobsNode) Lint(sink *lint.ProblemSink) error {
// 	if len(node.Raw.Content)%2 != 0 {
// 		// Uneven set of key value pairs (this shouldn't happen)
// 		sink.Record(node.Raw, "%d:%d  error  Expected even number of key value pairs", node.Raw.Line, node.Raw.Column)
// 	}
// 	return nil
// }


// func lintWorkflowName(sink *ProblemSink, target mock_gen_schema.NameRaw) error {
// 	name := target.Value
// 	if strings.HasPrefix(name, "HelloWorld") {
// 		sink.RecordProblem(target.Raw, "you're not allowed to say Hello World")
// 	}
// 	return nil
// }

// func lintWorkflowJobs(sink *ProblemSink, target gen_schema.StepsRaw) error {
// 	if len(target.Value) >= 4 {
// 		sink.RecordProblem(target.Raw, "can't use more than 3 steps")
// 	}
// 	return nil
// }

// if len(value.Steps) > 0 && value.Uses != "" {
// 	sink.RecordProblem(target.Raw, `can't use "steps" with "uses"`)
// }

// if err := lintWorkflowOn(sink, workflow.On); err != nil {
// 	return err
// }
