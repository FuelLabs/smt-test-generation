package marshalling

import (
	"bytes"
	"encoding/json"
)

func marshalJson[T any](ts ...T) ([]byte, error) {
	if len(ts) == 1 {
		t := ts[0]
		return json.Marshal(t)
	} else {
		combined := make([]T, 0)
		for _, t := range ts {
			combined = append(combined, t)
		}
		return json.Marshal(combined)
	}
}

func formatJson(data []byte) ([]byte, error) {
	formatted := &bytes.Buffer{}
	err := json.Indent(formatted, data, "", "    ")
	if err != nil {
		return nil, err
	}
	return formatted.Bytes(), nil
}

// MarshalJson is an instance of a MarshalFunc for JSON
func MarshalJson[T any](ts ...T) ([]byte, error) {
	data, err := marshalJson(ts...)
	if err != nil {
		return nil, err
	}
	data, err = formatJson(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// NewJsonMarshaller creates an instance of Marshaller that produces YAML
func NewJsonMarshaller[T any]() Marshaller[T] {
	return Marshaller[T]{
		Func:      MarshalJson[T],
		Extension: "json",
	}
}
