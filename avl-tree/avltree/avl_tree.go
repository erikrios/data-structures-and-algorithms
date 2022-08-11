package avltree

import (
	"fmt"
)

type aVLNode struct {
	leftChild  *aVLNode
	rightChild *aVLNode
	value      int
	height     int
}

type AVLTree struct {
	root *aVLNode
}

func NewAVLTree() *AVLTree {
	return &AVLTree{root: nil}
}

func (a *AVLTree) Insert(item int) {
	a.root = a.insert(item, a.root)
}

func (a *AVLTree) insert(item int, root *aVLNode) *aVLNode {
	if root == nil {
		node := &aVLNode{value: item}
		root = node
		return root
	}

	if item < root.value {
		root.leftChild = a.insert(item, root.leftChild)
	} else {
		root.rightChild = a.insert(item, root.rightChild)
	}

	var height int

	right := a.height(root.rightChild)
	left := a.height(root.leftChild)

	if left < right {
		height = right
	} else {
		height = left
	}

	root.height = height + 1

	_ = a.balanceFactor(root)

	if a.isLeftHeavy(root) {
		fmt.Println(root.value, "is left heavy")
	} else if a.isRightHeavy(root) {
		fmt.Println(root.value, "is right heavy")
	}

	return root
}

func (a *AVLTree) isLeftHeavy(node *aVLNode) bool {
	return a.balanceFactor(node) > 1
}

func (a *AVLTree) isRightHeavy(node *aVLNode) bool {
	return a.balanceFactor(node) < -1
}

func (a *AVLTree) balanceFactor(node *aVLNode) int {
	if node == nil {
		return 0
	} else {
		return a.height(node.leftChild) - a.height(node.rightChild)
	}
}

func (a *AVLTree) height(node *aVLNode) int {
	if node == nil {
		return -1
	} else {
		return node.height
	}
}
