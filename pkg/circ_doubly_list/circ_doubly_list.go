package ord_doubly_list

import (
	"fmt"
)

type Node struct {
	data     int
	next     *Node
	previous *Node
}

func NewList() *Node {
	return nil
}

func (lst *Node) Insert(value int) *Node {
	node := &Node{data: value}

	if lst.IsEmpty() {
		node.next = node
		node.previous = node
	} else {

		node.next = lst
		node.previous = lst.previous

		lst.previous.next = node
		lst.previous = node
	}

	return node
}

func (lst *Node) IsEmpty() bool {
	return lst == nil
}

func (lst *Node) Find(value int) *Node {
	p := lst

	for ; p.next != lst; p = p.next {
		if p.data == value {
			return p
		}
	}

	// check last element
	if p.data == value {
		return p
	}

	return nil
}

func Free(lst **Node) {
	if *lst == nil {
		return
	}

	Free(&(*lst).next)
	(*lst).next = nil
	(*lst).previous = nil
	*lst = nil
}

func (lst *Node) Remove(value int) *Node {
	if lst == nil {
		return lst
	}

	// check if is the first and only element
	if lst.data == value && lst.next == lst {
		return nil
	}

	p := lst

	// iterate until find the node or value not present
	for ; p.next != lst && p.data != value; p = p.next {
	}

	// value not present case
	if p.next == lst && p.data != value {
		return lst
	}

	// if it's the first element
	if p == lst {
		p.next.previous = p.previous
		p.previous.next = p.next

		lst = p.next

		p = nil
	} else {
		p.next.previous = p.previous
		p.previous.next = p.next

		p = nil
	}

	return lst
}

func (lst *Node) RemoveRec(value int, sentinel *Node) *Node {
	if !lst.IsEmpty() {
		if lst.data == value {
			if lst == sentinel {
				lst.next.previous = lst.previous
				lst.previous.next = lst.next

				lst = lst.next
			} else {
				lst.next.previous = lst.previous
				lst.previous.next = lst.next
			}
		} else {
			lst.next.RemoveRec(value, sentinel)
		}
	}

	return lst
}

func (lst *Node) Print() {
	fmt.Printf("start \n")

	p := lst

	for ; p.next != lst; p = p.next {
		fmt.Printf("data=%d \n", p.data)
	}
	fmt.Printf("data=%d \n", p.data)

	fmt.Printf("end \n")
}

func (lst *Node) PrintRec(sentinel *Node) {
	if lst.next != sentinel {
		fmt.Printf("data = %d\n", lst.data)
		lst.next.PrintRec(sentinel)
	} else {
		fmt.Printf("data = %d\n", lst.data)
	}
}

func (lst *Node) PrintRecInv(sentinel *Node) {
	if lst.next != sentinel {
		lst.next.PrintRecInv(sentinel)
		fmt.Printf("data = %d\n", lst.data)
	} else {
		fmt.Printf("data = %d\n", lst.data)
	}
}
