package queues

import "github.com/erikrios/queues/stacks"

type queueWithTwoStacks[T any] struct {
	firstStack  stacks.Stacks[T]
	secondStack stacks.Stacks[T]
}

func NewQueueWithTwoStacks[T any]() *queueWithTwoStacks[T] {
	return &queueWithTwoStacks[T]{}
}

func (q *queueWithTwoStacks[T]) Enqueue(item T) error {
	q.firstStack.Push(item)
	return nil
}

func (q *queueWithTwoStacks[T]) Dequeue() (item T, err error) {
	item, err = q.Peek()
	if err != nil {
		return
	}

	item, err = q.secondStack.Pop()
	return
}

func (q *queueWithTwoStacks[T]) Peek() (item T, err error) {
	if q.IsEmpty() {
		err = ErrQueueEmpty
		return
	}

	if q.secondStack.IsEmpty() {
		q.fillStack()
	}

	item, err = q.secondStack.Peek()
	return
}

func (q *queueWithTwoStacks[T]) IsEmpty() bool {
	return q.firstStack.IsEmpty() && q.secondStack.IsEmpty()
}

func (q *queueWithTwoStacks[T]) IsFull() bool {
	return false
}

func (q *queueWithTwoStacks[T]) fillStack() {
	for item, err := q.firstStack.Pop(); err == nil; item, err = q.firstStack.Pop() {
		q.secondStack.Push(item)
	}
}
