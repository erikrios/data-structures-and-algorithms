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

	a.setHeight(root)

	return a.balance(root)
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

func (a *AVLTree) balance(root *aVLNode) *aVLNode {
	if a.isLeftHeavy(root) {
		if a.balanceFactor(root.leftChild) < 0 {
			root.leftChild = a.rotateLeft(root.leftChild)
		}
		return a.rotateRight(root)
	} else if a.isRightHeavy(root) {
		if a.balanceFactor(root.rightChild) > 0 {
			root.rightChild = a.rotateRight(root.rightChild)
		}
		return a.rotateLeft(root)
	}
	return root
}

func (a *AVLTree) rotateLeft(root *aVLNode) *aVLNode {
	newRoot := root.rightChild

	root.rightChild = newRoot.leftChild
	newRoot.leftChild = root

	a.setHeight(root)
	a.setHeight(newRoot)

	return newRoot
}

func (a *AVLTree) rotateRight(root *aVLNode) *aVLNode {
	newRoot := root.leftChild

	root.leftChild = newRoot.rightChild
	newRoot.rightChild = root

	a.setHeight(root)
	a.setHeight(newRoot)

	return newRoot
}

func (a *AVLTree) setHeight(node *aVLNode) {
	var height int

	right := a.height(node.rightChild)
	left := a.height(node.leftChild)

	if left < right {
		height = right
	} else {
		height = left
	}

	node.height = height + 1
}
