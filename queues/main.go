package main

import (
	"fmt"

	"github.com/erikrios/queues/queues"
)

func main() {
	var queue queues.Queue[int] = queues.NewArrayQueue[int](5)

	fmt.Println(queue.IsEmpty())
	fmt.Println(queue.IsFull())

	if err := queue.Enqueue(5); err != nil {
		fmt.Println(err.Error())
	}
	if err := queue.Enqueue(4); err != nil {
		fmt.Println(err.Error())
	}
	if err := queue.Enqueue(3); err != nil {
		fmt.Println(err.Error())
	}
	if err := queue.Enqueue(2); err != nil {
		fmt.Println(err.Error())
	}
	if err := queue.Enqueue(1); err != nil {
		fmt.Println(err.Error())
	}
	if err := queue.Enqueue(0); err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(queue.IsEmpty())
	fmt.Println(queue.IsFull())

	fmt.Println(queue)
}
