package marshalling

type MarshalFunc[T any] func(ts ...T) ([]byte, error)

type Marshaller[T any] struct {
	Func      MarshalFunc[T]
	Extension string
}

func (m Marshaller[T]) Marshal(ts ...T) ([]byte, error) {
	return m.Func(ts...)
}
