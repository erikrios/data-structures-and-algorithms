package binarytrees

import (
	"fmt"
	"math"
)

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

	return 1 + int(math.Max(float64(t.height(root.leftChild)), float64(t.height(root.rightChild))))
}

func (t *Tree[T]) isLeaf(node *node[T]) bool {
	return node.leftChild == nil && node.rightChild == nil
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

	return t.isBinarySearchTree(root.leftChild, min, root.val-1) &&
		t.isBinarySearchTree(root.rightChild, root.val+1, max)
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
