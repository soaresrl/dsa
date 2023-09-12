package main

import (
	"fmt"
	"github.com/soaresrl/dsa/pkg/static_hash"
)

func main() {
	hash := static_hash.NewHashTable(10)

	for i := 1; i <= 10; i++ {
		hash.Insert(i)
	}

	hash.Print()

	hash.Remove(5)
	
	fmt.Println()

	fmt.Println("After removing from key 5")

	fmt.Println()

	hash.Print()

	static_hash.Free(&hash)
}
