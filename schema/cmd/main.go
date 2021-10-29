// The schema-generate binary reads the JSON schema files passed as arguments
// and outputs the corresponding Go structs.
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	// "github.com/davecgh/go-spew/spew"

	"c2c-actions-mlh-workflow-parser/schema"
)

var (
	o                     = flag.String("o", "", "The output file for the schema.")
	p                     = flag.String("p", "gen_schema", "The package that the structs are created in.")
	i                     = flag.String("i", "", "A single file path (used for backwards compatibility).")
	schemaKeyRequiredFlag = flag.Bool("schemaKeyRequired", false, "Allow input files with no $schema key.")
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "  paths")
		fmt.Fprintln(os.Stderr, "\tThe input JSON Schema files.")
	}

	flag.Parse()

	inputFiles := flag.Args()
	if *i != "" {
		inputFiles = append(inputFiles, *i)
	}
	if len(inputFiles) == 0 {
		fmt.Fprintln(os.Stderr, "No input JSON Schema files.")
		flag.Usage()
		os.Exit(1)
	}

	schemas, err := generate.ReadInputFiles(inputFiles, *schemaKeyRequiredFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	g := generate.New(schemas...)

	// spew.Dump(g)
	err = g.CreateTypes()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failure generating structs: ", err)
		os.Exit(1)
	}

	var w io.Writer = os.Stdout

	if *o != "" {
		w, err = os.Create(*o)

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error opening output file: ", err)
			return
		}
	}
	generate.Output(w, g, *p)
}
