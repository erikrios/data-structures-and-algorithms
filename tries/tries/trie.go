package tries

import "fmt"

type Trie struct {
	root *node
}

func New() *Trie {
	return &Trie{
		root: &node{
			val:         ' ',
			isEndOfWord: true,
		},
	}
}

func (t *Trie) Insert(word string) {
	cur := t.root
	for i := 0; i < len(word); i++ {
		c := word[i]
		index := c - 'a'
		if n := cur.children[index]; n == nil {
			n = &node{
				val:         char(c),
				isEndOfWord: true,
			}
			cur.isEndOfWord = false
			cur.children[index] = n
			cur = n
		} else {
			n.isEndOfWord = true
			cur.isEndOfWord = false
			cur.children[index] = n
			cur = n
		}
	}
}

func (t *Trie) String() string {
	var result string

	traverse(t.root)

	return result
}

func traverse(n *node) {
	if n == nil {
		return
	}

	fmt.Println("-", string(n.val), *n)
	for i := 0; i < len(n.children); i++ {
		child := n.children[i]
		traverse(child)
	}
}
