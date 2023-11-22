package container

type BaseSet[T any] interface {
	Assignable[T]
	Container
	Iterable[T]

	Add(value T)
	Remove(value T)
	Exist(value T) bool
}
