package container

type Iterator[T any] interface {
	Next() Iterator[T]
	Prev() Iterator[T]
	Value() T
}
