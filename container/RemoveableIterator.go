package container

type RemoveableIterator[T any] interface {
	Remove(iterator Iterator[T])
}
