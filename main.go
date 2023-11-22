package main

import (
	"DataStructure/container"
	"fmt"
)

type Comparator struct{}

func (Comparator) Compare(v1 int, v2 int) bool {
	return v1 < v2
}

type Comparator2 struct{}

func (Comparator2) Compare(v1 int, v2 int) int {
	return v1 - v2
}

type Comparator3 struct{}

func (Comparator3) Compare(v1 int, v2 int) int {
	return v1 - v2
}

func main() {

	var tree container.AVLTree[int, Comparator2] = container.AVLTree[int, Comparator2]{}
	for i := 0; i < 25; i++ {
		tree.Add(i)
	}
	for i := 0; i < 25; i++ {
		tree.Remove(i)
	}

	var treeMap container.TreeMap[int, int, Comparator3] = container.CreateTreeMap[int, int, Comparator3](true)
	for i := 0; i < 25; i++ {
		treeMap.Add(i, i)
	}
	for i := 0; i < 25; i++ {
		treeMap.Remove(i)
	}

	for it := treeMap.Iterator(); it != nil; it = it.Next() {
		fmt.Println(it.Value().Key, it.Value().Value)
	}
}
