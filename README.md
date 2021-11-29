# Github Workflow Parser

Engineering demonstration by GitHub's MLH externs for an Action workflow parser and integration with developer tools. 

## Usage

Clone this repository and run
``
go run ./cmd/parser/main.go yourfilehere.yaml
``
which will invoke `main.go` and attempt to lint your YAML workflow.  If any problems are detected, they will be emitted to stdout; if not, then the output will be empty.  You may pass multiple files.

Available flags:
- `-i filename` to specify an input file, identical to a command line argument.
- `-o filename` to specify an output file.
- `-d` for debug info --- will print the internal representation of the YAML file.  This is useful if, for instance, you suspect the code can't reach some part of your file and want to confirm.
- `-h` to display the help text.

For testing, you can run
```
go test ./test/
```
Further explanation can be found in `test/main_test.go` and `yaml/clean/` or `yaml/dirty`.

## Details

The code in `lint/` will leverage [this package](https://pkg.go.dev/gopkg.in/yaml.v3) to read the YAML file, and will attempt to cast the YAML it reads into the types specified in `lint/workflow/`.  Any mismatches or detected errors will be collected into a sink and dumped to stdout.  `yaml/` contains sample YAML files for testing, organized into `clean/` (no errors) and `dirty/` (deliberate errors specified in `filename.yaml.exp`).  See each directory for further information.

For future maintainers, a recommended development procedure is as follows:
1. Identify a new class of error you want to lint for.
2. Write a dirty YAML file and the corresponding .exp file for that lint.
3. Run `go test ./test` to make sure the files are being found correctly.  You should fail the newly introduced tests.
4. Code until you pass the above tests.  Make sure you run `go test ./test` so that you know if you fail any previous tests as well.

## Contact

This project is maintained by @jjayeon and @hankc97, under the guidance of @aybabtme and @joshmgross.  Please contact any of us for more information.
