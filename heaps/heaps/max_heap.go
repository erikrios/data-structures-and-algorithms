package heaps

import "errors"

func Heapify(items []int) {
	lastParentIndex := len(items)/2 - 1
	for i := lastParentIndex; i >= 0; i-- {
		heapify(items, i)
	}
}

func GetKthLargest(items []int, k int) (int, error) {
	if k < 1 || k > len(items) {
		return 0, errors.New("k must less than 1 or greater than length of items")
	}

	heap := NewHeap()
	for _, item := range items {
		heap.Insert(item)
	}

	for i := 0; i < k-1; i++ {
		heap.Remove()
	}

	return heap.Max()
}

func heapify(items []int, index int) {
	largerIndex := index

	leftIndex := index*2 + 1
	if leftIndex < len(items) && items[leftIndex] > items[largerIndex] {
		largerIndex = leftIndex
	}

	rightIndex := index*2 + 2
	if rightIndex < len(items) && items[rightIndex] > items[largerIndex] {
		largerIndex = rightIndex
	}

	if index == largerIndex {
		return
	}

	swap(items, index, largerIndex)
	heapify(items, largerIndex)
}

func swap(items []int, first, second int) {
	items[first], items[second] = items[second], items[first]
}
