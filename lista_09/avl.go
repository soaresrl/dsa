package main

import (
	"fmt"

	"github.com/soaresrl/dsa/pkg/avl"
)

func main() {
	var tree *avl.Node

	tree = tree.Insert(7)
	tree = tree.Insert(4)
	tree = tree.Insert(9)
	tree = tree.Insert(2)
	tree = tree.Insert(5)
	tree = tree.Insert(6)
	tree = tree.Insert(1)
	tree = tree.Insert(8)
	tree = tree.Insert(3)
	tree = tree.Insert(10)

	tree.PrintPreOrder()

	fmt.Println("Removing item 9...")
	tree = tree.Remove(9)

	tree.PrintPreOrder()

	fmt.Println("Free list")
	avl.Free(&tree)

	tree.PrintPreOrder()
}
