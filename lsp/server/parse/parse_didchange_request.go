package parse

import (
	"encoding/json"
	"fmt"

	"github.com/actions-mlh/go-workflows/lint"
)

func NewDidChangeRequest(params json.RawMessage) (*DidChangeRequest, error) {
	didChangeRequestStruct := DidChangeRequest{}
	err := json.Unmarshal(params, &didChangeRequestStruct)
	if err != nil {
		return nil, err
	}

	fileInfoStruct := FileInfo{}
	for _, fileInfo := range didChangeRequestStruct.ContentChanges {
		err := json.Unmarshal(fileInfo, &fileInfoStruct)
		if err != nil {
			return nil, err
		}
	}

	problems, err := lint.Lint("hanks workspace", []byte(fileInfoStruct.Text))
	if err != nil {
		return nil, err
	}

	for _, problem := range problems {
		fmt.Println(problem)
	}

	return nil, nil
}

type DidChangeRequest struct {
	TextDocument   textDocumentValue `json:"textDocument"`
	ContentChanges []json.RawMessage `json:"contentChanges"`
}

type textDocumentValue struct {
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
