package framework

import (
	"encoding/hex"
)

type EncodedValue struct {
	Value    string `json:"value"`
	Encoding string `json:"encoding"`
}

type Step struct {
	Action string        `json:"action"`
	Key    *EncodedValue `json:"key"`
	Data   *EncodedValue `json:"data,omitempty"`
}

type Test struct {
	Name  string       `json:"name"`
	Root  EncodedValue `json:"expected_root"`
	Steps []Step       `json:"steps"`
}

func (t Test) GetName() string {
	return t.Name
}

func HexValue(data []byte) EncodedValue {
	return EncodedValue{
		Value:    hex.EncodeToString(data),
		Encoding: "hex",
	}
}

func Utf8Value(data []byte) EncodedValue {
	return EncodedValue{
		Value:    string(data),
		Encoding: "utf-8",
	}
}
