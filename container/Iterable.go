package container

type Iterable[T any] interface {
	Iterator() Iterator[T]
}
