package main

import (
	"flag"
	"io"
	"log"
	"os"

	"gh-actions-checker/parser"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		data, err := io.ReadAll(os.Stdin)
		process("stdin", data, err)
	}

	for _, arg := range args {
		data, err := os.ReadFile(arg)
		process(arg, data, err)
	}
}

func process(name string, data []byte, err error) {
	if err != nil {
		log.Print(err)
		return
	}

	_, err = parser.Parse(data)
	if err != nil {
		log.Printf("error parsing %v:\n    %v", name, err)
	}
}
