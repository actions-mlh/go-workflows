package parse

import (
	"github.com/actions-mlh/go-workflows/lint"
)

func NewPublishDiagParams(currentWorkSpace string, fileInfo *FileInfo, textDocument *TextDocumentValue) (*PublishDiagnosticsParams, error) {
	diagnostics, err := NewDiagnostics(currentWorkSpace, fileInfo)
	if err != nil {
		return nil, err
	}
	// publish diagnostic parameters
	publishedDiagnosticParams := PublishDiagnosticsParams{}
	// uri
	publishedDiagnosticParams.URI = textDocument.Uri
	// version
	publishedDiagnosticParams.Version = textDocument.Version
	// diagnostics
	publishedDiagnosticParams.Diagnostics = *diagnostics

	return &publishedDiagnosticParams, nil
}

type PublishDiagnosticsParams struct {
	URI         string       `json:"uri"`
	Version     int          `json:"version"`
	Diagnostics []Diagnostic `json:"diagnostics"`
}

func NewDiagnostics(currentWorkSpace string, fileInfo *FileInfo) (*[]Diagnostic, error) {
	problemText := []byte(fileInfo.Text)
	problems, err := lint.Lint(currentWorkSpace, problemText)
	if err != nil {
		return nil, err
	}

	var diagnostics []Diagnostic
	for _, problem := range problems {
		diagnostic := Diagnostic{}
		// severity
		diagnostic.Severity = 2
		// range { Start, End }
		diagnostic.Range.Start.Line = problem.Range.Line - 1
		diagnostic.Range.Start.Character = problem.Range.Column - 1
		diagnostic.Range.End.Line = problem.Range.Line
		diagnostic.Range.End.Character = 0
		// message
		diagnostic.Message = problem.ProblemMsg
		// source
		diagnostic.Source = "mlh-golang-parser"

		diagnostics = append(diagnostics, diagnostic)
	}

	return &diagnostics, nil
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
