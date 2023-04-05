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
		},
	}
}

func (t *Trie) Insert(word string) {
	cur := t.root
	for i := 0; i < len(word); i++ {
		c := word[i]
		if !cur.hasChild(char(c)){
			cur.addChild(char(c))
		} 
		
		cur = cur.getChild(char(c))
	}

	cur.isEndOfWord = true
}

func (t *Trie) Contains(word string) bool {
	cur := t.root
	for i := 0; i < len(word); i++ {
		c := word[i]
		if !cur.hasChild(char(c)) {
			return false
		}

		cur = cur.getChild(char(c))
	}

	return cur.isEndOfWord 
}

func (t *Trie) Traverse() {
	traverse(t.root)
}

func (t *Trie) Remove(word string) {
	remove(t.root, word, 0)
}

func remove(n *node, word string, index int) {
	if index == len(word) {
		n.isEndOfWord = false
		return
	}

	c := word[index]
	child := n.getChild(char(c))
	if child == nil {
		return
	}

	remove(child, word, index+1)

	if !child.hasChildren() && !child.isEndOfWord {
		n.removeChild(char(c))
	}
}

func (t *Trie) String() string {
	var result string

	traverse(t.root)

	return result
}

func traverse(n *node) {
	// Pre-order traversal
	fmt.Println(string(n.val))

	for _, v := range n.getChildren() {
		traverse(v)
	}

	// Post-order traversal
	// fmt.Println(string(n.val))
}
