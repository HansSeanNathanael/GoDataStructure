package container

type Assignable[T any] interface {
	Assign(iterator Iterator[T])
}
