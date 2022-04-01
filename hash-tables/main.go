package main

import (
	"fmt"

	"github.com/hash-tables/exercise"
	"github.com/hash-tables/hashtables"
)

func main() {
	fmt.Println(string(exercise.FirstNonRepeatingChar("a green apple")))
	fmt.Println(string(exercise.FirstRepeatedChar("a green apple")))

	fmt.Println("\nHash Table")
	hashTable := hashtables.NewHashTable[int, string]()
	hashTable.Put(1, "Erik")
	hashTable.Put(3, "Rio")
	hashTable.Put(11, "Setiawan")
	hashTable.Put(10, "Reza")
	hashTable.Put(3, "Rio Setiawan")

	deleteKey := 11
	hashTable.Remove(deleteKey)

	key := 11
	if val, ok := hashTable.Get(key); !ok {
		fmt.Printf("Key %d not exists\n", key)
	} else {
		fmt.Printf("Key %d have value %s\n", key, val)
	}

	fmt.Println(hashTable)
}
