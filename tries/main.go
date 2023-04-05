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

	trie.Insert("canada")
	fmt.Println(trie.Contains(""))
	fmt.Println(trie.Contains("can"))
	fmt.Println(trie.Contains("canada"))

	trie.Insert("awesome")

	trie.Traverse()
}
