package container

type ArrayStack[T any] struct {
	stack ArrayList[T]
}

func (this *ArrayStack[T]) Size() int {
	return this.stack.Size()
}

func (this *ArrayStack[T]) IsEmpty() bool {
	return this.Size() == 0
}

func (this *ArrayStack[T]) Clear() {
	this.stack.Clear()
}

func (this *ArrayStack[T]) Top() T {
	return this.stack.Back()
}

func (this *ArrayStack[T]) Push(value T) {
	this.stack.PushBack(value)
}

func (this *ArrayStack[T]) Pop() T {
	return this.stack.PopBack()
}

func (this *ArrayStack[T]) Assign(iterator Iterator[T]) {
	this.stack.Assign(iterator)
}
