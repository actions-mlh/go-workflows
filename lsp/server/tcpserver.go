package tcpserver

import (
	"encoding/json"
	"github.com/actions-mlh/go-workflows/lsp/server/parse"
	"github.com/pkg/errors"
	"go.lsp.dev/protocol"
	"fmt"
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

func DidChange(body *parse.LspBody) (error) {
	params := body.Params
	didChangeRequest, err := parse.NewDidChangeRequest(params)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", didChangeRequest)

	return nil
}