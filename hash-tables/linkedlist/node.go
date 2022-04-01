package linkedlist

type node[T any] struct {
	val  T
	next *node[T]
}

func newNode[T any](val T, next *node[T]) *node[T] {
	return &node[T]{val: val, next: next}
}
