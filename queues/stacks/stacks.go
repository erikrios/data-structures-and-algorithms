package stacks

import (
	"errors"
	"fmt"
)

var (
	ErrEmptyStack = errors.New("stack is empty.")
)

type Stacks[T any] struct {
	elements []T
}

func NewStack[T any]() *Stacks[T] {
	return &Stacks[T]{}
}

func (s *Stacks[T]) String() string {
	return fmt.Sprintf("%v", s.elements)
}

func (s *Stacks[T]) Push(item T) {
	s.elements = append(s.elements, item)
}

func (s *Stacks[T]) Pop() (popped T, err error) {
	if s.IsEmpty() {
		err = ErrEmptyStack
		return
	}

	popped = s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return
}

func (s *Stacks[T]) Peek() (item T, err error) {
	if s.IsEmpty() {
		err = ErrEmptyStack
		return
	}

	item = s.elements[len(s.elements)-1]
	return
}

func (s *Stacks[T]) IsEmpty() bool {
	return len(s.elements) == 0
}
