package stacks

import (
	"errors"
	"fmt"
)

var (
	ErrStackOverflow = errors.New("stack is overflow")
	ErrStackEmpty    = errors.New("stack is empty")
)

type twoStack[T any] struct {
	items []T
	top1  int
	top2  int
}

func NewTwoStack[T any](size int) *twoStack[T] {
	items := make([]T, size)
	return &twoStack[T]{items: items, top1: -1, top2: size}
}

func (t *twoStack[T]) Push1(item T) error {
	if t.IsFull1() {
		return ErrStackOverflow
	}

	t.top1++
	t.items[t.top1] = item
	return nil
}

func (t *twoStack[T]) Pop1() (popped T, err error) {
	if t.IsEmpty1() {
		err = ErrStackEmpty
		return
	}

	popped = t.items[t.top1]
	t.top1--
	return
}

func (t *twoStack[T]) IsFull1() bool {
	return t.top1+1 == t.top2
}

func (t *twoStack[T]) IsEmpty1() bool {
	return t.top1 == -1
}

func (t *twoStack[T]) Push2(item T) error {
	if t.IsFull2() {
		return ErrStackOverflow
	}

	t.top2--
	t.items[t.top2] = item
	return nil
}

func (t *twoStack[T]) Pop2() (popped T, err error) {
	if t.IsEmpty2() {
		err = ErrStackEmpty
		return
	}

	popped = t.items[t.top2]
	t.top2++
	return
}

func (t *twoStack[T]) IsFull2() bool {
	return t.top2-1 == t.top1
}

func (t *twoStack[T]) IsEmpty2() bool {
	return t.top2 == len(t.items)
}

func (t *twoStack[T]) String() string {
	return fmt.Sprintf("%v", t.items)
}
