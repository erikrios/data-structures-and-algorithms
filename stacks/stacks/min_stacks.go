package stacks

type minimumStack[T int] struct {
	stack    *stack[int]
	minStack *stack[int]
}

func NewMinimumStack[T int]() *minimumStack[T] {
	return &minimumStack[T]{stack: NewStack[int](), minStack: NewStack[int]()}
}

func (m *minimumStack[T]) Push(item int) {
	m.stack.Push(item)

	min, _ := m.minStack.Peek()
	if m.minStack.IsEmpty() {
		m.minStack.Push(item)
	} else if item < min {
		m.minStack.Push(item)
	}
}

func (m *minimumStack[T]) Pop() error {
	if m.stack.IsEmpty() {
		return ErrEmptyStack
	}

	top, err := m.stack.Pop()
	if err != nil {
		return err
	}

	min, err := m.minStack.Peek()
	if err != nil {
		return err
	}

	if min == top {
		m.minStack.Pop()
	}

	return nil
}

func (m *minimumStack[T]) Min() int {
	min, _ := m.minStack.Peek()
	return min
}
