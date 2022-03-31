package queues

import "fmt"

type circularArrayQueue[T any] struct {
	items []T
	front int
	rear  int
	count int
}

func NewCircularArrayQueue[T any](size int) *circularArrayQueue[T] {
	items := make([]T, size)
	return &circularArrayQueue[T]{items: items}
}

func (a *circularArrayQueue[T]) String() string {
	return fmt.Sprintf("%v\nfront: %d\nrear: %d\ncount: %d\n", a.items, a.front, a.rear, a.count)
}

func (a *circularArrayQueue[T]) Enqueue(item T) error {
	if a.IsFull() {
		return ErrQueueFull
	}

	a.items[a.rear] = item
	a.rear = (a.rear + 1) % len(a.items)
	a.count++

	return nil
}

func (a *circularArrayQueue[T]) Dequeue() (item T, err error) {
	item, err = a.Peek()
	if err != nil {
		err = err
		return
	}

	var defValue T
	a.items[a.front] = defValue
	a.front = (a.front + 1) % len(a.items)
	a.count--
	return
}

func (a *circularArrayQueue[T]) Peek() (item T, err error) {
	if a.IsEmpty() {
		err = ErrQueueEmpty
		return
	}

	item = a.items[a.front]
	return
}

func (a *circularArrayQueue[T]) IsEmpty() bool {
	return a.count == 0
}

func (a *circularArrayQueue[T]) IsFull() bool {
	return a.count == len(a.items)
}
