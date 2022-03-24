package linkedlist

import (
	"fmt"
)

type linkedList[T any] struct {
	first *node[T]
	last  *node[T]
}

func NewLinkedList[T any]() *linkedList[T] {
	return &linkedList[T]{}
}

func (l *linkedList[T]) String() (out string) {
	node := l.first

	for node != nil {
		if node == l.first && node == l.last {
			out += fmt.Sprintf("first-last:%v->", node.val)
		} else if node == l.first {
			out += fmt.Sprintf("first:%v->", node.val)
		} else if node == l.last {
			out += fmt.Sprintf("last:%v->", node.val)
		} else {
			out += fmt.Sprintf("%v->", node.val)
		}
		node = node.next
	}

	return
}

func (l *linkedList[T]) AddFirst(v T) {
	node := newNode(v, nil)

	if l.first == nil {
		l.first = node
		l.last = node
		return
	}

	node.next = l.first
	l.first = node
}

func (l *linkedList[T]) AddLast(v T) {
	node := newNode(v, nil)

	if l.last == nil {
		l.first = node
		l.last = node
		return
	}

	l.last.next = node
	l.last = node
}
