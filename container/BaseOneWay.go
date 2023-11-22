package container

type BaseOneWay[T any] interface {
	Assignable[T]
	Container
	ExpandableOneWay[T]
}
