package lint

import (
	"fmt"
	"io"
	"gopkg.in/yaml.v3"
)

type ProblemSink struct {
	Filename string
	Problems []string
}

func (sink *ProblemSink) Record(raw *yaml.Node, format string, args ...interface{}) {
	sink.Problems = append(sink.Problems,
		fmt.Sprintf("%s:%d:%d\terror:\t", sink.Filename, raw.Line, raw.Column) + 
		fmt.Sprintf(format, args...),
	)
}

func (sink *ProblemSink) RecordMultiple(raw *yaml.Node, format string, args ...interface{}) {
	for _, node := range raw.Content {
		sink.Record(node, format, args)
	}
}

func (sink *ProblemSink) Render(w io.Writer) {
	for _, problem := range sink.Problems {
		fmt.Fprint(w, problem)
	}
}

