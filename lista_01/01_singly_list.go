package main 

import (
	"fmt"
	"github.com/soaresrl/dsa/pkg/linked_list"
)

func main() {
	lst := linked_list.NewList()

	lst = lst.Insert(10)
	lst = lst.Insert(12)
	lst = lst.Insert(608)
	lst = lst.Insert(42)
	lst = lst.Insert(50)
	lst = lst.Insert(102)

	fmt.Println(">> Printing list <<")
	lst.Print()
	fmt.Println(">>      End      <<")

	fmt.Println(">> Printing list recursively <<")
	lst.PrintRec()
	fmt.Println(">>            End            <<")

	fmt.Println(">> Printing list inverted <<")
	lst.PrintRecInv()
	fmt.Println(">>            End            <<")

	fmt.Println(">> Checking if is empty   <<")
	fmt.Printf("Is Empty = %v\n", lst.IsEmpty())
	fmt.Println(">>          End           <<")

	fmt.Println(">> Searching element 42   <<")

	node1 := lst.Find(42)

	node1_next := node1.Next

	fmt.Printf("Data = %d\n", node1.Data)
	fmt.Printf("Next = %d\n", node1_next.Data)
	fmt.Println(">>          End           <<")

	fmt.Println(">> Removing from the beginning <<")
	lst = lst.Remove(102)
	lst.Print()
	fmt.Println(">>             End             <<")

	fmt.Println(">> Removing from the end <<")
	lst = lst.Remove(10)
	lst.Print()
	fmt.Println(">>             End             <<")

	fmt.Println(">> Removing from the middle <<")
	lst = lst.Remove(42)
	lst.Print()
	fmt.Println(">>             End             <<")

	linked_list.Free(&lst)
	
	fmt.Println(">> Checking if is empty   <<")
	fmt.Printf("Is Empty = %v\n", lst.IsEmpty())
	fmt.Println(">>          End           <<")
}
