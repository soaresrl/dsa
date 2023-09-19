package main

import (
	"github.com/soaresrl/dsa/pkg/extendible_hash"
	"fmt"
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

	fmt.Println("Recuperando item (14)")

	item := hash.Get(14)
	
	fmt.Printf("item = %d\n", item) 
	
	fmt.Println()

	fmt.Println("Escrevendo a hash completa")
	hash.Print()

	fmt.Println()

	fmt.Println("Removendo itens 40 e 14.")

	hash.Remove(40)
	hash.Remove(14)

	fmt.Println()

	fmt.Println("Recuperando item (40)")

	item = hash.Get(40)

	if item == -1 {
		fmt.Println("Não encontrado")
	} else {
		fmt.Printf("item = %d\n", item)
	}

	fmt.Println()

	fmt.Println("Escrevendo a hash completa")
	hash.Print()

	extendible_hash.Free(&hash)
}
