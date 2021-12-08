package sink

import (
	"fmt"
	"io"

	"gopkg.in/yaml.v3"
)

type ProblemSink struct {
	Filename string
	Problems []ProblemsValue
}

type ProblemsValue struct {
	Range   RangeValue
	ProblemMsg string
}

type RangeValue struct {
	Line   int
	Column int
}

func (sink *ProblemSink) Record(raw *yaml.Node, format string, args ...interface{}) {
	problem := fmt.Sprintf("%s:%d:%d\terror:\t", sink.Filename, raw.Line, raw.Column) + fmt.Sprintf(format, args...)
	rangeValue := RangeValue{}

	rangeValue.Line = raw.Line
	rangeValue.Column = raw.Column

	problemsValue := ProblemsValue{}
	problemsValue.Range = rangeValue
	problemsValue.ProblemMsg = problem

	sink.Problems = append(sink.Problems, problemsValue)
}

func (sink *ProblemSink) Render(w io.Writer) {
	for _, problem := range sink.Problems {
		fmt.Fprint(w, problem.ProblemMsg)
	}
}
