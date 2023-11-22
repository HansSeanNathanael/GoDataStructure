package container

type ExpandableList[T any] interface {
	PushBack(value T)
	PopBack() T
	Back() T
}
