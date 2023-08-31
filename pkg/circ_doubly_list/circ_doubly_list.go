package circ_doubly_list

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

func (lst *Node) Insert(value int) *Node {
	node := &Node{Data: value}

	if lst.IsEmpty() {
		node.Next = node
		node.Previous = node
	} else {

		node.Next = lst
		node.Previous = lst.Previous

		lst.Previous.Next = node
		lst.Previous = node
	}

	return node
}

func (lst *Node) IsEmpty() bool {
	return lst == nil
}

func (lst *Node) Find(value int) *Node {
	p := lst

	for ; p.Next != lst; p = p.Next {
		if p.Data == value {
			return p
		}
	}

	// check last element
	if p.Data == value {
		return p
	}

	return nil
}

func Free(lst **Node) {
	if *lst == nil {
		return
	}

	if (*lst).Next == *lst {
		(*lst).Next = nil
		(*lst).Previous = nil

		*lst = nil
		return
	}

	p := *lst

	for ; p.Next != *lst; p = p.Next {
		temp := p.Next

		p.Next = nil
		p.Previous = nil

		p = temp
	}

	p.Next = nil
	p.Previous = nil

	*lst = nil
}

func (lst *Node) Remove(value int) *Node {
	if lst == nil {
		return lst
	}

	// check if is the first and only element
	if lst.Data == value && lst.Next == lst {
		lst.Next = nil
		lst.Previous = nil
		
		return nil
	}

	p := lst

	// iterate until find the node or value not present
	for ; p.Next != lst && p.Data != value; p = p.Next {
	}

	// value not present case
	if p.Next == lst && p.Data != value {
		return lst
	}

	// if it's the first element
	if p == lst {
		p.Next.Previous = p.Previous
		p.Previous.Next = p.Next

		lst = p.Next

		p = nil
	} else {
		p.Next.Previous = p.Previous
		p.Previous.Next = p.Next

		p = nil
	}

	return lst
}

func (lst *Node) RemoveRec(value int, sentinel *Node) *Node {
	if !lst.IsEmpty() {
		if lst.Data == value {
			if lst == sentinel {
				lst.Next.Previous = lst.Previous
				lst.Previous.Next = lst.Next

				lst = lst.Next
			} else {
				lst.Next.Previous = lst.Previous
				lst.Previous.Next = lst.Next
			}
		} else {
			lst.Next.RemoveRec(value, sentinel)
		}
	}

	return lst
}

func (lst *Node) Print() {
	fmt.Printf("start \n")

	p := lst

	for ; p.Next != lst; p = p.Next {
		fmt.Printf("Data=%d \n", p.Data)
	}
	fmt.Printf("Data=%d \n", p.Data)

	fmt.Printf("end \n")
}

func (lst *Node) PrintRec(sentinel *Node) {
	if lst.Next != sentinel {
		fmt.Printf("Data = %d\n", lst.Data)
		lst.Next.PrintRec(sentinel)
	} else {
		fmt.Printf("Data = %d\n", lst.Data)
	}
}

func (lst *Node) PrintRecInv(sentinel *Node) {
	if lst.Next != sentinel {
		lst.Next.PrintRecInv(sentinel)
		fmt.Printf("Data = %d\n", lst.Data)
	} else {
		fmt.Printf("Data = %d\n", lst.Data)
	}
}
