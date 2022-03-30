package main

import (
	"errors"
	"fmt"

	"github.com/erikrios/linked-list/linkedlist"
)

func main() {
	ll := linkedlist.NewLinkedList[int]()
	fmt.Println(ll)
	fmt.Println(ll.Size())
	fmt.Println(ll.ToSlice())
	ll.AddFirst(5)
	fmt.Println(ll)
	ll.AddFirst(4)
	fmt.Println(ll)
	ll.AddFirst(3)
	fmt.Println(ll)
	ll.AddFirst(2)
	fmt.Println(ll)
	ll.AddFirst(1)
	fmt.Println(ll)
	fmt.Println(ll.Contains(0))
	fmt.Println(ll.IndexOf(4))
	fmt.Println(ll.Size())
	fmt.Println(ll.ToSlice())

	ll2 := linkedlist.NewLinkedList[string]()
	fmt.Println(ll2)
	ll2.AddLast("five")
	fmt.Println(ll2)
	ll2.AddLast("four")
	fmt.Println(ll2)
	ll2.AddLast("three")
	fmt.Println(ll2)
	ll2.AddLast("two")
	fmt.Println(ll2)
	ll2.AddLast("one")
	fmt.Println(ll2)
	fmt.Println(ll2.Contains("three"))
	fmt.Println(ll2.IndexOf("ten"))
	fmt.Println(ll2.Size())
	fmt.Println(ll2.ToSlice())

	ll3 := linkedlist.NewLinkedList[int8]()
	fmt.Println(ll3)
	if err := ll3.DeleteFirst(); err != nil {
		if errors.Is(err, linkedlist.ErrNoSuchElement) {
			fmt.Println(err.Error())
		}
	}
	ll3.AddFirst(5)
	ll3.AddFirst(4)
	ll3.AddFirst(3)
	ll3.AddFirst(2)
	ll3.AddFirst(1)
	fmt.Println(ll3)
	if err := ll3.DeleteFirst(); err != nil {
		if errors.Is(err, linkedlist.ErrNoSuchElement) {
			fmt.Println(err.Error())
		}
	}
	if err := ll3.DeleteFirst(); err != nil {
		if errors.Is(err, linkedlist.ErrNoSuchElement) {
			fmt.Println(err.Error())
		}
	}
	if err := ll3.DeleteFirst(); err != nil {
		if errors.Is(err, linkedlist.ErrNoSuchElement) {
			fmt.Println(err.Error())
		}
	}
	fmt.Println(ll3)
	fmt.Println(ll3.Contains(5))
	fmt.Println(ll3.IndexOf(4))
	fmt.Println(ll3.Size())
	fmt.Println(ll3.ToSlice())

	ll4 := linkedlist.NewLinkedList[float32]()
	fmt.Println(ll4)
	if err := ll4.DeleteLast(); err != nil {
		if errors.Is(err, linkedlist.ErrNoSuchElement) {
			fmt.Println(err.Error())
		}
	}
	ll4.AddLast(5.1)
	ll4.AddLast(4.2)
	ll4.AddLast(3.3)
	ll4.AddLast(2.4)
	ll4.AddLast(1.5)
	if err := ll4.DeleteLast(); err != nil {
		if errors.Is(err, linkedlist.ErrNoSuchElement) {
			fmt.Println(err.Error())
		}
	}
	fmt.Println(ll4)
	fmt.Println(ll4.Contains(1.5))
	fmt.Println(ll4.IndexOf(2.4))
	fmt.Println(ll4.Size())
	fmt.Println(ll4.ToSlice())

	ll5 := linkedlist.NewLinkedList[float64]()
	fmt.Println(ll5)
	ll5.AddFirst(9.3)
	ll5.AddFirst(10.5)
	ll5.AddFirst(11.4)
	ll5.AddFirst(12.8)
	ll5.AddFirst(13.9)
	fmt.Println(ll5)
	ll5.Reverse()
	fmt.Println(ll5)
	fmt.Println(ll5.GetKthFromTheEnd(1))
	fmt.Println(ll5.GetKthFromTheEnd(2))
	fmt.Println(ll5.GetKthFromTheEnd(3))
	fmt.Println(ll5.GetKthFromTheEnd(4))
	fmt.Println(ll5.GetKthFromTheEnd(5))
	fmt.Println(ll5.GetKthFromTheEnd(6))

	ll5.PrintMiddle()
	fmt.Println(ll5.HasLoop())

	ll6 := linkedlist.CreateWithLoop()
	fmt.Println(ll6.HasLoop())

}
