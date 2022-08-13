package heaps

func Heapify(items []int) {
	lastParentIndex := len(items)/2 - 1
	for i := lastParentIndex; i >= 0; i-- {
		heapify(items, i)
	}
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
