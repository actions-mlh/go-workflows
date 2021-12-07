package parse

import (
	"encoding/json"
	"github.com/actions-mlh/go-workflows/lint"
)

func NewDidChangeDiagnostic(params json.RawMessage) (*TextDocumentValue, *[]Diagnostic, error) {
	didChangeRequestStruct := DidChangeRequest{}
	err := json.Unmarshal(params, &didChangeRequestStruct)
	if err != nil {
		return nil, nil, err
	}

	fileInfoStruct := FileInfo{}
	for _, fileInfo := range didChangeRequestStruct.ContentChanges {
		err := json.Unmarshal(fileInfo, &fileInfoStruct)
		if err != nil {
			return nil, nil, err
		}
	}

	problems, err := lint.Lint("hanks_workspace.yaml", []byte(fileInfoStruct.Text))
	if err != nil {
		return nil, nil, err
	}
	
	var diagnostics []Diagnostic
	for _, problem := range problems {
		diagnostic := Diagnostic{}
		// severity
		diagnostic.Severity = 2
		// range { Start, End }
		diagnostic.Range.Start.Line = problem.Range.Line - 1
		diagnostic.Range.Start.Character = problem.Range.Column
		diagnostic.Range.End.Line = problem.Range.Line
		diagnostic.Range.End.Character = 0
		// message
		diagnostic.Message = problem.ProblemMsg
		// source
		diagnostic.Source = "typescript"

		diagnostics = append(diagnostics, diagnostic)
	}

	return &didChangeRequestStruct.TextDocument, &diagnostics, nil
}

type DidChangeRequest struct {
	TextDocument   TextDocumentValue `json:"textDocument"`
	ContentChanges []json.RawMessage `json:"contentChanges"`
}

type TextDocumentValue struct {
	Uri     string `json:"uri"`
	Version int    `json:"version"`
}

type FileInfo struct {
	Range       rangeValue `json:"range"`
	RangeLength int        `json:"rangeLength"`
	Text        string     `json:"text"`
}

type rangeValue struct {
	Start startValue `json:"start"`
	End   endValue   `json:"end"`
}

type startValue struct {
	Line      int `json:"line"`
	Character int `json:"character"`
}

type endValue struct {
	Line      int `json:"line"`
	Character int `json:"character"`
}

type Diagnostic struct {
	Severity int        `json:"severity"`
	Range    RangeValue `json:"range"`
	Message  string     `json:"message"`
	Source   string     `json:"source"`
}

type RangeValue struct {
	Start StartPosition `json:"start"`
	End   EndPosition   `json:"end"`
}

type StartPosition struct {
	Line      int `json:"line"`
	Character int `json:"character"`
}

type EndPosition struct {
	Line      int `json:"line"`
	Character int `json:"character"`
}
