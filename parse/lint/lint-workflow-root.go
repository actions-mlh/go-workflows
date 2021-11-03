package lint

import (
	"c2c-actions-mlh-workflow-parser/gen_mock"
	"c2c-actions-mlh-workflow-parser/parse/lint/util"
	"c2c-actions-mlh-workflow-parser/parse/sink"
	"gopkg.in/yaml.v3"
	"reflect"
	"strings"
	"fmt"
)

// issues
// 1) Is Kind a marshal error? -> Decided to only add support for kind of scalar type (!!bool, !!float, !!int, !!str, ...)

func LintWorkflowRoot(sink *sink.ProblemSink, target *gen_mock.WorkflowNode) error {
	workflowKeyNodes := []*yaml.Node{}
	workflowValueNodes := []*yaml.Node{}

	for i := 0; i < len(target.Raw.Content); i += 2 {
		workflowKeyNodes = append(workflowKeyNodes, target.Raw.Content[i])
		workflowValueNodes = append(workflowValueNodes, target.Raw.Content[i+1])
	}

	requiredKeys := map[string]bool{"on": false, "jobs": false}
	if err := util.CheckRequiredKeys(target.Raw, sink, workflowKeyNodes, requiredKeys); err != nil {
		return err
	}

	if err := util.CheckNullPointer(sink, workflowKeyNodes, workflowValueNodes); err != nil {
		return err
	}

	if len(workflowKeyNodes) != len(target.Value) {
		if err := util.CheckDuplicateKeys(sink, workflowKeyNodes); err != nil {
			return err
		}

		expectedKeys := []string{}
		reflectedStruct := reflect.ValueOf(*target.Value[0])
		typeOfStruct := reflectedStruct.Type()
		for i := 0; i < reflectedStruct.NumField(); i++ {
			expectedKeys = append(expectedKeys, strings.ToLower(typeOfStruct.Field(i).Name))
		}

		if err := util.CheckUnexpectedKeys(sink, expectedKeys, workflowKeyNodes); err != nil {
			return err
		}
	}

	workflowValues := make(map[string]int)
	for _, value := range target.Value {
		reflectedValues := reflect.ValueOf(*value)
		typeOfValue := reflectedValues.Type()
		for i := 0; i< reflectedValues.NumField(); i++ {
			key := typeOfValue.Field(i).Name
			fieldVal := reflectedValues.Field(i)

			if _, contains := workflowValues[key]; !contains && !fieldVal.IsNil() {
				workflowValues[key] = 0

				switch key {
				case "Name":
					if len(value.Name.ScalarTypes) != 0 {
						util.CheckUnexpectedScalarTypes(sink, value.Name.Raw, value.Name.ScalarTypes)
					}
				case "On":
					if len(value.On.ScalarTypes) != 0 {
						util.CheckUnexpectedScalarTypes(sink, value.On.Raw, value.On.ScalarTypes)
					}
				case "Jobs":
					if len(value.Jobs.ScalarTypes) != 0 {
						util.CheckUnexpectedScalarTypes(sink, value.Jobs.Raw, value.Jobs.ScalarTypes)
					}
				} 
			}
		}
	}

	fmt.Println("-------TESTING--------")
	for i := 0; i < len(workflowValueNodes); i += 1 {
		// fmt.Printf("%+v\n", target.Value[i])
		// fmt.Printf("%+v\n", workflowKeyNodes[i])
		// fmt.Printf("%+v\n", workflowValueNodes[i])
	}


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

// 			remainingString := jobValueID[1:]
// 			alphabetValidation := regexp.MustCompile(`^[A-Za-z]+$`).MatchString
// 			if !alphabetValidation(remainingString) {
// 				sink.Record(jobRaw, "job ID's must contain only alphanumeric characters \"-\", or \"_\"")
// 			}
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