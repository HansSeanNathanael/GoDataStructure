package container

type ExpandableQueue[T any] interface {
	ExpandableList[T]
	PushFront(value T)
	PopFront() T
	Front() T
}
