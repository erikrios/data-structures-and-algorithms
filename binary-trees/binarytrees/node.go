package binarytrees

type Number interface {
	uint8 |
		uint16 |
		uint32 |
		uint64 |
		uint |
		int8 |
		int16 |
		int32 |
		int64 |
		int |
		float32 |
		float64
}

type node[T Number] struct {
	leftChild  *node[T]
	rightChild *node[T]
	val        T
}
