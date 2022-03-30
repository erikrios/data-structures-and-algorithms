package main

import (
	"fmt"

	"github.com/erikrios/stacks/stacks"
)

func main() {
	stack := stacks.NewStack[int]()
	fmt.Println(stack)

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	fmt.Println(stack)
	fmt.Println(stack.IsEmpty())

	if peeked, err := stack.Peek(); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(peeked)
	}

	if popped, err := stack.Pop(); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(popped)
	}

	if popped, err := stack.Pop(); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(popped)
	}

	if popped, err := stack.Pop(); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(popped)
	}

	if popped, err := stack.Pop(); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(popped)
	}

	if popped, err := stack.Pop(); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(popped)
	}

	fmt.Println(stack)
	if peeked, err := stack.Peek(); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(peeked)
	}
	fmt.Println(stack.IsEmpty())

	fmt.Println("Two Stacks")
	twoStack := stacks.NewTwoStack[int](5)
	fmt.Println(twoStack.IsEmpty1())
	fmt.Println(twoStack.IsEmpty2())
	fmt.Println(twoStack.IsFull1())
	fmt.Println(twoStack.IsFull2())

	twoStack.Push1(1)
	twoStack.Push2(5)
	fmt.Println(twoStack)
	fmt.Println(twoStack.IsEmpty1())
	fmt.Println(twoStack.IsEmpty2())
	fmt.Println(twoStack.IsFull1())
	fmt.Println(twoStack.IsFull2())

	fmt.Println(twoStack.Pop1())
	fmt.Println(twoStack.Pop2())
	fmt.Println(twoStack.IsEmpty1())
	fmt.Println(twoStack.IsEmpty2())
	fmt.Println(twoStack.IsFull1())
	fmt.Println(twoStack.IsFull2())

	fmt.Println("Min Stacks")
	minStack := stacks.NewMinimumStack[int]()
	minStack.Push(4)
	minStack.Push(3)
	minStack.Push(5)
	fmt.Println(minStack.Min())
}
