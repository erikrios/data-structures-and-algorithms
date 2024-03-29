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

	fmt.Println("Circular Array Queue")
	var caq queues.Queue[int] = queues.NewCircularArrayQueue[int](5)

	if err := caq.Enqueue(50); err != nil {
		fmt.Println("enqueue error:", err.Error())
	}

	if err := caq.Enqueue(40); err != nil {
		fmt.Println("enqueue error:", err.Error())
	}

	if err := caq.Enqueue(30); err != nil {
		fmt.Println("enqueue error:", err.Error())
	}

	if err := caq.Enqueue(20); err != nil {
		fmt.Println("enqueue error:", err.Error())
	}

	if err := caq.Enqueue(10); err != nil {
		fmt.Println("enqueue error:", err.Error())
	}

	if err := caq.Enqueue(50); err != nil {
		fmt.Println("enqueue error:", err.Error())
	}

	if dequeued, err := caq.Dequeue(); err != nil {
		fmt.Println("dequeued error:", err.Error())
	} else {
		fmt.Println("dequeued:", dequeued)
	}

	if peeked, err := caq.Peek(); err != nil {
		fmt.Println("peeked error:", err.Error())
	} else {
		fmt.Println("peeked:", peeked)
	}

	fmt.Println(caq)

	fmt.Println("Queue with Two Stacks")

	var qwts queues.Queue[int] = queues.NewQueueWithTwoStacks[int]()

	if err := qwts.Enqueue(50); err != nil {
		fmt.Println("enqueue error:", err.Error())
	}

	if err := qwts.Enqueue(40); err != nil {
		fmt.Println("enqueue error:", err.Error())
	}

	if err := qwts.Enqueue(30); err != nil {
		fmt.Println("enqueue error:", err.Error())
	}

	if err := qwts.Enqueue(20); err != nil {
		fmt.Println("enqueue error:", err.Error())
	}

	if err := qwts.Enqueue(10); err != nil {
		fmt.Println("enqueue error:", err.Error())
	}

	fmt.Println(qwts)

	if peeked, err := qwts.Peek(); err != nil {
		fmt.Println("peeked error:", err.Error())
	} else {
		fmt.Println("peeked:", peeked)
	}

	if dequeued, err := qwts.Dequeue(); err != nil {
		fmt.Println("dequeued error:", err.Error())
	} else {
		fmt.Println("dequeued:", dequeued)
	}

	if dequeued, err := qwts.Dequeue(); err != nil {
		fmt.Println("dequeued error:", err.Error())
	} else {
		fmt.Println("dequeued:", dequeued)
	}

	fmt.Println(qwts)

	fmt.Println("\nPriority Queue")

	var pq queues.Queue[int] = queues.NewPriorityQueue[int]()

	if err := pq.Enqueue(20); err != nil {
		fmt.Println("enqueue error:", err.Error())
	}

	if err := pq.Enqueue(15); err != nil {
		fmt.Println("enqueue error:", err.Error())
	}

	if err := pq.Enqueue(25); err != nil {
		fmt.Println("enqueue error:", err.Error())
	}

	if peeked, err := pq.Peek(); err != nil {
		fmt.Println("peeked error:", err.Error())
	} else {
		fmt.Println("peeked:", peeked)
	}

	if dequeued, err := pq.Dequeue(); err != nil {
		fmt.Println("dequeued error:", err.Error())
	} else {
		fmt.Println("dequeued:", dequeued)
	}

	if dequeued, err := pq.Dequeue(); err != nil {
		fmt.Println("dequeued error:", err.Error())
	} else {
		fmt.Println("dequeued:", dequeued)
	}

	if dequeued, err := pq.Dequeue(); err != nil {
		fmt.Println("dequeued error:", err.Error())
	} else {
		fmt.Println("dequeued:", dequeued)
	}

	if dequeued, err := pq.Dequeue(); err != nil {
		fmt.Println("dequeued error:", err.Error())
	} else {
		fmt.Println("dequeued:", dequeued)
	}

	fmt.Println(pq)
}
