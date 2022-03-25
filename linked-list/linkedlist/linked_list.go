package linkedlist

import (
	"fmt"
	"reflect"
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

func (l *linkedList[T]) DeleteFirst() {
	if l.first == nil {
		return
	}

	l.first = l.first.next
}

func (l *linkedList[T]) DeleteLast() {
	if l.last == nil {
		return
	}

	if l.first == l.last {
		l.first = nil
		l.last = nil
		return
	}

	prev := l.first

	for prev.next != l.last {
		prev = prev.next
	}

	prev.next = nil
	l.last = prev
}

func (l *linkedList[T]) Contains(v T) bool {
	for node := l.first; node != nil; node = node.next {
		if reflect.DeepEqual(node.val, v) {
			return true
		}
	}

	return false
}

func (l *linkedList[T]) IndexOf(v T) (index int) {
	index = -1

	i := 0
	for node := l.first; node != nil; node = node.next {
		if reflect.DeepEqual(node.val, v) {
			index = i
			return
		}

		i++
	}

	return index
}
