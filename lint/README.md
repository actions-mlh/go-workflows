## explanation

The essential package is provided in `lint.go`, which just provides two functions, `Lint()` and `Spew()`.  `Lint()` is the important function; it will take a filename and a `[]byte` representing an input file, and return a `[]string` where each line contains an error message.  `Spew()` will print the internal representation of a YAML file for debug purposes.

Each subdirectory contains a small package used in `lint.go`.  `workflow/` contains the Go types used to represent the YAML structure, separated into different files for the different possible pieces of a GitHub workflow.  Each other directory contains a small package for linting that section of the workflow, supplying a single `Lint()` function to be called in `lint.go`.
