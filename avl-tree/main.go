package main

import (
	"avl-tree/avltree"
)

func main() {
	avlTree := avltree.NewAVLTree()
	avlTree.Insert(10)
	avlTree.Insert(30)
	avlTree.Insert(20)
}
