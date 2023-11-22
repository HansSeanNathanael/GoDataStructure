package container

type Container interface {
	Size() int
	IsEmpty() bool
	Clear()
}
