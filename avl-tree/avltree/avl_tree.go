package avltree

type aVLNode struct {
	leftChild  *aVLNode
	rightChild *aVLNode
	value      int
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

	return root
}
