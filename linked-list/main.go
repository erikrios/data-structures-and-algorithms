package main

import (
	"fmt"

	"github.com/erikrios/linked-list/linkedlist"
)

func main() {
	ll := linkedlist.NewLinkedList[int]()
	fmt.Println(ll)
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

	ll3 := linkedlist.NewLinkedList[int8]()
	fmt.Println(ll3)
	ll3.AddFirst(5)
	ll3.AddFirst(4)
	ll3.AddFirst(3)
	ll3.AddFirst(2)
	ll3.AddFirst(1)
	fmt.Println(ll3)
	ll3.DeleteFirst()
	ll3.DeleteFirst()
	ll3.DeleteFirst()
	fmt.Println(ll3)

	ll4 := linkedlist.NewLinkedList[float32]()
	fmt.Println(ll4)
	ll4.AddLast(5.1)
	ll4.AddLast(4.2)
	ll4.AddLast(3.3)
	ll4.AddLast(2.4)
	ll4.AddLast(1.5)
	ll4.DeleteLast()
	fmt.Println(ll4)
}
