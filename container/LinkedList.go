package container

type _LinkedNode[T any] struct {
	_next  *_LinkedNode[T]
	_prev  *_LinkedNode[T]
	_value T
}

type LinkedList[T any] struct {
	_head *_LinkedNode[T]
	_tail *_LinkedNode[T]
	_size int
}

func (this *LinkedList[T]) Size() int {
	return this._size
}

func (this *LinkedList[T]) IsEmpty() bool {
	return this.Size() == 0
}

func (this *LinkedList[T]) Clear() {
	this._head = nil
	this._tail = nil
	this._size = 0
}

func (this *LinkedList[T]) PushBack(value T) {
	if this.IsEmpty() {
		var firstNode *_LinkedNode[T] = &_LinkedNode[T]{
			_next:  nil,
			_prev:  nil,
			_value: value,
		}
		this._head = firstNode
		this._tail = firstNode
	} else {
		var newTail *_LinkedNode[T] = &_LinkedNode[T]{
			_next:  nil,
			_prev:  this._tail,
			_value: value,
		}
		this._tail._next = newTail
		this._tail = newTail
	}
	this._size++
}

func (this *LinkedList[T]) PopBack() T {
	var value T = this._tail._value

	if this.Size() == 1 {
		this._head = nil
		this._tail = nil
	} else {
		this._tail = this._tail._prev
		this._tail._next._prev = nil
		this._tail._next = nil
	}

	this._size--
	return value
}

func (this *LinkedList[T]) Back() T {
	var value T = this._tail._value
	return value
}

func (this *LinkedList[T]) PushFront(value T) {
	if this.IsEmpty() {
		var firstNode *_LinkedNode[T] = &_LinkedNode[T]{
			_next:  nil,
			_prev:  nil,
			_value: value,
		}
		this._head = firstNode
		this._tail = firstNode
	} else {
		var newHead *_LinkedNode[T] = &_LinkedNode[T]{
			_next:  this._head,
			_prev:  nil,
			_value: value,
		}
		this._head._prev = newHead
		this._head = newHead
	}
	this._size++
}

func (this *LinkedList[T]) PopFront() T {
	var value T = this._head._value

	if this.Size() == 1 {
		this._head = nil
		this._tail = nil
	} else {
		this._head = this._head._next
		this._head._prev._next = nil
		this._head._prev = nil
	}

	this._size--
	return value
}

func (this *LinkedList[T]) Front() T {
	var value T = this._head._value
	return value
}

func (this *LinkedList[T]) Remove(iterator *LinkedListIterator[T]) {
	if this != iterator._linked_list {
		return
	}

	if iterator._node == this._head {
		this.PopFront()
	} else if iterator._node == this._tail {
		this.PopBack()
	} else {
		iterator._node._prev._next = iterator._node._next
		iterator._node._next._prev = iterator._node._prev
		iterator._node._next = nil
		iterator._node._prev = nil
		this._size--
	}
}

func (this *LinkedList[T]) Iterator() Iterator[T] {
	if this.IsEmpty() {
		return nil
	}
	return &LinkedListIterator[T]{
		_linked_list: this,
		_node:        this._head,
	}
}

func (this *LinkedList[T]) Assign(iterator Iterator[T]) {
	for iterator != nil {
		this.PushBack(iterator.Value())
		iterator = iterator.Next()
	}
}

/*
LinkedListIterator adalah Iterator untuk LinkedList
*/
type LinkedListIterator[T any] struct {
	_linked_list *LinkedList[T]
	_node        *_LinkedNode[T]
}

func (this *LinkedListIterator[T]) Next() Iterator[T] {
	if this._node._next == nil {
		return nil
	}

	return &LinkedListIterator[T]{
		_linked_list: this._linked_list,
		_node:        this._node._next,
	}
}

func (this *LinkedListIterator[T]) Prev() Iterator[T] {
	if this._node._prev == nil {
		return nil
	}

	return &LinkedListIterator[T]{
		_linked_list: this._linked_list,
		_node:        this._node._next,
	}
}

func (this *LinkedListIterator[T]) Value() T {
	return this._node._value
}
