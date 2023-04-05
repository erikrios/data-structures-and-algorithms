package tries

import "fmt"

type Trie struct {
	root *node
}

func New() *Trie {
	return &Trie{
		root: &node{
			val:         ' ',
			children: make(map[char]*node),
			isEndOfWord: true,
		},
	}
}

func (t *Trie) Insert(word string) {
	cur := t.root
	for i := 0; i < len(word); i++ {
		c := word[i]
		if n, ok := cur.children[char(c)]; !ok{
			n = &node{
				val:         char(c),
			children: make(map[char]*node),
				isEndOfWord: true,
			}
			cur.isEndOfWord = false
			cur.children[char(c)] = n
			cur = n
		} else {
			n.isEndOfWord = true
			cur.isEndOfWord = false
			cur.children[char(c)] = n
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
	for _, v := range n.children{
		traverse(v)
	}
}
