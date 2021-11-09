package lint

import (
	"reflect"
	"strings"
	"gopkg.in/yaml.v3"
)

// issues
// 1) Is Kind a marshal error? -> Decided to only add support for kind of scalar type (!!bool, !!float, !!int, !!str, ...)

func lintWorkflowRoot(sink *problemSink, target *WorkflowNode) error {
	workflowKeyNodes := []*yaml.Node{}
	workflowValueNodes := []*yaml.Node{}

	for i := 0; i < len(target.Raw.Content); i += 2 {
		workflowKeyNodes = append(workflowKeyNodes, target.Raw.Content[i])
		workflowValueNodes = append(workflowValueNodes, target.Raw.Content[i+1])
	}

	requiredKeys := []string{"on", "jobs"}
	if err := checkRequiredKeys(target.Raw, sink, workflowKeyNodes, requiredKeys); err != nil {
		return err
	}

	if err := checkNullPointer(sink, workflowKeyNodes, workflowValueNodes); err != nil {
		return err
	}

	if err := checkDuplicateKeys(sink, workflowKeyNodes); err != nil {
		return err
	}

	expectedKeys := []string{}
	reflectedStruct := reflect.ValueOf(*target.Value)
	typeOfStruct := reflectedStruct.Type()
	for i := 0; i < reflectedStruct.NumField(); i++ {
		expectedKeys = append(expectedKeys, strings.ToLower(typeOfStruct.Field(i).Name))
	}

	if err := checkUnexpectedKeys(sink, expectedKeys, workflowKeyNodes); err != nil {
		return err
	}

	if target.Value.Name != nil {
		if err := checkUnexpectedScalarTypes(sink, target.Value.Name.Raw, []string{"!!str"}); err != nil {
			return err
		}
	}

	if target.Value.Jobs != nil && target.Value.Jobs.Raw != nil {
		if err := checkUnexpectedScalarTypes(sink, target.Value.Jobs.Raw, []string{"!!map"}); err != nil {
			return err
		}
		if err := checkJobNames(sink, target.Value.Jobs.Raw); err != nil {
			return err
		}
	}

	// fmt.Println("-------TESTING--------")
	// fmt.Println("-------HERE--------")

	// if err := lintWorkflowName(sink, workflow.Name, target.Raw); err != nil {
	// 	return err
	// }

	// if err := lintWorkflowJobs(sink, workflow.Jobs, target.Raw); err != nil {
	// 	return err
	// }
	return nil
}

// func lintWorkflowName(sink *ProblemSink, target *gen_mock.WorkflowNameNode, raw *yaml.Node) error {
// 	nameNode := target

// 	if nameNode != nil {
// 		if nameNode.Raw == nil {
// 			sink.Record(raw, "name cannot be null")
// 		}
// 	}

// 	return nil
// }

// func lintWorkflowJobs(sink *ProblemSink, target *gen_mock.WorkflowJobsNode, raw *yaml.Node) error {
// 	jobsNode := target

// 	if jobsNode != nil {
// 		if jobsNode.Raw == nil {
// 			sink.Record(raw, "jobs cannot be null")
// 		}

// 		jobIDArray := []string{}
// 		validateJobID := func(jobValueID string, jobRaw *yaml.Node) error {
// 			firstLetter := jobValueID[0:1]
// 			for _, runeVal := range firstLetter {
// 				if (runeVal < 'a' || runeVal > 'z') && (runeVal < 'A' || runeVal > 'Z') && (runeVal != '_') {
// 					sink.Record(jobRaw, "job ID's must start with a letter or \"_\"")
// 				}
// 			}

			// remainingString := jobValueID[1:]
			// alphabetValidation := regexp.MustCompile(`"^[_a-zA-Z][a-zA-Z0-9_-]*$"`).MatchString
			// if !alphabetValidation(remainingString) {
			// 	sink.Record(jobRaw, "job ID's must contain only alphanumeric characters \"-\", or \"_\"")
			// }
// 			return nil
// 		}

// 		validateCircularNeeds := func(needsNode *gen_mock.JobsPatternPropertiesNeedsNode, jobValueID string) {
// 			raw := needsNode.Raw
// 			sequence := *needsNode.OneOf.SequenceNode

// 			for _, sequenceID := range sequence {
// 				contains := false
// 				for _, jobID := range jobIDArray {
// 					if jobID == sequenceID {
// 						contains = true
// 					}
// 				}
// 				if !contains {
// 					sink.Record(raw, "job %s does not exist", sequenceID)
// 				}
// 				contains = false
// 			}

// 			for _, sequenceID := range sequence {
// 				if sequenceID == jobValueID {
// 					sink.Record(raw, "cannot contain itself within its job needs")
// 				}
// 			}
// 		}

// 		for _, jobValue := range jobsNode.Value {
// 			jobIDArray = append(jobIDArray, jobValue.ID)
// 		}

// 		for _, jobValue := range jobsNode.Value {
// 			validateJobID(jobValue.ID, jobValue.PatternProperties.Raw)
// 		}

// 		for _, jobValue := range jobsNode.Value {
// 			validateCircularNeeds(jobValue.PatternProperties.Value.Needs , jobValue.ID)
// 		}
// 	}

// 	return nil
// }

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
