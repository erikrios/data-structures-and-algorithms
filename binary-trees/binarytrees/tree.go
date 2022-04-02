package binarytrees

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
