package binarytrees

import "fmt"

type Tree[T Number] struct {
	root *node[T]
}

func NewTree[T Number]() *Tree[T] {
	return &Tree[T]{root: nil}
}

func (t *Tree[T]) Insert(val T) {
	node := &node[T]{val: val}

	if t.root == nil {
		t.root = node
		return
	}

	currentNode := t.root
	for {
		if val < currentNode.val {
			if currentNode.leftChild == nil {
				currentNode.leftChild = node
				break
			}
			currentNode = currentNode.leftChild
		} else if val > currentNode.val {
			if currentNode.rightChild == nil {
				currentNode.rightChild = node
				break
			}
			currentNode = currentNode.rightChild
		}
	}
}

func (t *Tree[T]) Find(val T) bool {
	currentNode := t.root
	for currentNode != nil {
		if val < currentNode.val {
			currentNode = currentNode.leftChild
		} else if val > currentNode.val {
			currentNode = currentNode.rightChild
		} else {
			return true
		}
	}
	return false
}

func (t *Tree[T]) TraversePreOrder() {
	t.traversePreOrder(t.root)
}

func (t *Tree[T]) traversePreOrder(root *node[T]) {
	if root == nil {
		return
	}

	fmt.Print(root.val, " ")
	t.traversePreOrder(root.leftChild)
	t.traversePreOrder(root.rightChild)
}

func (t *Tree[T]) TraverseInOrder() {
	t.traverseInOrder(t.root)
}

func (t *Tree[T]) traverseInOrder(root *node[T]) {
	if root == nil {
		return
	}

	t.traverseInOrder(root.leftChild)
	fmt.Print(root.val, " ")
	t.traverseInOrder(root.rightChild)
}

func (t *Tree[T]) TraversePostOrder() {
	t.traversePostOrder(t.root)
}

func (t *Tree[T]) traversePostOrder(root *node[T]) {
	if root == nil {
		return
	}

	t.traversePostOrder(root.leftChild)
	t.traversePostOrder(root.rightChild)
	fmt.Print(root.val, " ")
}
