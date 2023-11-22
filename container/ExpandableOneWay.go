package container

type ExpandableOneWay[T any] interface {
	Top() T
	Push(value T)
	Pop() T
}
