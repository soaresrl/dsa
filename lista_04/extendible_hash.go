package main

import (
	"github.com/soaresrl/dsa/pkg/extendible_hash"
	//"fmt"
)

func main() {
	hash := extendible_hash.NewHashTable(4)


	hash.Insert(40)
	hash.Insert(44)
	
	hash.Insert(49)
	hash.Insert(21)
	hash.Insert(43)
	hash.Insert(51)

	hash.Insert(10)
	hash.Insert(14)
	hash.Insert(22)
	hash.Insert(42)

	hash.Print()

	hash.Remove(40)
	hash.Remove(44)
	//fmt.Println()

	hash.Print()

	//hash.Remove(14)
	//fmt.Println()

	//hash.Print()
}
