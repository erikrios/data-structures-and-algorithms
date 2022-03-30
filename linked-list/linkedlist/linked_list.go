package linkedlist

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	ErrNoSuchElement = errors.New("no such element")
	ErrOutOfRange    = errors.New("out of range")
)

type linkedList[T any] struct {
	first *node[T]
	last  *node[T]
	size  int
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

	l.size++

	if l.isEmpty() {
		l.first, l.last = node, node
		return
	}

	node.next = l.first
	l.first = node
}

func (l *linkedList[T]) AddLast(v T) {
	node := newNode(v, nil)

	l.size++

	if l.isEmpty() {
		l.first, l.last = node, node
		return
	}

	l.last.next = node
	l.last = node
}

func (l *linkedList[T]) DeleteFirst() (err error) {
	if l.isEmpty() {
		err = ErrNoSuchElement
		return
	}

	temp := l.first
	l.first = l.first.next
	temp.next = nil
	l.size--
	return
}

func (l *linkedList[T]) DeleteLast() (err error) {
	if l.isEmpty() {
		err = ErrNoSuchElement
		return
	}

	l.size--
	if l.first == l.last {
		l.first, l.last = nil, nil
		return
	}

	prev := l.getPrevious(l.last)

	prev.next = nil
	l.last = prev
	return
}

func (l *linkedList[T]) Contains(v T) bool {
	return l.IndexOf(v) != -1
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

func (l *linkedList[T]) Size() int {
	return l.size
}

func (l *linkedList[T]) ToSlice() []T {
	results := make([]T, l.Size())

	node := l.first

	for i := 0; i < len(results); i++ {
		results[i] = node.val
		node = node.next
	}

	return results
}

func (l *linkedList[T]) Reverse() {
	if l.isEmpty() {
		return
	}

	prev := l.first
	cur := l.first.next

	for cur != nil {
		next := cur.next
		cur.next = prev
		prev = cur
		cur = next
	}

	l.last = l.first
	l.last.next = nil
	l.first = prev
}

func (l *linkedList[T]) GetKthFromTheEnd(k int) (val T, err error) {
	if l.isEmpty() {
		err = ErrOutOfRange
		return
	}

	a := l.first
	b := l.first

	for i := 0; i < k-1; i++ {
		b = b.next
		if b == nil {
			err = ErrOutOfRange
			return
		}
	}

	for b != l.last {
		a = a.next
		b = b.next
	}

	val = a.val
	return
}

func (l *linkedList[T]) PrintMiddle() {
	if l.isEmpty() {
		fmt.Println(ErrOutOfRange.Error())
		return
	}

	a := l.first
	b := l.first

	for b != l.last && b.next != l.last {
		b = b.next.next
		a = a.next
	}

	if b == l.last {
		fmt.Println(a.val)
	} else {
		fmt.Println(a.val, ",", a.next.val)
	}
}

func (l *linkedList[T]) isEmpty() bool {
	return l.first == nil
}

func (l *linkedList[T]) getPrevious(n *node[T]) *node[T] {
	if l.isEmpty() || l.first == l.last || n == l.first {
		return nil
	}

	prev := l.first

	for prev.next != n {
		prev = prev.next
	}

	return prev
}
