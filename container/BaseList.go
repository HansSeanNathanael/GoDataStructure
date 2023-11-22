package container

type BaseList[T any] interface {
	Assignable[T]
	Container
	ExpandableList[T]
	Iterable[T]
}
