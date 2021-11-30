package tcpserver

import (
	"encoding/json"
	"lsp/server/parse"
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
