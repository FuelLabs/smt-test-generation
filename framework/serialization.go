package framework

import (
	"bytes"
	"encoding/json"

	"github.com/ghodss/yaml"
)

type MarshalFunc[T any] func(ts ...T) ([]byte, error)

// MarshalJson is an instance of a MarshalFunc for JSON
func MarshalJson[T any](ts ...T) ([]byte, error) {
	combined := make([]T, 0)
	for _, t := range ts {
		combined = append(combined, t)
	}

	data, err := json.Marshal(combined)
	if err != nil {
		return nil, err
	}

	formatted := &bytes.Buffer{}
	err = json.Indent(formatted, data, "", "    ")
	if err != nil {
		return nil, err
	}

	return formatted.Bytes(), nil
}

// MarshalYaml is an instance of a MarshalFunc for YAML
func MarshalYaml[T any](ts ...T) ([]byte, error) {
	if len(ts) == 1 {
		data, err := yaml.Marshal(ts[0])
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
	Func MarshalFunc[T]
}

func (m Marshaller[T]) Marshal(ts ...T) ([]byte, error) {
	return m.Func(ts...)
}

func NewJsonMarshaller[T any]() Marshaller[T] {
	return Marshaller[T]{
		Func: MarshalJson[T],
	}
}

func NewYamlMarshaller[T any]() Marshaller[T] {
	return Marshaller[T]{
		Func: MarshalYaml[T],
	}
}
