package framework

import (
    "bytes"
    "encoding/json"

    "github.com/ghodss/yaml"
)

type MarshalFunc func([]Test) ([]byte, error)

// MarshalJson is an instance of a MarshalFunc for JSON
func MarshalJson(tests []Test) ([]byte, error) {
    data, err := json.Marshal(tests)
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
func MarshalYaml(tests []Test) ([]byte, error) {
    data, err := yaml.Marshal(tests)
    if err != nil {
        return nil, err
    }

    return data, nil
}

type Marshaller struct {
    Func MarshalFunc
}

func (m Marshaller) Marshal(tests []Test) ([]byte, error) {
    return m.Func(tests)
}

func NewJsonMarshaller() Marshaller {
    return Marshaller{
        Func: MarshalJson,
    }
}

func NewYamlMarshaller() Marshaller {
    return Marshaller{
        Func: MarshalYaml,
    }
}
