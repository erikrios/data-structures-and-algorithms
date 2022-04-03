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

func (t *Tree[T]) Equals(other *Tree[T]) bool {
	if other == nil {
		return false
	}

	return t.equals(t.root, other.root)
}

func (t *Tree[T]) equals(first, second *node[T]) bool {
	if first == nil && second == nil {
		return true
	}

	if first != nil && second != nil {
		return first.val == second.val &&
			t.equals(first.leftChild, second.leftChild) &&
			t.equals(first.rightChild, second.rightChild)
	}

	return false
}

func (t *Tree[T]) IsBinarySearchTree() bool {
	return t.isBinarySearchTree(t.root, T(0), T(127))
}

func (t *Tree[T]) isBinarySearchTree(root *node[T], min, max T) bool {
	if root == nil {
		return true
	}

	if root.val < min || root.val > max {
		return false
	}

	return t.isBinarySearchTree(root.leftChild, min, root.val-1) && t.isBinarySearchTree(root.rightChild, root.val+1, max)
}

func (t *Tree[T]) GetNodesAtDistance(distance int) (nodeValues []T) {
	t.printNodesAtDistance(t.root, distance, &nodeValues)
	return
}

func (t *Tree[T]) printNodesAtDistance(root *node[T], distance int, nodeValues *[]T) {
	if root == nil {
		return
	}

	if distance == 0 {
		*nodeValues = append(*nodeValues, root.val)
		return
	}

	t.printNodesAtDistance(root.leftChild, distance-1, nodeValues)
	t.printNodesAtDistance(root.rightChild, distance-1, nodeValues)
}

func (t *Tree[T]) TraverseLevelOrder() {
	for i := 0; i <= t.Height(); i++ {
		for _, v := range t.GetNodesAtDistance(i) {
			fmt.Printf("%v ", v)
		}
	}
}
