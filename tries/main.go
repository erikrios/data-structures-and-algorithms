package main

import (
	"fmt"
	"tries/tries"
)

func main() {
	trie := tries.New()

	trie.Insert("cat")
	fmt.Println(trie)

	trie.Insert("car")
	fmt.Println(trie)
}
