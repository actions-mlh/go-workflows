package lint

import (
	"c2c-actions-mlh-workflow-parser/gen_mock"
	// "fmt"
	// "reflect"
	"strings"
	"gopkg.in/yaml.v3"
	"regexp"
)

func checkRequiredKeys(raw *yaml.Node, sink *ProblemSink, keys []*yaml.Node, requiredKeys map[string]bool) error {
	for _, key := range keys {
		if _, contains := requiredKeys[key.Value]; contains {
			requiredKeys[key.Value] = true
		}
	}

	required := []string{}
	for k, v := range requiredKeys {
		if !v{
			required = append(required, k)
		}
	}

	if len(required) != 0 {
		sink.Record(raw, "Required Keys: %s", strings.Join(required, ","))
	}


	return nil
}

func checkDuplicateKeys(raw *yaml.Node, sink *ProblemSink, nodeKeys []*yaml.Node ) error {
	nonDuplicateKeys := make(map[string]int) 

	for _, key := range nodeKeys {
		if _, contains := nonDuplicateKeys[key.Value]; !contains {
			nonDuplicateKeys[key.Value] = 1
		} else {
			nonDuplicateKeys[key.Value]++
			sink.Record(key, "Duplicate Keys: %s", key.Value)
		}
	}
	return nil
}

func LintWorkflow(sink *ProblemSink, target *gen_mock.WorkflowNode) error {
	workflowKeyNodes := []*yaml.Node{}

	for i := 0; i < len(target.Raw.Content); i+=2 {
		workflowKeyNodes = append(workflowKeyNodes, target.Raw.Content[i])
	}

	if err := checkDuplicateKeys(target.Raw, sink, workflowKeyNodes); err != nil {
		return err
	}

	requiredKeys := map[string]bool{"on": false, "jobs": false}
	if err := checkRequiredKeys(target.Raw, sink, workflowKeyNodes, requiredKeys); err != nil {
		return err
	}

	// if err := lintWorkflowName(sink, workflow.Name, target.Raw); err != nil {
	// 	return err
	// }

	// if err := lintWorkflowJobs(sink, workflow.Jobs, target.Raw); err != nil {
	// 	return err
	// }
	return nil
}

func lintWorkflowName(sink *ProblemSink, target *gen_mock.WorkflowNameNode, raw *yaml.Node) error {
	nameNode := target

	if nameNode != nil {
		if nameNode.Raw == nil {
			sink.Record(raw, "name cannot be null")
		}
	}

	return nil
}

func lintWorkflowJobs(sink *ProblemSink, target *gen_mock.WorkflowJobsNode, raw *yaml.Node) error {
	jobsNode := target

	if jobsNode != nil {
		if jobsNode.Raw == nil {
			sink.Record(raw, "jobs cannot be null")
		}

		jobIDArray := []string{}
		validateJobID := func(jobValueID string, jobRaw *yaml.Node) error {
			firstLetter := jobValueID[0:1]
			for _, runeVal := range firstLetter {
				if (runeVal < 'a' || runeVal > 'z') && (runeVal < 'A' || runeVal > 'Z') && (runeVal != '_') {
					sink.Record(jobRaw, "job ID's must start with a letter or \"_\"")
				}
			}

			remainingString := jobValueID[1:]
			alphabetValidation := regexp.MustCompile(`^[A-Za-z]+$`).MatchString
			if !alphabetValidation(remainingString) {
				sink.Record(jobRaw, "job ID's must contain only alphanumeric characters \"-\", or \"_\"")
			}
			return nil
		}

		validateCircularNeeds := func(needsNode *gen_mock.JobsPatternPropertiesNeedsNode, jobValueID string) {
			raw := needsNode.Raw
			sequence := *needsNode.OneOf.SequenceNode

			for _, sequenceID := range sequence {
				contains := false
				for _, jobID := range jobIDArray {
					if jobID == sequenceID {
						contains = true
					}
				}
				if !contains {
					sink.Record(raw, "job %s does not exist", sequenceID)
				}
				contains = false
			}

			for _, sequenceID := range sequence {
				if sequenceID == jobValueID {
					sink.Record(raw, "cannot contain itself within its job needs")
				}
			}
		}

		for _, jobValue := range jobsNode.Value { 
			jobIDArray = append(jobIDArray, jobValue.ID)
		}


		for _, jobValue := range jobsNode.Value {
			validateJobID(jobValue.ID, jobValue.PatternProperties.Raw)
		} 

		for _, jobValue := range jobsNode.Value {
			validateCircularNeeds(jobValue.PatternProperties.Value.Needs , jobValue.ID)
		} 
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
