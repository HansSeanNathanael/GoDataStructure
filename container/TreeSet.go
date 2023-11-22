package container

type TreeSetFunction[T any] interface {
	AVLTreeFunction[T]
}

type TreeSet[T any, Func TreeSetFunction[T]] struct {
	BaseSet[T]
	tree AVLTree[T, Func]
}

func CreateTreeSet[T any, Func TreeSetFunction[T]](multi bool) TreeSet[T, Func] {
	return TreeSet[T, Func]{
		tree: CreateAVLTree[T, Func](multi),
	}
}

func (this *TreeSet[T, F]) Size() int {
	return this.tree.Size()
}

func (this *TreeSet[T, F]) IsEmpty() bool {
	return this.tree.IsEmpty()
}

func (this *TreeSet[T, F]) Clear() {
	this.tree.Clear()
}

func (this *TreeSet[T, F]) Add(value T) {
	this.tree.Add(value)
}

func (this *TreeSet[T, F]) Remove(value T) {
	this.tree.Remove(value)
}

func (this *TreeSet[T, F]) Exist(value T) bool {
	var val *T = this.tree.Find(value)
	return val != nil
}

func (this *TreeSet[T, F]) Iterator() Iterator[T] {
	return this.tree.Iterator()
}

func (this *TreeSet[T, F]) Assign(iterator Iterator[T]) {
	var it = iterator
	for it != nil {
		this.Add(it.Value())
		it = it.Next()
	}
}

/*
TreeSetIterator[T] adalah Iterator[T] untuk TreeSet
*/
type TreeSetIterator[T any] struct {
	AVLTreeIterator[T]
}
