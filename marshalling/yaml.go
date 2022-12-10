package marshalling

import (
	"github.com/ghodss/yaml"
)

func marshalYaml[T any](ts ...T) ([]byte, error) {
	if len(ts) == 1 {
		t := ts[0]
		return yaml.Marshal(t)
	} else {
		combined := make([]T, 0)
		for _, t := range ts {
			combined = append(combined, t)
		}
		return yaml.Marshal(combined)
	}
}

// MarshalYaml is an instance of a MarshalFunc for YAML
func MarshalYaml[T any](ts ...T) ([]byte, error) {
	return marshalYaml(ts...)
}

// NewYamlMarshaller creates an instance of Marshaller that produces YAML
func NewYamlMarshaller[T any]() Marshaller[T] {
	return Marshaller[T]{
		Func:      MarshalYaml[T],
		Extension: "yaml",
	}
}
