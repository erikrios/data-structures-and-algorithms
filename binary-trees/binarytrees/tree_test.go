package binarytrees

import (
	"testing"
)

func TestNewTree(t *testing.T) {
	t.Run("it should return not nill instance, when creating instance with NewTree method", func(t *testing.T) {
		tree := NewTree[int]()

		if tree == nil {
			t.Fatalf("got %v, want %v", tree, "not nil")
		}
	})

	t.Run("it should have nil root, when first initialization", func(t *testing.T) {
		tree := NewTree[float32]()

		got := tree.root

		if got != nil {
			t.Fatalf("got %v, want %v", got, nil)
		}
	})
}

func TestInsert(t *testing.T) {
	t.Run("it should have expected root value, when inserting a first value", func(t *testing.T) {
		val := 5.1
		tree := NewTree[float64]()
		tree.Insert(val)

		got := tree.root.val

		if got != val {
			t.Fatalf("got %f, want %f", got, val)
		}
	})

	t.Run("it should have expected root left child, when inserting value that less than root", func(t *testing.T) {
		vals := []int{4, 3}
		expected := vals[len(vals)-1]
		tree := NewTree[int]()
		for _, val := range vals {
			tree.Insert(val)
		}

		got := tree.root.leftChild.val

		if got != expected {
			t.Fatalf("got %d, want %d", got, expected)
		}
	})

	t.Run("it should have expected root right child, when inserting value that greater than root", func(t *testing.T) {
		vals := []int{8, 12}
		expected := vals[len(vals)-1]
		tree := NewTree[int]()
		for _, val := range vals {
			tree.Insert(val)
		}

		got := tree.root.rightChild.val

		if got != expected {
			t.Fatalf("got %d, want %d", got, expected)
		}
	})

	t.Run("it should have expected root right and left child, when inserting three values", func(t *testing.T) {
		vals := []float32{4.1, 3.2, 5.8}
		expectedLeftChild := vals[len(vals)-2]
		expectedRightChild := vals[len(vals)-1]
		tree := NewTree[float32]()
		for _, val := range vals {
			tree.Insert(val)
		}

		gotLeftChild := tree.root.leftChild.val
		gotRightChild := tree.root.rightChild.val

		if gotLeftChild != expectedLeftChild || gotRightChild != expectedRightChild {
			t.Fatalf("got %f and %f, want %f and %f", gotLeftChild, gotRightChild, expectedLeftChild, expectedRightChild)
		}
	})

	t.Run("it should have expected left leaf, when inserting values that less than root", func(t *testing.T) {
		vals := []int{8, 5, 3}
		expected := vals[len(vals)-1]
		tree := NewTree[int]()
		for _, val := range vals {
			tree.Insert(val)
		}

		got := tree.root.leftChild.leftChild.val

		if got != expected {
			t.Fatalf("got %d, want %d", got, expected)
		}
	})

	t.Run("it should have expected right leaf, when inserting values that greater than root", func(t *testing.T) {
		vals := []int{3, 5, 8}
		expected := vals[len(vals)-1]
		tree := NewTree[int]()
		for _, val := range vals {
			tree.Insert(val)
		}

		got := tree.root.rightChild.rightChild.val

		if got != expected {
			t.Fatalf("got %d, want %d", got, expected)
		}
	})

	t.Run("it should have expected left and right leaf, when inserting values", func(t *testing.T) {
		vals := []int{10, 5, 15, 6, 1, 8, 12, 18, 17}
		expectedLeftLeaf := 8
		expectedRightLeaf := 17
		tree := NewTree[int]()
		for _, val := range vals {
			tree.Insert(val)
		}

		gotLeftChild := tree.root.leftChild.rightChild.rightChild.val
		gotRightChild := tree.root.rightChild.rightChild.leftChild.val

		if gotLeftChild != expectedLeftLeaf || gotRightChild != expectedRightLeaf {
			t.Fatalf("got %d and %d, want %d and %d", gotLeftChild, gotRightChild, expectedLeftLeaf, expectedRightLeaf)
		}
	})
}

func TestFind(t *testing.T) {
	t.Run("it should return false, when value not exists in tree", func(t *testing.T) {
		expected := false
		tree := NewTree[int]()
		got := tree.Find(5)

		if got != expected {
			t.Fatalf("got %v, want %v", got, expected)
		}
	})

	t.Run("it should return true, when value exists in tree", func(t *testing.T) {
		expected := true
		tree := NewTree[int]()
		tree.Insert(8)
		tree.Insert(3)
		tree.Insert(5)
		tree.Insert(10)
		got := tree.Find(5)

		if got != expected {
			t.Fatalf("got %v, want %v", got, expected)
		}
	})
}
