package heaps

import (
	"errors"
	"fmt"
)

type Heap struct {
	items []int
	size  int
}

func NewHeap() *Heap {
	items := make([]int, 10)
	return &Heap{items: items}
}

func (h *Heap) Insert(value int) error {
	if h.isFull() {
		return errors.New("The items is full")
	}

	h.items[h.size] = value
	h.size++

	h.bubbleUp()
	return nil
}

func (h *Heap) Remove() {

}

func (h *Heap) swap(first, second int) {
	h.items[first], h.items[second] = h.items[second], h.items[first]
}

func (h *Heap) parent(index int) int {
	return (index - 1) / 2
}

func (h *Heap) bubbleUp() {
	index := h.size - 1

	for index > 0 && h.items[index] > h.items[h.parent(index)] {
		h.swap(index, h.parent(index))
		index = h.parent(index)
	}
}

func (h *Heap) isFull() bool {
	return h.size == len(h.items)
}

func (h *Heap) String() string {
	return fmt.Sprintf("%v\n", h.items)
}
