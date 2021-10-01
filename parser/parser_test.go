package parser

import (
	"os"
	"strings"
	"strconv"
	"testing"
)

func TestParse(t *testing.T) {
	dir, err := os.Open("./yaml/")
	if err != nil {
		panic("error reading dir yaml")
	}
	defer dir.Close()

	files, err := dir.Readdirnames(0)
	if err != nil {
		panic("error reading files from yaml")
	}

	for _, file := range files {
		data, err := os.ReadFile("./yaml/" + file)
		if err != nil {
			panic("error reading test file " + file)
		}

		// separate string into lines
		lines := strings.Split(string(data), "\n")
		// get first line and split on : to get just the number in the second half
		s := strings.Split(lines[0], ":")[1]
		// cut the first char to remove space
		s = s[1:]
		// convert to int
		lineNoExpected, err := strconv.Atoi(s)
		if err != nil {
			panic("error getting line no. from yaml file " + file)
		}
		passedExpected := lineNoExpected == -1
		
		passedActual, lineNoActual, err := Parse(data)

		if (passedExpected != passedActual) {
			t.Errorf(
				"error linting file %v: expected %v, got %v.",
				file, passedExpected, passedActual,
			)
		} else if (lineNoExpected != lineNoActual) {
			t.Errorf(
				"error linting file %v: typo is on line %v, but returned line %v.\n" +
					"line %v: %v\n" +
					"line %v: %v",
				file, lineNoExpected, lineNoActual,
				lineNoExpected, lines[lineNoExpected-1],
				lineNoActual, lines[lineNoActual-1],
				// line numbers are 1-indexed in common usage
			)
		}
	}
}
