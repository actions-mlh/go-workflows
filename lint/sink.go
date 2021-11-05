package lint

import (
	"fmt"
	"io"
	"gopkg.in/yaml.v3"
)

type problemSink struct {
	Filename string
	Problems []string
}

func (sink *problemSink) record(raw *yaml.Node, format string, args ...interface{}) {
	sink.Problems = append(sink.Problems,
		fmt.Sprintf("%s:%d:%d\terror:\t", sink.Filename, raw.Line, raw.Column) + 
		fmt.Sprintf(format, args...),
	)
}

func (sink *problemSink) recordMultiple(raw *yaml.Node, format string, args ...interface{}) {
	for _, node := range raw.Content {
		sink.record(node, format, args)
	}
}

func (sink *problemSink) render(w io.Writer) {
	for _, problem := range sink.Problems {
		fmt.Fprint(w, problem)
	}
}

