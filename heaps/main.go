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
	fmt.Println("Done", heap)
}
