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
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}

		_, err = parser.Parse(data)
		if err != nil {
			log.Fatalf("error parsing stdin:\n    %v", err)
		}

		fmt.Println("passed!")
	}

	for _, arg := range args {
		data, err := os.ReadFile(arg)
		if err != nil {
			log.Fatal(err)
		}

		_, err = parser.Parse(data)
		if err != nil {
			log.Fatalf("error parsing %v:\n    %v", arg, err)
		}

		// success case?
		fmt.Println("passed!")
	}
}
