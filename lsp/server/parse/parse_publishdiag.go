package parse

func NewPublishDiagParams(diagnostics *[]Diagnostic, textDocument *TextDocumentValue) (*PublishDiagnosticsParams, error) {
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
