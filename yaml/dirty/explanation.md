This directory will be scanned for .yaml files.  For each .yaml file, we'll read the corresponding .exp(ected) file to see the expected output, then compare the two.

Please note that the tester is going to be VERY finicky --- the output string must match the .exp file EXACTLY.  Spaces instead of tabs?  Error.  Trailing space at the end of the line?  Error.  Some hacking may be necessary to get the .exp file and output to match.  That being said, I'd rather our tester throw too many errors than too few.
