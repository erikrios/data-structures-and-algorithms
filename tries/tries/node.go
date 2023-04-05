package tries

type char byte

type node struct {
	val         char
	children    map[char]*node
	isEndOfWord bool
}
