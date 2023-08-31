package ord_doubly_list

import (
	"fmt"
)

type Node struct {
	Data     int
	Next     *Node
	Previous *Node
}

func NewList() *Node {
	return nil
}

func AreEqual(lst1 *Node, lst2 *Node) bool {
	a, b := lst1, lst2
	for ; a != nil && b != nil; a, b = a.Next, b.Next {
		if a.Data != b.Data {
			return false
		}
	}

	return a == b
}

func AreEqualRec(lst1 *Node, lst2 *Node) bool {
	if lst1 == nil && lst2 == nil {
		return true
	} else if lst1 == nil || lst2 == nil {
		return false
	} else {
		return (lst1.Data == lst2.Data) && AreEqualRec(lst1.Next, lst2.Next)
	}
}

func (lst *Node) Insert(value int) *Node {
	node := &Node{Data: value}

	var previous *Node

	p := lst

	if lst == nil {
		return node
	}

	for ; p != nil && p.Data < value; p = p.Next {
		previous = p
	}

	if previous == nil {
		node.Next = p
		p.Previous = node

		lst = node
	} else {
		previous.Next = node

		node.Previous = previous
		node.Next = p

		if p != nil {
			p.Previous = node
		}
	}

	return lst
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
		p.Previous = nil

		p = temp
	}

	*lst = nil
}

func (lst *Node) Remove(value int) *Node {
	p := lst.Find(value)

	if p == nil {
		return lst
	}

	// if it's the first element
	if lst == p {
		lst = p.Next
		lst.Previous = nil
	} else { // if not the first
		p.Previous.Next = p.Next
	}

	// if not the last element
	if p.Next != nil {
		p.Next.Previous = p.Previous
	}

	p = nil

	return lst
}

func (lst *Node) RemoveRec(value int) *Node {
	if !lst.IsEmpty() {
		if lst.Data == value {
			if lst.Previous == nil {
				lst = lst.Next
				lst.Previous = nil
			} else {
				lst.Previous.Next = lst.Next
			}

			if lst.Next != nil {
				lst.Next.Previous = lst.Previous
			}
		} else {
			lst.Next.RemoveRec(value)
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
