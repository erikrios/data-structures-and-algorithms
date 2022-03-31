package queues

import "fmt"

type arrayQueue[T any] struct {
	items []T
	front int
	rear  int
}

func NewArrayQueue[T any](size int) *arrayQueue[T] {
	items := make([]T, size)
	return &arrayQueue[T]{items: items, front: -1, rear: -1}
}

func (a *arrayQueue[T]) String() string {
	return fmt.Sprintf("%v\nfront: %d\nrear: %d\n", a.items, a.front, a.rear)
}

func (a *arrayQueue[T]) Enqueue(item T) error {
	if a.IsFull() {
		return ErrQueueFull
	}

	if a.front > 0 && a.rear == len(a.items)-1 {
		for i := 0; i <= a.rear-a.front; i++ {
			a.items[i] = a.items[i+1]
		}

		a.rear, a.front = a.rear-a.front, 0
		a.rear++
		a.items[a.rear] = item
	} else {
		if a.IsEmpty() {
			a.front++
		}
		a.rear++
		a.items[a.rear] = item
	}

	return nil
}

func (a *arrayQueue[T]) Dequeue() (item T, err error) { return }

func (a *arrayQueue[T]) Peek() (item T, err error) {
	if a.IsEmpty() {
		err = ErrQueueEmpty
		return
	}

	item = a.items[a.front]
	return
}

func (a *arrayQueue[T]) IsEmpty() bool {
	return a.front == -1 && a.rear == -1
}

func (a *arrayQueue[T]) IsFull() bool {
	return a.front == 0 && a.rear == len(a.items)-1
}
