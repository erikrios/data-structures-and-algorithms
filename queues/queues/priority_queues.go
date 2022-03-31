package queues

import "fmt"

type priorityQueue[T int | float64] struct {
	items []T
}

func NewPriorityQueue[T int | float64]() *priorityQueue[T] {
	return &priorityQueue[T]{}
}

func (p *priorityQueue[T]) String() string {
	return fmt.Sprintf("%v\n", p.items)
}

func (p *priorityQueue[T]) Enqueue(item T) error {
	p.items = append(p.items, item)

	for i := len(p.items) - 1; i > 0; i-- {
		if p.items[i] < p.items[i-1] {
			p.items[i], p.items[i-1] = p.items[i-1], p.items[i]
			continue
		}
		return nil
	}
	return nil
}

func (p *priorityQueue[T]) Dequeue() (item T, err error) {
	item, err = p.Peek()
	if err != nil {
		return
	}

	newItems := make([]T, len(p.items)-1)
	for i := 1; i < len(p.items); i++ {
		newItems[i-1] = p.items[i]
	}
	p.items = newItems

	return
}

func (p *priorityQueue[T]) Peek() (item T, err error) {
	if p.IsEmpty() {
		err = ErrQueueEmpty
		return
	}

	item = p.items[0]
	return
}

func (p *priorityQueue[T]) IsEmpty() bool {
	return len(p.items) == 0
}

func (p *priorityQueue[T]) IsFull() bool {
	return false
}
