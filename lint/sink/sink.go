package sink

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
)

type ProblemSink struct {
	Filename string
	Problems []string
}

func (sink *ProblemSink) Record(raw *yaml.Node, format string, args ...interface{}) {
	sink.Problems = append(sink.Problems,
		fmt.Sprintf("%s:%d:%d\terror:\t", sink.Filename, raw.Line, raw.Column)+
			fmt.Sprintf(format, args...),
	)
}

func (sink *ProblemSink) Render(w io.Writer) {
	for _, problem := range sink.Problems {
		fmt.Fprint(w, problem)
	}
}
