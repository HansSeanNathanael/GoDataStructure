package container

type ExpandableTree[T any] interface {
	Add(value T)
	Remove(value T) T
	Find(value T) T
}
