package binarytrees

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
