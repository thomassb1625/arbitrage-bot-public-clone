// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    response, err := UnmarshalResponse(bytes)
//    bytes, err = response.Marshal()

package arbitrage

import "encoding/json"

func UnmarshalResponse(data []byte) (Response, error) {
	var r Response
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Response) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Response struct {
	Name      string   `json:"name"`
	Timestamp string   `json:"timestamp"`
	Version   Version  `json:"version"`
	Tags      Tags     `json:"tags"`
	LogoURI   string   `json:"logoURI"`
	Keywords  []string `json:"keywords"`
	Tokens    []Token  `json:"tokens"`
}

type Tags struct {
}

type Token struct {
	ChainID    int64       `json:"chainId"`
	Address    string      `json:"address"`
	Name       string      `json:"name"`
	Symbol     string      `json:"symbol"`
	Decimals   int64       `json:"decimals"`
	LogoURI    string      `json:"logoURI"`
	Extensions *Extensions `json:"extensions,omitempty"`
}

type Extensions struct {
	BridgeInfo map[string]BridgeInfo `json:"bridgeInfo"`
}

type BridgeInfo struct {
	TokenAddress string `json:"tokenAddress"`
}

type Version struct {
	Major int64 `json:"major"`
	Minor int64 `json:"minor"`
	Patch int64 `json:"patch"`
}
