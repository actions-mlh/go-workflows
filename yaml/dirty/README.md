## explanation

This directory will be scanned for .yaml files.  For each .yaml file, we'll read the corresponding .exp(ected) file to see the expected output, then compare the two.  Any expected error not found will be reported, and any found error that is not expected will be reported as well.

For future maintainers, a recommended development procedure is as follows:
1. Identify a new class of error you want to lint for.
2. Write a dirty YAML file and the corresponding .exp file for that lint.
3. Run `go test ./test` to make sure the files are being found correctly.  You should fail the newly introduced tests.
4. Code until you pass the above tests.  Make sure you run `go test ./test` so that you know if you fail any previous tests as well.

Please note that the tester is going to be VERY finicky --- the output string must match the .exp file EXACTLY.  Spaces instead of tabs?  Error.  Trailing space at the end of the line?  Error.  Some hacking may be necessary to get the .exp file and output to match.
