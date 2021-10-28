package lint

import (
	"fmt"
	"io"
	"gopkg.in/yaml.v3"
)

type ProblemSink struct {
	Filename string
	Output   io.Writer
	Problems []string
}

func (sink *ProblemSink) Record(node *yaml.Node, format string, args ...interface{}) {
	sink.Problems = append(sink.Problems,
		fmt.Sprintf("%s:%d:%d\n", sink.Filename, node.Line, node.Column) +
		fmt.Sprintf("\t"+format+" \n", args...),
	)
}

func (sink *ProblemSink) Render() {
	for _, problem := range sink.Problems {
		fmt.Fprint(sink.Output, problem)
	}
}
