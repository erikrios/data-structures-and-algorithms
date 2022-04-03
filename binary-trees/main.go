package main

import (
	"fmt"

	"github.com/erikrios/binary-trees/binarytrees"
)

func main() {
	tree := binarytrees.NewTree[int]()
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(6)
	tree.Insert(1)
	tree.Insert(8)
	tree.Insert(12)
	tree.Insert(18)
	tree.Insert(17)
	fmt.Println(tree.Find(18))

	newTree := binarytrees.NewTree[int]()
	newTree.Insert(20)
	newTree.Insert(10)
	newTree.Insert(30)
	newTree.Insert(6)
	newTree.Insert(14)
	newTree.Insert(3)
	newTree.Insert(8)
	newTree.Insert(24)
	newTree.Insert(26)
	fmt.Println(newTree)
	newTree.TraversePreOrder()
	fmt.Println()
	newTree.TraverseInOrder()
	fmt.Println()
	newTree.TraversePostOrder()
	fmt.Println()
	fmt.Println(newTree.Height())
	fmt.Println(newTree.Min())
}