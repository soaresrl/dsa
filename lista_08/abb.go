package main

import (
	"fmt"

	"github.com/soaresrl/dsa/pkg/abb"
)

func main() {
	tree := abb.NewAbb()

	tree.InsertIterative(7)
	tree.InsertIterative(4)
	tree.InsertIterative(9)
	tree.InsertIterative(2)
	tree.InsertIterative(5)
	tree.InsertIterative(6)
	tree.InsertIterative(1)
	tree.InsertIterative(8)
	tree.InsertIterative(3)
	tree.InsertIterative(10)

	tree.PrintPreOrder()

	tree.Remove(9)

	fmt.Println("Search for element 11")
	node := tree.Find(11)

	if node != nil {
		fmt.Printf("found=%d\n", node.Key)
	} else {
		fmt.Println("not found")
	}

	fmt.Println("Search for element 6")
	node = tree.Find(6)

	if node != nil {
		fmt.Printf("found=%d\n", node.Key)
	} else {
		fmt.Println("not found")
	}

	tree.PrintPreOrder()

	abb.FreeAbb(&tree)
}
