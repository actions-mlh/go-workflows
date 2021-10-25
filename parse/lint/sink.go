package lint

import (
	"fmt"
	"io"
	"gopkg.in/yaml.v3"
)

type ProblemSink struct {
	Filename string
	Output   io.Writer
}

func (sink *ProblemSink) RecordProblem(node *yaml.Node, format string, args ...interface{}) {
	fmt.Fprintf(sink.Output, "%s:%d:%d\n", sink.Filename, node.Line, node.Column)
	fmt.Fprintf(sink.Output, "\t"+format+" \n", args...)
}