package hashtables

import (
	"fmt"
	"log"

	"github.com/hash-tables/linkedlist"
)

type hashTable[K int, V any] struct {
	list []linkedlist.LinkedList[entry[K, V]]
}

func NewHashTable[K int, V any]() *hashTable[K, V] {
	list := make([]linkedlist.LinkedList[entry[K, V]], 5)
	return &hashTable[K, V]{list: list}
}

func (h *hashTable[K, V]) String() string {
	logger := fmt.Sprintf("%+v\n", h.list)
	log.Println(logger)

	allEntries := make([]entry[K, V], 0)
	for _, ll := range h.list {
		entries := ll.ToSlice()
		allEntries = append(allEntries, entries...)
	}

	return fmt.Sprintf("%+v\n", allEntries)
}

func (h *hashTable[K, V]) Put(key K, value V) {
	hashed := hash(int(key))
	e := entry[K, V]{
		key:   key,
		value: value,
	}
	if _, ok := h.Get(key); ok {
		h.Remove(key)
	}
	h.list[hashed].AddLast(e)
}

func (h *hashTable[K, V]) Get(key K) (value V, ok bool) {
	hashed := hash(int(key))

	entries := h.list[hashed].ToSlice()
	for _, entry := range entries {
		if key == entry.key {
			ok = true
			value = entry.value
			return
		}
	}

	return
}

func (h *hashTable[K, V]) Remove(key K) {
	hashed := hash(int(key))

	entries := h.list[hashed].ToSlice()
	for _, entry := range entries {
		if key == entry.key {
			h.list[hashed].Remove(entry)
		}
	}
}

func hash(num int) int {
	return num % 5
}
