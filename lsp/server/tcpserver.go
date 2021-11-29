package tcpserver

import (
	"encoding/json"
	"lsp/server/parse"
	"fmt"
	"github.com/pkg/errors"
	"go.lsp.dev/protocol"
)

func Initialize(body *parse.LspBody) (*InitializeResult, error) {
	params := body.Params
	initializeParamStruct := protocol.InitializeParams{}
	err := json.Unmarshal(params, &initializeParamStruct)
	if err != nil {
		return nil, errors.New("decoding lsp body params")
	}

	fmt.Printf("%+v\n", initializeParamStruct)

	
	if err != nil {
		return nil, errors.Wrap(err, "decoding initialized params")
	}


	result := InitializeResult{
		Capabilities: ServerCapabilities{
			TextDocumentSync: 1,
			CompletionProvider: ResolveProviderValue{
				ResolveProvider: true,
			},
			Workspace: WorkspaceValue{
				WorkspaceFolders: WorkspaceFoldersValue{
					Supported: true,
				},
			},
		},
		ServerInfo: ServerInfoValue{
			Name: "vscode",
			Version: "1.62.3",
		},
	}

	return &result, nil
}

type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities"`
	ServerInfo   ServerInfoValue    `json:"serverinfo"`
}

type ServerCapabilities struct {
	TextDocumentSync   int                  `json:"textDocumentSync"`
	CompletionProvider ResolveProviderValue `json:"completionProvider"`
	Workspace          WorkspaceValue       `json:"workspace"`
}

type TextDocumentSyncValue struct {
	Kind int `json:"kind"`
}

type ResolveProviderValue struct {
	ResolveProvider bool `json:"resolveProvider"`
}

type WorkspaceValue struct {
	WorkspaceFolders WorkspaceFoldersValue `json:"workspaceFolders"`
}

type WorkspaceFoldersValue struct {
	Supported bool `json:"supported"`
}

type ServerInfoValue struct {
	Name string `json:"name"`
	Version string `json:"version"`
}

// requestBody := Resp{
// 	Jsonrpc: "2.0",
// 	Id:      0,
// 	Result: ResultValue{
// 		Capabilities: CapabilitiesValue{
// 			TextDocumentSync: 2,
// 			CompletionProvider: ResolveProviderValue{
// 				ResolveProvider: true,
// 			},
// 			Workspace: WorkspaceValue{
// 				WorkspaceFolders: WorkspaceFoldersValue{
// 					Supported: true,
// 				},
// 			},
// 		},
// 	},
// }
