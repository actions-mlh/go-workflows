package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"io"
	
	"gh-actions-checker/parser"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		process(io.ReadAll(os.Stdin))
	}

	for _, arg := range args {
		process(os.ReadFile(arg))
	}
}

func process(data []byte, err error) error {
	if err != nil {
		log.Fatal(err)
	}

	_, err = parser.Parse(data)
	if err != nil {
		log.Fatalf("error parsing stdin:\n    %v", err)
	}

	fmt.Println("passed!")
	return nil
}
