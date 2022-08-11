package avltree

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

	return root
}

func (a *AVLTree) height(node *aVLNode) int {
	if node == nil {
		return 0
	} else {
		return node.height
	}
}
