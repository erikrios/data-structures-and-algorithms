package binarytrees

import "testing"

func TestNewTree(t *testing.T) {
	t.Run("it should return not nill instance, when creating instance with NewTree method", func(t *testing.T) {
		tree := NewTree[int]()
		if tree == nil {
			t.Fatalf("got %v, want %v", tree, "not nil")
		}
	})

	t.Run("it should have nil root, when first initialization", func(t *testing.T) {
		tree := NewTree[float32]()
		if tree.root != nil {
			t.Fatalf("got %v, want %v", tree.root, nil)
		}
	})
}
