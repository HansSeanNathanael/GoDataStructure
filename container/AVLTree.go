package container

type avlTreeNode[T any] struct {
	depth  int
	value  T
	parent *avlTreeNode[T]
	left   *avlTreeNode[T]
	right  *avlTreeNode[T]

	next *avlTreeNode[T]
	prev *avlTreeNode[T]
}

type AVLTreeFunction[T any] interface {
	Compare(v1 T, v2 T) int
}

func (this *avlTreeNode[T]) update_depth() {
	var leftDepth = 0
	var rightDepth = 0

	if this.left != nil {
		leftDepth = this.left.depth + 1
	}
	if this.right != nil {
		rightDepth = this.right.depth + 1
	}
	this.depth = max(leftDepth, rightDepth)
}

func (this *avlTreeNode[T]) rotate_left() *avlTreeNode[T] {
	var child *avlTreeNode[T] = this.right

	this.right = child.left
	if this.right != nil {
		this.right.parent = this
	}

	child.left = this
	this.parent = child

	this.update_depth()
	child.update_depth()

	return child
}

func (this *avlTreeNode[T]) rotate_right() *avlTreeNode[T] {
	var child *avlTreeNode[T] = this.left

	this.left = child.right
	if this.left != nil {
		this.left.parent = this
	}

	child.right = this
	this.parent = child

	this.update_depth()
	child.update_depth()

	return child
}

type AVLTree[T any, C AVLTreeFunction[T]] struct {
	root          *avlTreeNode[T]
	head          *avlTreeNode[T]
	comparator    C
	support_multi bool
	size          int
}

func CreateAVLTree[T any, C AVLTreeFunction[T]](multi bool) AVLTree[T, C] {
	return AVLTree[T, C]{
		support_multi: multi,
	}
}

func (this *AVLTree[T, C]) fix(node *avlTreeNode[T]) {
	for node != nil {
		node.update_depth()

		var leftDepth = 0
		var rightDepth = 0
		if node.left != nil {
			leftDepth = node.left.depth + 1
		}
		if node.right != nil {
			rightDepth = node.right.depth + 1
		}

		var leftBigger bool = (leftDepth - rightDepth) > 1
		var rightBigger bool = (rightDepth - leftDepth) > 1
		var childLeftDepth = 0
		var childRightDepth = 0
		if leftBigger {
			var child *avlTreeNode[T] = node.left
			if child.left != nil {
				childLeftDepth = child.left.depth + 1
			}
			if child.right != nil {
				childRightDepth = child.right.depth + 1
			}
			if (childRightDepth - childLeftDepth) > 0 {
				var newChild *avlTreeNode[T] = child.rotate_left()
				newChild.parent = node
				node.left = newChild
			}
			var parent *avlTreeNode[T] = node.parent
			var direction = 1
			if parent != nil && parent.left == node {
				direction = -1
			}

			node = node.rotate_right()
			node.parent = parent

			if parent == nil {
				this.root = node
			} else if direction == -1 {
				parent.left = node
			} else {
				parent.right = node
			}
		} else if rightBigger {
			var child *avlTreeNode[T] = node.right
			if child.left != nil {
				childLeftDepth = child.left.depth + 1
			}
			if child.right != nil {
				childRightDepth = child.right.depth + 1
			}
			if (childLeftDepth - childRightDepth) > 0 {
				var newChild *avlTreeNode[T] = child.rotate_right()
				newChild.parent = node
				node.right = newChild
			}
			var parent *avlTreeNode[T] = node.parent
			var direction = 1
			if parent != nil && parent.left == node {
				direction = -1
			}

			node = node.rotate_left()
			node.parent = parent

			if parent == nil {
				this.root = node
			} else if direction == -1 {
				parent.left = node
			} else {
				parent.right = node
			}
		}

		node = node.parent
	}
}

func (this *AVLTree[T, C]) Size() int {
	return this.size
}

func (this *AVLTree[T, C]) IsEmpty() bool {
	return this.Size() == 0
}

func (this *AVLTree[T, C]) Clear() {
	this.root = nil
}

func (this *AVLTree[T, C]) Add(value T) {
	if this.root == nil {
		this.root = &avlTreeNode[T]{
			depth:  0,
			value:  value,
			parent: nil,
			left:   nil,
			right:  nil,
			next:   nil,
			prev:   nil,
		}
		this.head = this.root
		this.size++
		return
	}

	var walker *avlTreeNode[T] = this.root
	var diff int = 0
	for {
		diff = this.comparator.Compare(value, walker.value)
		if diff < 0 {
			if walker.left == nil {
				break
			}
			walker = walker.left
		} else if diff >= 0 {
			if diff == 0 && !this.support_multi {
				return
			}
			if walker.right == nil {
				diff = 1
				break
			}
			walker = walker.right
		}
	}

	var newNode *avlTreeNode[T] = &avlTreeNode[T]{
		depth:  0,
		value:  value,
		parent: walker,
		left:   nil,
		right:  nil,
		next:   nil,
		prev:   nil,
	}

	if diff < 0 {
		walker.left = newNode

		newNode.prev = walker.prev
		if newNode.prev != nil {
			newNode.prev.next = newNode
		} else {
			this.head = newNode
		}

		walker.prev = newNode
		newNode.next = walker
	} else {
		walker.right = newNode

		newNode.next = walker.next
		if newNode.next != nil {
			newNode.next.prev = newNode
		}
		walker.next = newNode
		newNode.prev = walker
	}

	var node *avlTreeNode[T] = newNode
	this.fix(node)
	this.size++
}

func (this *AVLTree[T, C]) Remove(value T) {
	if this.IsEmpty() {
		return
	}
	if this.Size() == 1 {
		this.root = nil
		this.head = nil
		this.size--
		return
	}

	var walker *avlTreeNode[T] = this.root
	var diff int = 0
	for walker != nil {
		diff = this.comparator.Compare(value, walker.value)
		if diff < 0 {
			walker = walker.left
		} else if diff > 0 {
			walker = walker.right
		} else {
			break
		}
	}

	if walker == nil {
		return
	}

	var update *avlTreeNode[T] = nil
	if walker.right == nil && walker.left == nil {
		update = walker.parent
		walker.parent = nil
		if update.left == walker {
			update.left = nil
			update.prev = walker.prev
			if update.prev != nil {
				update.prev.next = update
			} else {
				this.head = update
			}
		} else {
			update.right = nil
			update.next = walker.next
			if update.next != nil {
				update.next.prev = update
			}
		}
	} else if walker.right == nil {
		var left = walker.left
		left.parent = walker.parent

		if walker.parent == nil {
			this.root = left
		} else if walker.parent.left == walker {
			walker.parent.left = left
		} else {
			walker.parent.right = left
		}
		update = left

		left.next = walker.next
		if left.next != nil {
			left.next.prev = left
		}
	} else {
		var next = walker.next
		if next.parent != walker {
			next.parent.left = next.right
			if next.right != nil {
				next.right.parent = next.parent
				update = next.right
			} else {
				update = next.parent
			}
		} else {
			walker.right = nil
			update = next
		}

		next.left = walker.left
		next.right = walker.right
		if next.left != nil {
			next.left.parent = next
		}
		if next.right != nil {
			next.right.parent = next
		}

		next.prev = walker.prev
		if next.prev == nil {
			this.head = next
		} else {
			next.prev.next = next
		}

		next.parent = walker.parent
		if walker.parent == nil {
			this.root = next
		} else if walker.parent.left == walker {
			walker.parent.left = next
		} else {
			walker.parent.right = next
		}
	}

	this.fix(update)
	this.size--
}

func (this *AVLTree[T, C]) Find(value T) *T {
	if this.IsEmpty() {
		return nil
	}

	var walker *avlTreeNode[T] = this.root
	var diff int = 0
	for walker != nil {
		diff = this.comparator.Compare(value, walker.value)
		if diff < 0 {
			walker = walker.left
		} else if diff > 0 {
			walker = walker.right
		} else {
			break
		}
	}

	if walker == nil {
		return nil
	}
	return &walker.value
}

func (this *AVLTree[T, C]) Iterator() Iterator[T] {
	if this.head == nil {
		return nil
	}

	return &AVLTreeIterator[T]{
		node: this.head,
	}
}

/*
AVLTreeIterator[T] adalah Iterator[T] untuk AVLTree
*/
type AVLTreeIterator[T any] struct {
	node *avlTreeNode[T]
}

func (this *AVLTreeIterator[T]) Next() Iterator[T] {
	if this.node.next != nil {
		return &AVLTreeIterator[T]{
			node: this.node.next,
		}
	}
	return nil
}

func (this *AVLTreeIterator[T]) Prev() Iterator[T] {
	if this.node.prev != nil {
		return &AVLTreeIterator[T]{
			node: this.node.prev,
		}
	}
	return nil
}

func (this *AVLTreeIterator[T]) Value() T {
	return this.node.value
}
