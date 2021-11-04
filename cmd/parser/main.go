package main

import (
	"flag"
	"log"
	"os"
	"gopkg.in/yaml.v3"
<<<<<<< HEAD:parse/cmd/parser/main.go
	"c2c-actions-mlh-workflow-parser/parse/lint"
	"c2c-actions-mlh-workflow-parser/parse/sink"
	// "c2c-actions-mlh-workflow-parser/gen_mock"
	"c2c-actions-mlh-workflow-parser/gen"
=======
	"c2c-actions-mlh-workflow-parser/lint"
	"c2c-actions-mlh-workflow-parser/sink"
	"c2c-actions-mlh-workflow-parser/workflow"
>>>>>>> a1699b0f54f71ac4d70ae5d6e1a227cf78ff1d54:cmd/parser/main.go
)

func main() {
	inputFilename := flag.String("i", "", "name of the file to lint")
	flag.Parse()

	if *inputFilename == "" {
		log.Fatalf("-i is required")
	}
	if err := realMain(*inputFilename); err != nil {
		log.Fatal(err)
	}
}

func realMain(inputFilename string) error {
	input, err := os.Open(inputFilename)
	if err != nil {
		return err
	}
	defer input.Close()

	sink := &sink.ProblemSink{Filename: inputFilename, Output: os.Stdout}
<<<<<<< HEAD:parse/cmd/parser/main.go
	node := new(gen.Root)
=======
	node := new(workflow.WorkflowNode)
>>>>>>> a1699b0f54f71ac4d70ae5d6e1a227cf78ff1d54:cmd/parser/main.go
	
	if err := yaml.NewDecoder(input).Decode(&node); err != nil {
		return err
	}

	if err := lint.LintWorkflowRoot(sink, node); err != nil {
		return err
	}
	sink.Render()

	return nil
}


