package parse

import (
	"go.lsp.dev/protocol"
)

var (
	hasConfigurationCapability                = false
	hasWorkspaceFolderCapability              = false
	hasDiagnosticRelatedInformationCapability = false
)

func NewInitializeResult(initializeParamStruct protocol.InitializeParams) (*InitializeResult, error) {
	result := InitializeResult{}

	// Server Information
	result.ServerInfo.Name = initializeParamStruct.ClientInfo.Name
	result.ServerInfo.Version = initializeParamStruct.ClientInfo.Version

	// Text Document Synchronization
	/**
	 * Defines how text documents are synced. Is either a detailed structure
	 * defining each notification or for backwards compatibility the
	 * TextDocumentSyncKind number. If omitted it defaults to
	 * `TextDocumentSyncKind.None`.
	 */
	// Non-Detailed
	result.Capabilities.TextDocumentSync = 1

	// Completion Request
	result.Capabilities.CompletionProvider.ResolveProvider = true

	// Workspace Folder Capability
	if (initializeParamStruct.Capabilities.Workspace != nil) && initializeParamStruct.Capabilities.Workspace.WorkspaceFolders {
		hasWorkspaceFolderCapability = true
	}

	if hasWorkspaceFolderCapability {
		result.Capabilities.Workspace = WorkspaceValue{
			WorkspaceFolders: WorkspaceFoldersValue{
				Supported: true,
			},
		}
	}

	return &result, nil
}

type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities"`
	ServerInfo   ServerInfoValue    `json:"serverinfo"`
}

type ServerCapabilities struct {
	TextDocumentSync         int                    `json:"textDocumentSync"`
	CompletionProvider       CompletionOptions      `json:"completionProvider"`
	Workspace                WorkspaceValue         `json:"workspace"`
}

type TextDocumentSyncValue struct {
	TextDocumentSyncKind int `json:"textDocumentSyncKind"`
}

type CompletionOptions struct {
	ResolveProvider bool `json:"resolveProvider"`
}

type WorkspaceValue struct {
	WorkspaceFolders WorkspaceFoldersValue `json:"workspaceFolders"`
}

type WorkspaceFoldersValue struct {
	Supported bool `json:"supported"`
}

type ServerInfoValue struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}
