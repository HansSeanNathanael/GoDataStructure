package container

type PriorityQueueFunction[T any] interface {
	Compare(v1 T, v2 T) bool
}

type PriorityQueue[T any, C PriorityQueueFunction[T]] struct {
	_container  ArrayList[T]
	_comparator C
}

func (this *PriorityQueue[T, C]) Size() int {
	return this._container.Size()
}

func (this *PriorityQueue[T, C]) IsEmpty() bool {
	return this._container.IsEmpty()
}

func (this *PriorityQueue[T, C]) Clear() {
	this._container.Clear()
}

func (this *PriorityQueue[T, C]) Top() T {
	return this._container.Get(0)
}

func (this *PriorityQueue[T, C]) Push(value T) {
	var index int = this._container.Size()
	this._container.PushBack(value)

	for index > 0 {
		var parent int = (index - 1) / 2
		if this._comparator.Compare(this._container.Get(index), this._container.Get(parent)) {
			var temp T = this._container.Get(index)
			this._container.Set(index, this._container.Get(parent))
			this._container.Set(parent, temp)
		} else {
			break
		}
		index = parent
	}
}

func (this *PriorityQueue[T, C]) Pop() T {
	var top T = this.Top()

	this._container.Set(0, this._container.Back())
	this._container.PopBack()

	var index int = 0
	for index < this._container.Size() {
		var child []int = []int{
			(index << 1) + 1,
			(index << 1) + 2,
		}
		var biggerIndex int = index

		for _, element := range child {
			if element < this._container.Size() && this._comparator.Compare(this._container.Get(element), this._container.Get(biggerIndex)) {
				biggerIndex = element
			}
		}

		if biggerIndex == index {
			break
		}
		var temp T = this._container.Get(index)
		this._container.Set(index, this._container.Get(biggerIndex))
		this._container.Set(biggerIndex, temp)

		index = biggerIndex
	}

	return top
}

func (this *PriorityQueue[T, C]) Assign(iterator Iterator[T]) {
	for iterator != nil {
		this.Push(iterator.Value())
		iterator = iterator.Next()
	}
}
