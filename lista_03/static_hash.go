package main

import (
	"fmt"
	"github.com/soaresrl/dsa/pkg/static_hash"
)

func main() {
	hash := static_hash.NewHashTable(10)

	for i := 0; i < 10; i++ {
		hash.Insert(i)
	}

	hash.Print()
	
	fmt.Println()

	fmt.Println("Searching element 7: ")

	node := hash.Get(7)

	if node == nil {
		fmt.Println("Key not found")
	} else {
		fmt.Printf("value = %v\n", node.Data)
	} 

	hash.Remove(3)

	fmt.Println()

	fmt.Println("After removing from key 3")

	fmt.Println()

	hash.Print()

	static_hash.Free(&hash)
}
