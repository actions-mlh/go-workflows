package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"c2c-actions-mlh-workflow-parser/lint"
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
	input, err := os.ReadFile(inputFilename)
	if err != nil {
		return err
	}
	problems, err := lint.Lint(inputFilename, input)
	if err != nil {
		return err
	}
	for _, problem := range problems {
		fmt.Println(problem)
	}
	return nil
}


