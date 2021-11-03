# Github Workflow Parser

A Go command line tool for linting and validating GitHub Actions.

## Usage

Clone this repository and run

``
go run ./cmd/parser/main.go yourfilehere.yaml
``

will invoke `main.go` and attempt to lint your YAML workflow.  If any problems are detected, they will be emitted to stdout; if not, then the output will be empty.

## Details

The code in `parse/` will leverage [this package](https://pkg.go.dev/gopkg.in/yaml.v3) to read the YAML file, and will attempt to cast the YAML it reads into the types specified in `workflow/`.  Any mismatches or detected errors will be collected into a sink and dumped to stdout.  `yaml/` contains sample YAML files for testing.

## Contact

This project is maintained by @jjayeon and @hankc97, under the guidance of @aybabtme and @joshmgross.  Please contact any of us for more information.
