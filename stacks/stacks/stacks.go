package stacks

import (
	"errors"
	"fmt"
)

var (
	ErrEmptyStack = errors.New("stack is empty.")
)

type stack[T any] struct {
	elements []T
}

func NewStack[T any]() *stack[T] {
	return &stack[T]{}
}

func (s *stack[T]) String() string {
	return fmt.Sprintf("%v", s.elements)
}

func (s *stack[T]) Push(item T) {
	s.elements = append(s.elements, item)
}

func (s *stack[T]) Pop() (popped T, err error) {
	if s.IsEmpty() {
		err = ErrEmptyStack
		return
	}

	popped = s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return
}

func (s *stack[T]) Peek() (item T, err error) {
	if s.IsEmpty() {
		err = ErrEmptyStack
		return
	}

	item = s.elements[len(s.elements)-1]
	return
}

func (s *stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}
