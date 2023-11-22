package container

// type _RedBlackTreeNode[T any] struct {
// 	_color  int8 // 0 means black, 1 means red
// 	_value  T
// 	_parent *_RedBlackTreeNode[T]
// 	_left   *_RedBlackTreeNode[T]
// 	_right  *_RedBlackTreeNode[T]

// 	_next *_RedBlackTreeNode[T]
// 	_prev *_RedBlackTreeNode[T]
// }

// func (this *_RedBlackTreeNode[T]) _color_of_parent() int8 {
// 	return this._parent._color
// }

// func (this *_RedBlackTreeNode[T]) _color_of_uncle() int8 {
// 	if this._parent != this._parent._parent._left {
// 		if this._parent._parent._left == nil {
// 			return 0
// 		}
// 		return this._parent._parent._left._color
// 	}
// 	if this._parent._parent._right == nil {
// 		return 0
// 	}
// 	return this._parent._parent._right._color
// }

// func (this *_RedBlackTreeNode[T]) _left_or_right() int8 {
// 	if this._parent._left == this {
// 		return -1
// 	}
// 	return 1
// }

// // Attention!
// // For RedBlackTree doing what you intended, you must read important message below!
// //
// // RedBlackTreeFunction[T] used for comparison in insert, delete, and find of node of RedBlackTree[T]
// // v1 is the inserted value, and v2 is the node on tree.
// // The result of Compare should be (v1 - v2), it means negative if v1 is "less" than v2, positive if
// // v1 is "more" than v2, and 0 if v1 and v2 is equal.
// type RedBlackTreeFunction[T any] interface {
// 	Compare(v1 T, v2 T) int
// }

// // Attention!
// // Don't use this struct bare, instead make a wrapping struct!
// // _support_multi is a flag telling if the tree supporting multiple equal data
// // to be saved on the tree.
// type RedBlackTree[T any, C RedBlackTreeFunction[T]] struct {
// 	_root          *_RedBlackTreeNode[T]
// 	_head          *_RedBlackTreeNode[T]
// 	_comparator    C
// 	_support_multi bool
// 	_size          int
// }

// func (this *RedBlackTree[T, C]) _rotate_node_left(node *_RedBlackTreeNode[T]) {
// 	var parent *_RedBlackTreeNode[T] = node._parent
// 	var child *_RedBlackTreeNode[T] = node._right

// 	node._right = child._left
// 	if node._right != nil {
// 		node._right._parent = node
// 	}

// 	child._left = node
// 	node._parent = child
// 	child._parent = parent

// 	if parent == nil {
// 		this._root = child
// 	} else if parent._left == node {
// 		parent._left = child
// 	} else {
// 		parent._right = child
// 	}
// }

// func (this *RedBlackTree[T, C]) _rotate_node_right(node *_RedBlackTreeNode[T]) {
// 	var parent *_RedBlackTreeNode[T] = node._parent
// 	var child *_RedBlackTreeNode[T] = node._left

// 	node._left = child._right
// 	if node._left != nil {
// 		node._left._parent = node
// 	}

// 	child._right = node
// 	node._parent = child
// 	child._parent = parent

// 	if parent == nil {
// 		this._root = child
// 	} else if parent._left == node {
// 		parent._left = child
// 	} else {
// 		parent._right = child
// 	}
// }

// func (this *RedBlackTree[T, C]) _fix_node(node *_RedBlackTreeNode[T]) *_RedBlackTreeNode[T] {
// 	if node._color == 0 {
// 		return nil
// 	}
// 	if node._parent == nil {
// 		node._color = 0
// 		return nil
// 	}
// 	if node._color_of_parent() == 0 {
// 		return nil
// 	}

// 	var colorOfUncle int8 = node._color_of_uncle()
// 	var locParent int8 = node._parent._left_or_right()
// 	var loc int8 = node._left_or_right()

// 	node._parent._parent._color = 1
// 	if colorOfUncle == 1 {
// 		node._parent._parent._left._color = 0
// 		node._parent._parent._right._color = 0
// 		return node._parent._parent
// 	}

// 	var grandParent *_RedBlackTreeNode[T] = node._parent._parent
// 	if locParent != loc {
// 		node._color = 0
// 		if loc == 1 {
// 			this._rotate_node_left(node._parent)
// 		} else {
// 			this._rotate_node_right(node._parent)
// 		}
// 	} else {
// 		node._parent._color = 0
// 	}

// 	if locParent == 1 {
// 		this._rotate_node_left(grandParent)
// 	} else {
// 		this._rotate_node_right(grandParent)
// 	}

// 	return nil
// }

// // func (this *RedBlackTree[T, C]) _transplant(n1 *_RedBlackTreeNode[T], n2 *_RedBlackTreeNode[T]) {
// // 	if n1._parent == nil {
// // 		this._root = n2
// // 	} else if n1 == n1._parent._left {
// // 		n1._parent._left = n2
// // 	} else {
// // 		n1._parent._right = n2
// // 	}

// // 	n2._parent = n1._parent
// // }

// func (this *RedBlackTree[T, C]) Size() int {
// 	return this._size
// }

// func (this *RedBlackTree[T, C]) IsEmpty() bool {
// 	return this.Size() == 0
// }

// func (this *RedBlackTree[T, C]) Clear() {
// 	this._root = nil
// }

// func (this *RedBlackTree[T, C]) Add(value T) {
// 	if this._root == nil {
// 		this._root = &_RedBlackTreeNode[T]{
// 			_color:  0,
// 			_value:  value,
// 			_parent: nil,
// 			_left:   nil,
// 			_right:  nil,
// 			_next:   nil,
// 			_prev:   nil,
// 		}
// 		this._head = this._root
// 		this._size++
// 		return
// 	}

// 	var walker *_RedBlackTreeNode[T] = this._root
// 	var diff int = 0
// 	for {
// 		diff = this._comparator.Compare(value, walker._value)
// 		if diff < 0 {
// 			if walker._left == nil {
// 				break
// 			}
// 			walker = walker._left
// 		} else if diff >= 0 {
// 			if diff == 0 && !this._support_multi {
// 				return
// 			}
// 			if walker._right == nil {
// 				diff = 1
// 				break
// 			}
// 			walker = walker._right
// 		}
// 	}

// 	var newNode *_RedBlackTreeNode[T] = &_RedBlackTreeNode[T]{
// 		_color:  1,
// 		_value:  value,
// 		_parent: walker,
// 		_left:   nil,
// 		_right:  nil,
// 		_next:   nil,
// 		_prev:   nil,
// 	}
// 	if diff < 0 {
// 		walker._left = newNode

// 		newNode._prev = walker._prev
// 		if newNode._prev != nil {
// 			newNode._prev._next = newNode
// 		} else {
// 			this._head = newNode
// 		}

// 		walker._prev = newNode
// 		newNode._next = walker
// 	} else {
// 		walker._right = newNode

// 		newNode._next = walker._next
// 		if newNode._next != nil {
// 			newNode._next._prev = newNode
// 		}
// 		walker._next = newNode
// 		newNode._prev = walker
// 	}

// 	var node *_RedBlackTreeNode[T] = newNode
// 	for node != nil {
// 		node = this._fix_node(node)
// 	}
// 	this._size++
// }

// // func (this *RedBlackTree[T, C]) Remove(value T) T {

// // }

// // func (this *RedBlackTree[T, C]) Find(value T) T

// func (this *RedBlackTree[T, C]) Iterator() Iterator[T] {
// 	return &RedBlackTreeIterator[T, C]{
// 		_tree: this,
// 		_node: *this._head,
// 	}
// }

// type RedBlackTreeIterator[T any, C RedBlackTreeFunction[T]] struct {
// 	_tree *RedBlackTree[T, C]
// 	_node _RedBlackTreeNode[T]
// }

// func (this *RedBlackTreeIterator[T, C]) Next() Iterator[T] {
// 	if this._node._next == nil {
// 		return nil
// 	}
// 	return &RedBlackTreeIterator[T, C]{
// 		_tree: this._tree,
// 		_node: *this._node._next,
// 	}
// }

// func (this *RedBlackTreeIterator[T, C]) Prev() Iterator[T] {
// 	if this._node._prev == nil {
// 		return nil
// 	}
// 	return &RedBlackTreeIterator[T, C]{
// 		_tree: this._tree,
// 		_node: *this._node._prev,
// 	}
// }

// func (this *RedBlackTreeIterator[T, C]) Value() T {
// 	return this._node._value
// }
