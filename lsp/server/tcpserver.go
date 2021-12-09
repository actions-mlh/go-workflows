package tcpserver

import (
	"encoding/json"
	"github.com/pkg/errors"
	"go.lsp.dev/protocol"
	"github.com/actions-mlh/go-workflows/lsp/server/parse"
)

var (
	currentWorkspace string
)

func Initialize(body *parse.LspBody) (*parse.InitializeResult, error) {
	params := body.Params
	initializeParamStruct := protocol.InitializeParams{}
	err := json.Unmarshal(params, &initializeParamStruct)
	if err != nil {
		return nil, errors.New("decoding lsp body params")
	}

	currentWorkspace = initializeParamStruct.WorkspaceFolders[0].Name
	result, err := parse.NewInitializeResult(initializeParamStruct)

	if err != nil {
		return nil, errors.Wrap(err, "decoding initialized params")
	}

	return result, nil
}

func DidOpen(body *parse.LspBody) (*parse.PublishDiagnosticsParams, error) {
	params := body.Params
	fileInfo, textDocument, err := parse.NewDidOpenFileInfo(params)
	if err != nil {
		return nil, err
	}

	publishedDiagParams, err := parse.NewPublishDiagParams(currentWorkspace, fileInfo, textDocument)
	if err != nil {
		return nil, err
	}
	return publishedDiagParams, nil
}

func DidChange(body *parse.LspBody) (*parse.PublishDiagnosticsParams, error) {
	params := body.Params
	fileInfo, textDocument, err := parse.NewDidChangeFileInfo(params)
	if err != nil {
		return nil, err
	}

	publishedDiagParams, err := parse.NewPublishDiagParams(currentWorkspace, fileInfo, textDocument)
	if err != nil {
		return nil, err
	}
	return publishedDiagParams, nil
}

