package main

import (
	"fmt"
	"heaps/heaps"
)

func main() {
	heap := heaps.NewHeap()
	heap.Insert(10)
	heap.Insert(5)
	heap.Insert(17)
	heap.Insert(4)
	heap.Insert(22)
	heap.Remove()
	fmt.Println("Done", heap)

	numbers := []int{5, 3, 8, 4, 1, 2}
	fmt.Println(heaps.GetKthLargest(numbers, 2))
	heaps.Heapify(numbers)
	fmt.Println(numbers)
}
