package container

type TreeMapComparator[T any] interface {
	Compare(v1 T, v2 T) int
}

type TreeMapFunction[Key any, Value any, F TreeMapComparator[Key]] struct {
	AVLTreeFunction[Key]
	helper F
}

func (this TreeMapFunction[Key, Value, helper]) Compare(v1 MapPair[Key, Value], v2 MapPair[Key, Value]) int {
	return this.helper.Compare(v1.Key, v2.Key)
}

type MapPair[Key any, Value any] struct {
	Key   Key
	Value Value
}

type TreeMap[Key any, Value any, C TreeMapComparator[Key]] struct {
	BaseSet[Key]
	tree AVLTree[MapPair[Key, Value], TreeMapFunction[Key, Value, C]]
}

func CreateTreeMap[Key any, Value any, C TreeMapComparator[Key]](multi bool) TreeMap[Key, Value, C] {
	return TreeMap[Key, Value, C]{
		tree: CreateAVLTree[MapPair[Key, Value], TreeMapFunction[Key, Value, C]](multi),
	}
}

func (this *TreeMap[Key, Value, C]) Size() int {
	return this.tree.Size()
}

func (this *TreeMap[Key, Value, C]) IsEmpty() bool {
	return this.tree.IsEmpty()
}

func (this *TreeMap[Key, Value, C]) Clear() {
	this.tree.Clear()
}

func (this *TreeMap[Key, Value, C]) Add(key Key, value Value) {
	this.tree.Add(MapPair[Key, Value]{
		Key:   key,
		Value: value,
	})
}

func (this *TreeMap[Key, Value, C]) Remove(key Key) {
	this.tree.Remove(MapPair[Key, Value]{
		Key: key,
	})
}

func (this *TreeMap[Key, Value, C]) Get(key Key) *Value {
	var val *MapPair[Key, Value] = this.tree.Find(MapPair[Key, Value]{
		Key: key,
	})

	if val == nil {
		return nil
	}

	return &val.Value
}

func (this *TreeMap[Key, Value, C]) Iterator() Iterator[MapPair[Key, Value]] {
	return this.tree.Iterator()
}

/*
TreeMapIterator[Key, Value] adalah Iterator[Key, Value] untuk TreeMap
*/
type TreeMapIterator[Key any, Value any] struct {
	AVLTreeIterator[MapPair[Key, Value]]
}
