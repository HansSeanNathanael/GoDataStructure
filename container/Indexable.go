package container

type Indexable[K any, T any] interface {
	Get(key K) T
	Set(key K, value T)
}
