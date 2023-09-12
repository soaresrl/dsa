package static_hash

import (
	"fmt"
	"github.com/soaresrl/dsa/pkg/linked_list"
)

type HashTable struct {
	Size int
	Items []*linked_list.Node
}

func hashfunction(key, m int) int {
   return key % m
}

func NewHashTable(size int) *HashTable {
	table := HashTable{
		Size: size / 2,
		Items: make([]*linked_list.Node, size / 2),
	}

	return &table
}

func (h *HashTable) Insert(value int) {
	pos := hashfunction(value, h.Size)

	h.Items[pos] = h.Items[pos].Insert(value)
}

func (h *HashTable) Get(key int) *linked_list.Node {
	pos := hashfunction(key, h.Size)

	node := h.Items[pos].Find(key)

	return node
}

func (h *HashTable) Remove(key int) {
	pos := hashfunction(key, h.Size)

	h.Items[pos] = h.Items[pos].Remove(key)
}

func Free(h **HashTable) {
	for _, v := range (*h).Items {
		linked_list.Free(&v)
	}

	*h = nil
}

func (h *HashTable) Print() {
	for i, v := range h.Items {
		fmt.Printf("Slot %d \n", i)
		v.Print()
	}
}