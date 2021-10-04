package main

import (
	"flag"
	"io"
	"fmt"
	"os"

	"gh-actions-checker/parser"
)

func main() {
	flag.Parse()
	args := flag.Args()

	error := false
	if len(args) == 0 {
		data, err := io.ReadAll(os.Stdin)
		error = process("stdin", data, err)
	}

	for _, arg := range args {
		data, err := os.ReadFile(arg)
		error = error || process(arg, data, err)
	}
	if error {
		os.Exit(1)
	}
	os.Exit(0)
}

func process(name string, data []byte, err error) bool {
	error := false
	if err != nil {
		fmt.Println(err)
		error = true
	}

	_, err = parser.Parse(data)
	if err != nil {
		fmt.Printf("error parsing %v:\n    %v\n", name, err)
		error = true
	}
	return error
}
