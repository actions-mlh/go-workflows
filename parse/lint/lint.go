package lint

import (
	// "strings"
	// "c2c-actions-mlh-workflow-parser/gen_mock"
)

// use this to convert interface{} into slice of types
// s := reflect.ValueOf(*node.OneOf.ValueArray)
// 	for i := 0; i < s.Len(); i++ {
// 		fmt.Println(s.Index(i))
// }

func LintWorkflow() error { // (sink *ProblemSink, workflow *gen_mock.WorkflowNode) error {
	// fmt.Printf("%+v\n", *workflow.Value.On.OneOf.MappingNode)
	// fmt.Printf("%+v\n", *&workflow.Value.Concurrency.OneOf.MappingNode.CancelInProgress.Value)
	// fmt.Println("------------------------------")
	// fmt.Printf("%+v\n", *workflow.Value.On.OneOf.MappingNode)
	// sink.Record(workflow.Raw, "test error")
	

	// fmt.Printf("%+v\n", *workflow.Value.On.OneOf.MappingNode.CheckRun.OneOf.MappingNode.Types.Value)

	// fmt.Printf("%+v\n", *workflow.Value.On.OneOf.MappingNode.CheckRun.OneOf.MappingNode[0].Types.Value)
	// fmt.Printf("%+v\n", *workflow.Value.On.OneOf.MappingNode.CheckSuite)
	// for _, node := range workflow.Value.On.OneOf.MappingNode {
		// fmt.Printf("%+v\n", *node)
		// fmt.Printf("%+v\n", *node.CheckRun.OneOf.MappingNode[0].Types.Value)
	// }

	// if err := lintWorkflowName(sink, workflow.Name); err != nil {
	// 	return err
	// }

	// if err := lintWorkflowJobs(sink, workflow.Jobs.Value.Steps); err != nil {
	// 	return err
	// }

	// for _, step := range workflow.Jobs.Value.Steps.Value {
	// 	if err := lintWorkflowJobs(sink, step); err != nil {
	// 		return err
	// 	}
	// }
	return nil
}

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
