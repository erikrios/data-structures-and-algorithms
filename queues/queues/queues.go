package queues

import "errors"

type Queue[T any] interface {
	Enqueue(item T) error
	Dequeue() (item T, err error)
	Peek() (item T, err error)
	IsEmpty() bool
	IsFull() bool
}

var (
	ErrQueueEmpty = errors.New("queue is empty")
	ErrQueueFull  = errors.New("queue is full")
)
