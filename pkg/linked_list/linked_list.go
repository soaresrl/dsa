package linked_list

import (
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

func NewList() *Node {
	return nil
}

func (lst *Node) Insert(value int) *Node {
	node := &Node{Data: value}

	node.Next = lst

	return node
}

func (lst *Node) IsEmpty() bool {
	return lst == nil
}

func (lst *Node) Find(value int) *Node {
	for p := lst; p != nil; p = p.Next {
		if p.Data == value {
			return p
		}
	}

	return nil
}

func Free(lst **Node) {
	if *lst == nil {
		return
	}

	for p := *lst; p != nil; {
		temp := p.Next

		p.Next = nil

		p = temp
	}

	*lst = nil
}

func (lst *Node) Remove(value int) *Node {
	var slow *Node

	p := lst
	for ; p != nil && p.Data != value; p = p.Next {
		slow = p
	}

	// didn't find element
	if p == nil {
		return lst
	}

	// if it's the first element remove from beginning
	if slow == nil {
		lst = p.Next
	} else { // remove from middle
		slow.Next = p.Next
	}

	p.Next = nil
	p = nil

	return lst
}

func (lst *Node) RemoveRec(value int) *Node {
	if !lst.IsEmpty() {
		if lst.Data == value {
			t := lst

			lst = lst.Next

			t.Next = nil
			t = nil
		} else {
			lst.Next = lst.Next.RemoveRec(value)
		}
	}

	return lst
}

func (lst *Node) Print() {
	fmt.Printf("start \n")

	for p := lst; p != nil; p = p.Next {
		fmt.Printf("Data=%d \n", p.Data)
	}

	fmt.Printf("end \n")
}

func (lst *Node) PrintRec() {
	if lst != nil {
		fmt.Printf("Data = %d\n", lst.Data)
		lst.Next.PrintRec()
	}
}

func (lst *Node) PrintRecInv() {
	if lst != nil {
		lst.Next.PrintRecInv()
		fmt.Printf("Data = %d\n", lst.Data)
	}
}
