package tries

type char byte

type node struct {
	val         char
	children    map[char]*node
	isEndOfWord bool
}

func (n *node) hasChild(c char) bool {
	_, ok := n.children[c]
	return ok
}

func (n *node) addChild(c char) {
	n.children[c] = &node{val: c, children: make(map[char]*node)}
}

func (n *node) getChild(c char) *node {
	return n.children[c]
}

func (n *node) getChildren() []*node {
	children := make([]*node, 0, len(n.children))
	
	for _, v := range n.children {
		children = append(children, v)
	}

	return children
}

func (n *node) hasChildren() bool {
	return !(len(n.children) == 0)
}

func (n *node) removeChild(c char) {
	delete(n.children, c)
}
