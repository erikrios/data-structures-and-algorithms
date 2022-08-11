package avltree

import "testing"

func TestInsert(t *testing.T) {
	avlTree := NewAVLTree()
	avlTree.Insert(5)
	avlTree.Insert(3)
	avlTree.Insert(2)
	avlTree.Insert(8)
}
