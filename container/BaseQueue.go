package container

type BaseQueue[T any] interface {
	Assignable[T]
	Container
	ExpandableQueue[T]
	Iterable[T]
}
