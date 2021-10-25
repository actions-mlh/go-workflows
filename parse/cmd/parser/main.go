package main

import (
	"flag"
	"log"
	"os"
	"gopkg.in/yaml.v3"
	"c2c-actions-mlh-workflow-parser/gen_mock"
	"c2c-actions-mlh-workflow-parser/parse/lint"
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

	sink := &lint.ProblemSink{Filename: inputFilename, Output: os.Stdout}
	node := new(gen_mock.WorkflowNode)
	
	if err := yaml.NewDecoder(input).Decode(&node); err != nil {
		return err
	}

	if err := lint.LintWorkflow(sink, node); err != nil {
		return err
	}

	return nil
}


