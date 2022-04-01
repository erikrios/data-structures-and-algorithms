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

type LinkedList[T any] struct {
	first *node[T]
	last  *node[T]
	size  int
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func CreateWithLoop() *LinkedList[int] {
	list := NewLinkedList[int]()
	list.AddLast(10)
	list.AddLast(20)
	list.AddLast(30)

	// Get a reference to 30
	node := list.last

	list.AddLast(40)
	list.AddLast(50)

	// Create the loop
	list.last.next = node

	return list
}

func (l *LinkedList[T]) String() (out string) {
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

func (l *LinkedList[T]) AddFirst(v T) {
	node := newNode(v, nil)

	l.size++

	if l.isEmpty() {
		l.first, l.last = node, node
		return
	}

	node.next = l.first
	l.first = node
}

func (l *LinkedList[T]) AddLast(v T) {
	node := newNode(v, nil)

	l.size++

	if l.isEmpty() {
		l.first, l.last = node, node
		return
	}

	l.last.next = node
	l.last = node
}

func (l *LinkedList[T]) DeleteFirst() (err error) {
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

func (l *LinkedList[T]) DeleteLast() (err error) {
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

func (l *LinkedList[T]) Contains(v T) bool {
	return l.IndexOf(v) != -1
}

func (l *LinkedList[T]) IndexOf(v T) (index int) {
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

func (l *LinkedList[T]) Size() int {
	return l.size
}

func (l *LinkedList[T]) Remove(v T) {
	if l.isEmpty() {
		return
	}

	for node := l.first; node != nil; node = node.next {
		if reflect.DeepEqual(v, node.val) {
			if l.first == l.last {
				l.first = nil
				l.last = nil
				l.size--
				return

			}

			if node == l.first {
				l.first = node.next
				node.next = nil
				l.size--
				return
			}

			prev := l.getPrevious(node)
			if node == l.last {
				l.last = prev
			}
			prev.next = node.next
			node.next = nil
			l.size--
			return
		}
	}
}

func (l *LinkedList[T]) ToSlice() []T {
	results := make([]T, l.Size())

	node := l.first

	for i := 0; i < len(results); i++ {
		results[i] = node.val
		node = node.next
	}

	return results
}

func (l *LinkedList[T]) Reverse() {
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

func (l *LinkedList[T]) GetKthFromTheEnd(k int) (val T, err error) {
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

func (l *LinkedList[T]) PrintMiddle() {
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

func (l *LinkedList[T]) HasLoop() bool {
	slow, fast := l.first, l.first

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next

		if slow == fast {
			return true
		}
	}

	return false
}

func (l *LinkedList[T]) isEmpty() bool {
	return l.first == nil
}

func (l *LinkedList[T]) getPrevious(n *node[T]) *node[T] {
	if l.isEmpty() || l.first == l.last || n == l.first {
		return nil
	}

	prev := l.first

	for prev.next != n {
		prev = prev.next
	}

	return prev
}
