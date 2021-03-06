package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/actions-mlh/go-workflows/lint"
)

var i = flag.String("i", "", "Name of file to lint.  Equivalent to a command-line argument.")
var o = flag.String("o", "", "A custom output file.  Defaults to stdout.")
var h = flag.Bool("h", false, "Print instructions for how to use this tool.")
var d = flag.Bool("d", false, "Debug: see full tree of workflow.")

func main() {
	flag.Parse()
	args := flag.Args()

	if *i != "" {
		args = append(args, *i)
	}

	printHelp := false
	spew := false
	if len(args) == 0 {
		printHelp = true
	}

	flag.Visit(func(f *flag.Flag) {
		if f.Name == "h" {
			printHelp = true
		}
		if f.Name == "d" {
			spew = true
		}
	})

	if printHelp {
		fmt.Println("Usage: go run main.go [-i inputfile] [-o outputfile] inputfiles...")
		fmt.Println("Note that Go programs require all flags to appear before command line arguments.")
		flag.PrintDefaults()
		return
	}

	w := os.Stdout
	var err error
	if *o != "" {
		w, err = os.Create(*o)
		if err != nil {
			log.Fatal(err)
		}
	}
	defer w.Close()

	for _, filename := range args {
		input, err := os.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}
		if spew {
			lint.Spew(input)
		}
		problems, err := lint.Lint(filename, input)
		if err != nil {
			log.Fatal(err)
		}
		for _, problem := range problems {
			fmt.Fprintln(w, problem)
		}
	}
}
