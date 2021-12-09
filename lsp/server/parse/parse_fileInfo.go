package parse

import (
	"encoding/json"
)

func NewDidOpenFileInfo(params json.RawMessage) (*FileInfo, *TextDocumentValue, error) {
	didOpenNotifStruct := DidOpenNotif{}
	err := json.Unmarshal(params, &didOpenNotifStruct)
	if err != nil {
		return nil, nil, err
	}

	problemText := []byte(didOpenNotifStruct.TextDocument.Text)
	fileInfoStruct := FileInfo{}
	err = json.Unmarshal(problemText, &fileInfoStruct)
	if err != nil {
		return nil, nil, err
	}

	return &fileInfoStruct, &didOpenNotifStruct.TextDocument, nil
}

func NewDidChangeFileInfo(params json.RawMessage) (*FileInfo, *TextDocumentValue, error) {
	didChangeNotifStruct := DidChangeNotif{}
	err := json.Unmarshal(params, &didChangeNotifStruct)
	if err != nil {
		return nil, nil, err
	}

	fileInfoStruct := FileInfo{}
	for _, fileInfo := range didChangeNotifStruct.ContentChanges {
		err := json.Unmarshal(fileInfo, &fileInfoStruct)
		if err != nil {
			return nil, nil, err
		}
	}

	return &fileInfoStruct, &didChangeNotifStruct.TextDocument, nil
}

type DidOpenNotif struct {
	TextDocument TextDocumentValue `json:"textDocument"`
}

type DidChangeNotif struct {
	TextDocument   TextDocumentValue `json:"textDocument"`
	ContentChanges []json.RawMessage `json:"contentChanges"`
}

type TextDocumentValue struct {
	Uri        string `json:"uri"`
	Version    int    `json:"version"`
	LanguageId string `json:"languageId"`
	Text       string `json:"text"`
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
