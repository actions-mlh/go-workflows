package tcpserver

import (
	"encoding/json"
	"github.com/actions-mlh/go-workflows/lsp/server/parse"
	"github.com/pkg/errors"
	"go.lsp.dev/protocol"
)

func Initialize(body *parse.LspBody) (*parse.InitializeResult, error) {
	params := body.Params
	initializeParamStruct := protocol.InitializeParams{}
	err := json.Unmarshal(params, &initializeParamStruct)
	if err != nil {
		return nil, errors.New("decoding lsp body params")
	}

	result, err := parse.NewInitializeResult(initializeParamStruct)

	if err != nil {
		return nil, errors.Wrap(err, "decoding initialized params")
	}

	return result, nil
}

func DidChange(body *parse.LspBody) (*parse.PublishDiagnosticsParams, error) {
	params := body.Params
	textDocument, diagnostics, err := parse.NewDidChangeDiagnostic(params)
	if err != nil {
		return nil, err
	}

	publishedDiagParams, err := parse.NewPublishDiagParams(diagnostics, textDocument)
	if err != nil {
		return nil, err
	}

	return publishedDiagParams, nil
}
