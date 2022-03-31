package main

import (
	"fmt"

	"github.com/erikrios/queues/queues"
)

func main() {
	var queue queues.Queue[int] = queues.NewArrayQueue[int](5)

	fmt.Println(queue.IsEmpty())
	fmt.Println(queue.IsFull())

	if peeked, err := queue.Peek(); err != nil {
		fmt.Println("peeked error:", err.Error())
	} else {
		fmt.Println("peeked:", peeked)
	}

	if dequeued, err := queue.Dequeue(); err != nil {
		fmt.Println("dequeued error:", err.Error())
	} else {
		fmt.Println("dequeued:", dequeued)
	}

	fmt.Println(queue)

	if err := queue.Enqueue(5); err != nil {
		fmt.Println("enqueue error:", err.Error())
	}

	if peeked, err := queue.Peek(); err != nil {
		fmt.Println("peeked error:", err.Error())
	} else {
		fmt.Println("peeked:", peeked)
	}

	if err := queue.Enqueue(4); err != nil {
		fmt.Println("enqueue error:", err.Error())
	}

	if peeked, err := queue.Peek(); err != nil {
		fmt.Println("peeked error:", err.Error())
	} else {
		fmt.Println("peeked:", peeked)
	}

	if err := queue.Enqueue(3); err != nil {
		fmt.Println("enqueue error:", err.Error())
	}

	if peeked, err := queue.Peek(); err != nil {
		fmt.Println("peeked error:", err.Error())
	} else {
		fmt.Println("peeked:", peeked)
	}

	if err := queue.Enqueue(2); err != nil {
		fmt.Println("enqueue error:", err.Error())
	}

	if peeked, err := queue.Peek(); err != nil {
		fmt.Println("peeked error:", err.Error())
	} else {
		fmt.Println("peeked:", peeked)
	}

	if err := queue.Enqueue(1); err != nil {
		fmt.Println("enqueue error:", err.Error())
	}

	if peeked, err := queue.Peek(); err != nil {
		fmt.Println("peeked error:", err.Error())
	} else {
		fmt.Println("peeked:", peeked)
	}

	if err := queue.Enqueue(0); err != nil {
		fmt.Println("enqueue error:", err.Error())
	}

	if dequeued, err := queue.Dequeue(); err != nil {
		fmt.Println("dequeued error:", err.Error())
	} else {
		fmt.Println("dequeued:", dequeued)
	}

	if peeked, err := queue.Peek(); err != nil {
		fmt.Println("peeked error:", err.Error())
	} else {
		fmt.Println("peeked:", peeked)
	}

	fmt.Println(queue)

	fmt.Println(queue.IsEmpty())
	fmt.Println(queue.IsFull())

	fmt.Println(queue)
}
