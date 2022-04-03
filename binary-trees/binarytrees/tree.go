package binarytrees

import (
	"fmt"
)

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

func (t *Tree[T]) Height() int {
	return t.height(t.root)
}

func (t *Tree[T]) height(root *node[T]) int {
	if root == nil {
		return -1
	}

	if t.isLeaf(root) {
		return 0
	}

	maxLeft := t.height(root.leftChild)
	maxRight := t.height(root.rightChild)
	var max int
	if maxLeft > maxRight {
		max = maxLeft
	} else {
		max = maxRight
	}

	return 1 + max
}

func (t *Tree[T]) Min() T {
	return t.min(t.root)
}

func (t *Tree[T]) min(root *node[T]) T {
	if root == nil {
		return T(100)
	}
	if t.isLeaf(root) {
		return root.val
	}

	left := t.min(root.leftChild)
	right := t.min(root.rightChild)

	var minRightOrLeft T
	if left < right {
		minRightOrLeft = left
	} else {
		minRightOrLeft = right
	}

	var minimum T
	if minRightOrLeft < root.val {
		minimum = minRightOrLeft
	} else {
		minimum = root.val
	}

	return minimum
}

func (t *Tree[T]) isLeaf(root *node[T]) bool {
	return root.leftChild == nil && root.rightChild == nil
}
