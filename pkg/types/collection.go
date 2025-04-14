package types

type Collection[T any] struct {
	Items map[string]T `json:"items"`
}

func NewCollection[T any]() *Collection[T] {
	return &Collection[T]{
		Items: map[string]T{},
	}
}
