package container

type LinkedStack[T any] struct {
	_stack LinkedList[T]
}

func (this *LinkedStack[T]) Size() int {
	return this._stack.Size()
}

func (this *LinkedStack[T]) IsEmpty() bool {
	return this.Size() == 0
}

func (this *LinkedStack[T]) Clear() {
	this._stack.Clear()
}

func (this *LinkedStack[T]) Top() T {
	return this._stack.Back()
}

func (this *LinkedStack[T]) Push(value T) {
	this._stack.PushBack(value)
}

func (this *LinkedStack[T]) Pop() T {
	return this._stack.PopBack()
}

func (this *LinkedStack[T]) Assign(iterator Iterator[T]) {
	this._stack.Assign(iterator)
}
