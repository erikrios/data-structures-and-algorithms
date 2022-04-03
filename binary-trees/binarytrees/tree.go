package binarytrees

import "fmt"

type Tree[T Number] struct {
	root *node[T]
}

func NewTree[T Number]() *Tree[T] {
	return &Tree[T]{root: nil}
}

func (t *Tree[T]) Insert(value T) {
	newNode := &node[T]{val: value}

	if t.root == nil {
		t.root = newNode
		return
	}

	current := t.root
	for {
		if value < current.val {
			if current.leftChild == nil {
				current.leftChild = newNode
				break
			}
			current = current.leftChild
		} else {
			if current.rightChild == nil {
				current.rightChild = newNode
				break
			}
			current = current.rightChild
		}
	}
}

func (t *Tree[T]) Find(value T) bool {
	current := t.root

	for current != nil {
		if value == current.val {
			return true
		} else if value < current.val {
			current = current.leftChild
		} else {
			current = current.rightChild
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

	fmt.Printf("%v ", root.val)
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
	fmt.Printf("%v ", root.val)
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
	fmt.Printf("%v ", root.val)
}
