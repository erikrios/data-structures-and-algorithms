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

func (h *Heap) Remove() (int, error) {
	if h.isEmpty() {
		return 0, errors.New("The items is empty")
	}

	root := h.items[0]
	h.items[0] = h.items[h.size-1]
	h.size--

	h.bubbleDown()

	return root, nil
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

func (h *Heap) bubbleDown() {
	index := 0
	for index <= h.size && !h.isValidParent(index) {
		largerChildIndex := h.largerChildIndex(index)
		h.swap(index, largerChildIndex)
		index = largerChildIndex
	}
}

func (h *Heap) isFull() bool {
	return h.size == len(h.items)
}

func (h *Heap) isEmpty() bool {
	return h.size == 0
}

func (h *Heap) leftChildIndex(index int) int {
	return index*2 + 1
}

func (h *Heap) leftChild(index int) int {
	return h.items[h.leftChildIndex(index)]
}

func (h *Heap) rightChildIndex(index int) int {
	return index*2 + 2
}

func (h *Heap) rightChild(index int) int {
	return h.items[h.rightChildIndex(index)]
}

func (h *Heap) isValidParent(index int) bool {
	if !h.hasLeftChild(index) {
		return true
	}

	isValid := h.items[index] >= h.leftChild(index)

	if h.hasRightChild(index) {
		isValid = isValid && h.items[index] >= h.rightChild(index)
	}

	return isValid
}

func (h *Heap) largerChildIndex(index int) int {
	if !h.hasLeftChild(index) {
		return index
	}

	if !h.hasRightChild(index) {
		return h.leftChildIndex(index)
	}

	if h.leftChild(index) > h.rightChild(index) {
		return h.leftChildIndex(index)
	} else {
		return h.rightChildIndex(index)
	}
}

func (h *Heap) hasLeftChild(index int) bool {
	return h.leftChildIndex(index) <= h.size
}

func (h *Heap) hasRightChild(index int) bool {
	return h.rightChildIndex(index) <= h.size
}

func (h *Heap) String() string {
	return fmt.Sprintf("%#v\n", h)
}
