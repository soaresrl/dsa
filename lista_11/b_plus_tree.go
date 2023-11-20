package main

import (
	"fmt"

	bptree "github.com/soaresrl/dsa/pkg/b_plus_tree"
)

func main() {
	tree := bptree.NewBPTree(6)

	tree.Insert(1, "Data 1")
	tree.Insert(4, "Data 4")
	tree.Insert(7, "Data 7")
	tree.Insert(10, "Data 10")
	tree.Insert(5, "Data 5")
	tree.Insert(3, "Data 3")
	tree.Insert(2, "Data 2")
	tree.Insert(6, "Data 6")
	tree.Insert(8, "Data 8")
	tree.Insert(9, "Data 9")

	r := tree.Find(4)

	if r != nil {
		fmt.Println(r)
	} else {
		fmt.Println("Not Found")
	}

	tree.Remove(7)
	tree.Remove(6)
	tree.Remove(5)

	r = tree.Find(4)

	if r != nil {
		fmt.Println(r)
	} else {
		fmt.Println("Not Found")
	}

	bptree.FreeBTree(&tree)
}
