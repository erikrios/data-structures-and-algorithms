package tries

type char byte

type node struct {
	val         char
	children    [26]*node
	isEndOfWord bool
}
