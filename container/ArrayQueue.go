package container

type ArrayQueue[T any] struct {
	storage []T
	filled  int
	head    int
}

func (this *ArrayQueue[T]) Size() int {
	return this.filled
}

func (this *ArrayQueue[T]) IsEmpty() bool {
	return this.Size() == 0
}

func (this *ArrayQueue[T]) Clear() {
	this.storage = this.storage[:0]
	this.filled = 0
	this.head = 0
}

func (this *ArrayQueue[T]) PushBack(value T) {
	if this.Size() >= cap(this.storage) {
		var backup []T = this.storage
		var newArray []T
		newArray = append(newArray, backup[this.head:]...)
		newArray = append(newArray, backup[:this.head]...)
		newArray = append(newArray, value)
		this.storage = newArray

		this.head = 0
	} else {
		this.storage[(this.head+this.Size())%cap(this.storage)] = value
	}
	this.filled++
}

func (this *ArrayQueue[T]) PopBack() T {
	var value T = this.storage[(this.head+this.Size())%cap(this.storage)]
	this.filled--

	return value
}

func (this *ArrayQueue[T]) Back() T {
	var value T = this.storage[(this.head+this.Size())%cap(this.storage)]
	return value
}

func (this *ArrayQueue[T]) PushFront(value T) {
	if this.Size() >= cap(this.storage) {
		var backup []T = this.storage
		var newArray []T
		newArray = append(newArray, value)
		newArray = append(newArray, backup[this.head:]...)
		newArray = append(newArray, backup[:this.head]...)
		this.storage = newArray

		this.head = 0
	} else {
		var capacity int = cap(this.storage)
		this.head = (this.head - 1 + capacity) % capacity
		this.storage[this.head] = value
	}
	this.filled++
}

func (this *ArrayQueue[T]) PopFront() T {
	var value T = this.storage[this.head]
	this.head = (this.head + 1) % cap(this.storage)
	this.filled--

	return value
}

func (this *ArrayQueue[T]) Front() T {
	var value T = this.storage[this.head]
	return value
}

func (this *ArrayQueue[T]) Iterator() Iterator[T] {
	if this.IsEmpty() {
		return nil
	}

	return &ArrayQueueIterator[T]{
		arr:     this,
		current: 0,
	}
}

/*
ArrayQueueIterator[T] adalah Iterator[T] untuk ArrayQueue[T]
*/
type ArrayQueueIterator[T any] struct {
	arr     *ArrayQueue[T]
	current int
}

func (this *ArrayQueueIterator[T]) Next() Iterator[T] {
	if this.current >= this.arr.filled {
		return nil
	}
	return &ArrayQueueIterator[T]{
		arr:     this.arr,
		current: this.current + 1,
	}
}

func (this *ArrayQueueIterator[T]) Prev() Iterator[T] {
	if this.current <= 0 {
		return nil
	}

	return &ArrayQueueIterator[T]{
		arr:     this.arr,
		current: this.current - 1,
	}
}

func (this *ArrayQueueIterator[T]) Value() T {
	var value T = this.arr.storage[(this.arr.head+this.current)%cap(this.arr.storage)]

	return value
}
