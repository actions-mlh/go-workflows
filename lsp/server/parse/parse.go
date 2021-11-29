package parse

import (
	"encoding/json"
)

type LspRequest struct {
	Header *LspHeader
	Body   *LspBody
}

type LspHeader struct {
	ContentLength int64
	ContentType   string
}
type LspBody struct {
	Jsonrpc string          `json:"jsonrpc"`
	Id      int             `json:"id"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
}
