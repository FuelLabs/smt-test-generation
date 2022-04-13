package framework

import (
	"bytes"
	"encoding/json"

	"github.com/ghodss/yaml"
)

type MarshalFunc[T any] func(ts ...T) ([]byte, error)

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
	if len(ts) == 1 {
		t := ts[0]
		data, err := json.Marshal(t)
		if err != nil {
			return nil, err
		}
		return formatJson(data)
	} else {
		combined := make([]T, 0)
		for _, t := range ts {
			combined = append(combined, t)
		}
		data, err := json.Marshal(combined)
		if err != nil {
			return nil, err
		}
		return formatJson(data)
	}
}

// MarshalYaml is an instance of a MarshalFunc for YAML
func MarshalYaml[T any](ts ...T) ([]byte, error) {
	if len(ts) == 1 {
		t := ts[0]
		data, err := yaml.Marshal(t)
		if err != nil {
			return nil, err
		}
		return data, nil
	} else {
		combined := make([]T, 0)
		for _, t := range ts {
			combined = append(combined, t)
		}
		data, err := yaml.Marshal(combined)
		if err != nil {
			return nil, err
		}
		return data, nil
	}
}

type Marshaller[T any] struct {
	Func      MarshalFunc[T]
	Extension string
}

func (m Marshaller[T]) Marshal(ts ...T) ([]byte, error) {
	return m.Func(ts...)
}

func NewJsonMarshaller[T any]() Marshaller[T] {
	return Marshaller[T]{
		Func:      MarshalJson[T],
		Extension: "json",
	}
}

func NewYamlMarshaller[T any]() Marshaller[T] {
	return Marshaller[T]{
		Func:      MarshalYaml[T],
		Extension: "yaml",
	}
}
