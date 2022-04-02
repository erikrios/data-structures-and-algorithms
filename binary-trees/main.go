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
}
