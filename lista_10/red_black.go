package main

import (
	"fmt"

	redblack "github.com/soaresrl/dsa/pkg/red_black"
)

func main() {
	tree := redblack.NewTree()

	tree.Insert(11)
	tree.Insert(2)
	tree.Insert(1)
	tree.Insert(5)
	tree.Insert(4)
	tree.Insert(8)
	tree.Insert(7)
	tree.Insert(14)
	tree.Insert(15)

	tree.PrintPreOrder()

	fmt.Println("Searching node with data = 4")

	n := tree.Find(4)

	if n != nil {
		fmt.Printf("Found Node Data = %d\n", n.Data)
	} else {
		fmt.Println("Node not Found")
	}

	redblack.Free(&tree.Root)

	tree.PrintPreOrder()
}
