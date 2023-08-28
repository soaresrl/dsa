package linked_list

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

func NewList() *Node {
	return nil
}

func (lst *Node) Insert(value int) *Node {
	node := &Node{data: value}

	node.next = lst

	if lst != nil {
		last := lst

		for ; last.next != lst; last = last.next {
		}

		last.next = node
	} else {
		node.next = node
	}

	return node
}

func (lst *Node) IsEmpty() bool {
	return lst == nil
}

func (lst *Node) Find(value int) *Node {
	for p := lst; p.next != lst; p = p.next {
		if p.data == value {
			return p
		}
	}

	return nil
}

func Free(lst **Node) {
	if *lst == nil {
		return
	}

	Free(&(*lst).next)
	(*lst).next = nil
	*lst = nil
}

func (lst *Node) Remove(value int) *Node {
	var slow *Node

	if lst == nil {
		return lst
	}

	// The list contains only one node and we want to remove it
	if lst.data == value && lst.next == lst {
		return nil
	}

	// Want to remove the first element
	if lst.data == value {
		last := lst

		// Find reference to last element
		for ; last.next != lst; last = last.next {
		}

		// Make it point no next of first
		last.next = lst.next

		return last.next
	}

	// Search middle or end of list case
	last := lst
	for ; last.next != lst && last.next.data != value; last = last.next {
	}

	if last.next.data == value {
		p := last.next

		last.next = p.next
		p = nil
	}

	return lst
}

func (lst *Node) RemoveRec(value int) *Node {
	if !lst.IsEmpty() {
		if lst.data == value {
			lst = lst.next
		} else {
			lst.next = lst.next.RemoveRec(value)
		}
	}

	return lst
}

func (lst *Node) Print() {
	fmt.Printf("start \n")

	for p := lst; p != nil; p = p.next {
		fmt.Printf("data=%d \n", p.data)
	}

	fmt.Printf("end \n")
}

func (lst *Node) PrintRec() {
	if lst != nil {
		fmt.Printf("data = %d\n", lst.data)
		lst.next.PrintRec()
	}
}

func (lst *Node) PrintRecInv() {
	if lst != nil {
		lst.next.PrintRecInv()
		fmt.Printf("data = %d\n", lst.data)
	}
}
