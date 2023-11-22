package container

type ArrayList[T any] struct {
	storage []T
}

func (this *ArrayList[T]) Size() int {
	return len(this.storage)
}

func (this *ArrayList[T]) IsEmpty() bool {
	return this.Size() == 0
}

func (this *ArrayList[T]) Clear() {
	this.storage = this.storage[:0]
}

func (this *ArrayList[T]) PushBack(value T) {
	this.storage = append(this.storage, value)
}

func (this *ArrayList[T]) PopBack() T {
	var value T = this.Get(this.Size() - 1)
	this.storage = this.storage[:this.Size()-1]

	return value
}

func (this *ArrayList[T]) Back() T {
	var value T = this.Get(this.Size() - 1)

	return value
}

func (this *ArrayList[T]) Get(key int) T {
	return this.storage[key]
}

func (this *ArrayList[T]) Set(key int, value T) {
	this.storage[key] = value
}

func (this *ArrayList[T]) Iterator() Iterator[T] {
	return &ArrayListIterator[T]{
		arr:   this,
		index: 0,
	}
}

func (this *ArrayList[T]) Assign(iterator Iterator[T]) {
	for iterator != nil {
		this.PushBack(iterator.Value())
		iterator = iterator.Next()
	}
}

/*
ArrayListIterator adalah iterator untuk ArrayList
*/
type ArrayListIterator[T any] struct {
	arr   *ArrayList[T]
	index int
}

func (this *ArrayListIterator[T]) Next() Iterator[T] {
	if this.index >= this.arr.Size()-1 {
		return nil
	}

	return &ArrayListIterator[T]{
		arr:   this.arr,
		index: this.index + 1,
	}
}

func (this *ArrayListIterator[T]) Prev() Iterator[T] {
	if this.index <= 0 {
		return nil
	}

	return &ArrayListIterator[T]{
		arr:   this.arr,
		index: this.index - 1,
	}
}
func (this *ArrayListIterator[T]) Value() T {
	return this.arr.Get(this.index)
}
