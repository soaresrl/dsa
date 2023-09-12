package main

import (
	"github.com/soaresrl/dsa/pkg/static_hash"
)

func main() {
	hash := static_hash.NewHashTable(10)

	for i := 1; i <= 10; i++ {
		hash.Insert(i)
	}

	hash.Print()

	hash.Remove(5)

	hash.Print()

	static_hash.Free(&hash)
}