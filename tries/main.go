package main

import (
	"fmt"
	"tries/tries"
)

func main() {
	trie := tries.New()

	trie.Insert("cat")

	trie.Insert("car")

	trie.Insert("canada")
	fmt.Println(trie.Contains(""))
	fmt.Println(trie.Contains("can"))
	fmt.Println(trie.Contains("canada"))

	trie.Insert("awesome")

	trie.Traverse()

	trie.Insert("doctor")
	trie.Insert("doctoral")

	trie.Remove("doctoral")
	trie.Traverse()
	fmt.Println(trie.Contains("doctor"))
	fmt.Println(trie.Contains("doctoral"))
}
